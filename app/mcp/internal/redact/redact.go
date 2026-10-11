// Package redact 工具返回值中敏感值（密码、token、secret）的递归脱敏。
// 对任意 JSON 结构在序列化给 MCP 客户端前执行；脱敏始终开启。
package redact

import (
	"regexp"
	"strings"
)

const Mask = "******"

// 始终脱敏的键名（小写、分隔符归一化后匹配）
var sensitiveKeys = map[string]struct{}{
	"password": {}, "passwd": {}, "secret": {}, "token": {}, "accesstoken": {},
	"refreshtoken": {}, "apikey": {}, "accesskey": {}, "secretkey": {},
	"privatekey": {}, "jwt": {}, "authorization": {}, "cookie": {}, "session": {},
	"credential": {}, "credentials": {}, "webhooksecret": {}, "clientsecret": {},
	"sshkey": {}, "tlskey": {}, "certkey": {}, "passphrase": {},
}

var (
	normalizeRe = regexp.MustCompile(`[\s_\-.]+`)
	// 环境变量风格：NAME=VALUE 且 NAME 含敏感特征
	envRe = regexp.MustCompile(`(?i)\b([A-Z0-9_]*(?:PASSWORD|PASSWD|SECRET|TOKEN|KEY|CREDENTIAL|JWT)[A-Z0-9_]*)\s*=\s*([^\s"']+)`)
	// 自由文本中的 JWT（xxx.yyy.zzz）
	jwtRe = regexp.MustCompile(`\beyJ[A-Za-z0-9_-]{10,}\.[A-Za-z0-9_-]{10,}\.[A-Za-z0-9_-]{5,}\b`)
)

func normalizeKey(key string) string {
	return normalizeRe.ReplaceAllString(strings.ToLower(key), "")
}

func isSensitiveKey(key string) bool {
	norm := normalizeKey(key)
	for k := range sensitiveKeys {
		if norm == k || strings.HasSuffix(norm, k) {
			return true
		}
	}
	return false
}

// Text 对自由文本（如日志）中 KEY=VALUE 形式赋值与 JWT 做脱敏。
func Text(text string) string {
	text = envRe.ReplaceAllString(text, "$1="+Mask)
	text = jwtRe.ReplaceAllString(text, Mask)
	return text
}

// Value 递归脱敏 JSON 结构（map[string]any / []any / string）。
func Value(v any) any {
	switch t := v.(type) {
	case map[string]any:
		out := make(map[string]any, len(t))
		for k, val := range t {
			if isSensitiveKey(k) {
				out[k] = Mask
			} else {
				out[k] = Value(val)
			}
		}
		return out
	case []any:
		out := make([]any, len(t))
		for i, val := range t {
			out[i] = Value(val)
		}
		return out
	case string:
		if jwtRe.MatchString(t) || envRe.MatchString(t) {
			return Text(t)
		}
		return t
	default:
		return v
	}
}
