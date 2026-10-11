package tools

// system 模块工具（对齐 DPanel v1.11.1）。
func init() {
	registerSlice([]ToolDef{
		{
			Name:        "dpanel_system_info",
			Description: "DPanel 系统信息：版本、运行环境、Docker 环境状态、数据挂载",
			Route:       "/common/home/info",
			Minimum:     "read-only",
			Destructive: false,
		},
		{
			Name:        "dpanel_system_usage",
			Description: "DPanel 所在主机资源使用情况（CPU/内存/磁盘/面板占用）",
			Route:       "/common/panel/usage",
			Minimum:     "read-only",
			Destructive: false,
		},
		{
			Name:        "dpanel_system_stat_list",
			Description: "DPanel 多主机统计列表",
			Route:       "/common/home/get-stat-list",
			Minimum:     "read-only",
			Destructive: false,
		},
	})
}
