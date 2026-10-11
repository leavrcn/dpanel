// Package profile 能力分级门禁：read-only / read-write / admin 三级。
//
// 破坏性工具在任何分级下都要求调用方显式传 confirm=true（fail-closed 人在回路确认）。
package profile

import "fmt"

const (
	ReadReadOnly  = "read-only"
	ReadReadWrite = "read-write"
	Admin         = "admin"
)

var rank = map[string]int{ReadReadOnly: 0, ReadReadWrite: 1, Admin: 2}

// Gate 按 server 配置的 profile 检查工具可用性。
type Gate struct {
	profile string
}

func New(p string) *Gate {
	return &Gate{profile: p}
}

// Allowed 工具是否在当前分级下可用。
func (g *Gate) Allowed(tool string, minimum string) bool {
	r, ok := rank[g.profile]
	if !ok {
		return false
	}
	m, ok := rank[minimum]
	if !ok {
		return false
	}
	return r >= m
}

// ExplainDenial 拒绝原因（给 agent 的可操作提示）。
func (g *Gate) ExplainDenial(tool, minimum string) string {
	return fmt.Sprintf("%s 需要 %s 及以上能力分级，当前 server 分级为 %s（DPANEL_MCP_PROFILE）", tool, minimum, g.profile)
}

// ConfirmRequired 破坏性工具未确认时的 fail-closed 提示。
func ConfirmRequired(tool string) error {
	return fmt.Errorf("%s 是破坏性操作，请确认影响后重发请求并携带 confirm=true（fail-closed）", tool)
}
