#!/usr/bin/env python3
"""内置版 E2E：单容器 DPanel（feat/mcp-native）的 MCP 端点全链路验证。

覆盖：握手、tools/list 99、只读工具真实数据（回环调用）、
fail-closed 负例、脱敏、参数归一化、mountPoint 校验。
"""
import json, subprocess, sys, os, urllib.request

BASE = "http://127.0.0.1:18907"
MCP_TOKEN = os.environ.get("MCP_TOKEN", "")

def http_post(path, body, headers=None):
    req = urllib.request.Request(BASE + path, data=json.dumps(body).encode(),
                                 headers={"Content-Type": "application/json",
                                          "Accept": "application/json, text/event-stream",
                                          **(headers or {})})
    with urllib.request.urlopen(req) as resp:
        return resp.headers, resp.read().decode()

passed = failed = 0
def check(name, ok, detail=""):
    global passed, failed
    if ok:
        passed += 1
        print(f"PASS: {name}")
    else:
        failed += 1
        print(f"FAIL: {name} {detail[:200]}")

# 1. 握手 + session
hdrs, body = http_post("/dpanel/mcp", {
    "jsonrpc": "2.0", "id": 1, "method": "initialize",
    "params": {"protocolVersion": "2025-03-26", "capabilities": {},
               "clientInfo": {"name": "e2e", "version": "1.0"}}},
    {"Authorization": "Bearer " + MCP_TOKEN})
session = hdrs.get("mcp-session-id", "")
check("initialize 握手", "serverInfo" in body, body)
check("session id", bool(session))

H = {"Authorization": "Bearer " + MCP_TOKEN, "mcp-session-id": session}

def call(tool, args):
    _, body = http_post("/dpanel/mcp", {
        "jsonrpc": "2.0", "id": 2, "method": "tools/call",
        "params": {"name": tool, "arguments": args}}, H)
    return unwrap(body)

def unwrap(body):
    """SSE/JSON-RPC 解包到工具 text"""
    try:
        data = body
        if data.startswith("event:") or "\ndata:" in data:
            for line in data.split("\n"):
                if line.startswith("data:"):
                    data = line[5:].strip()
                    break
        envelope = json.loads(data)
        return envelope.get("result", {}).get("content", [{}])[0].get("text", "")
    except Exception:
        return body

# 2. tools/list
_, body = http_post("/dpanel/mcp", {"jsonrpc": "2.0", "id": 3, "method": "tools/list"}, H)
# tools/list 响应：result.tools 数组（非 content）
try:
    data = body
    if "data:" in data:
        for line in data.split("\n"):
            if line.startswith("data:"):
                data = line[5:].strip()
                break
    n = json.loads(data).get("result", {}).get("tools", []).__len__() if "tools" in json.loads(data).get("result", {}) else data.count('"name":"dpanel_')
except Exception:
    n = 0
check("tools/list 99 工具", n == 99, f"实际 {n}")

# 3. 只读工具（回环调用）
for tool, needle in [
    ("dpanel_system_info", '"version"'),
    ("dpanel_env_list", '"local"'),
    ("dpanel_container_list", '"list"'),
    ("dpanel_compose_list", '"list"'),
    ("dpanel_image_list", '"list"'),
    ("dpanel_network_list", '"bridge"'),
    ("dpanel_volume_list", '"list"'),
    ("dpanel_cron_list", '"list"'),
    ("dpanel_registry_list", '"docker.io"'),
]:
    text = call(tool, {})
    check(tool, needle in text, text[:200])

# 4. 参数归一化
text = call("dpanel_network_container_list", {"name": "bridge"})
check("network_container_list 归一化", "networkName" in text, text[:200])

# 5. fail-closed 负例
for tool, args in [
    ("dpanel_container_delete", {"md5": "nonexistent"}),
    ("dpanel_image_delete", {"md5": "nonexistent"}),
    ("dpanel_registry_delete", {"id": 999}),
]:
    text = call(tool, args)
    check(f"{tool} fail-closed/分级拦截", "fail-closed" in text or "能力分级" in text, text[:200])

# 6. 脱敏
text = call("dpanel_system_info", {})
check("founder.password 脱敏", '"password":"******"' in text, text[:200])

# 7. mountPoint 校验负例
text = call("dpanel_explorer_list", {"mount_point": "garbage", "path": "/"})
check("mount_point 校验负例", "volume:" in text, text[:200])

# 8. 必填参数缺失
text = call("dpanel_container_detail", {})
check("必填参数缺失负例", "缺少必填参数" in text, text[:200])

# 9. 门禁：错 token 401
try:
    http_post("/dpanel/mcp", {"jsonrpc": "2.0", "id": 9, "method": "tools/list"},
              {"Authorization": "Bearer wrong"})
    check("错 token 401", False, "未拒绝")
except urllib.error.HTTPError as e:
    check("错 token 401", e.code == 401, str(e.code))

print(f"\n=== 内置版 E2E: PASS={passed} FAIL={failed} ===")
sys.exit(0 if failed == 0 else 1)
