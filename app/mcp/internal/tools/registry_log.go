package tools

// log 模块工具（对齐 DPanel v1.11.1）。
func init() {
	registerSlice([]ToolDef{
		{
			Name:        "dpanel_log_list",
			Description: "DPanel 系统日志列表",
			Route:       "/common/log/get-list",
			Params: []ParamDef{
			{Name: "page", Type: "number", Description: "", Required: false, Normalize: "int", Default: 1},
			{Name: "page_size", Type: "number", Description: "", Required: false, PayloadKey: "pageSize", Normalize: "int", Default: 50},
		},
			Minimum:     "read-only",
			Destructive: false,
		},
	})
}
