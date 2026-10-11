package tools

// setting 模块工具（对齐 DPanel v1.11.1）。
func init() {
	registerSlice([]ToolDef{
		{
			Name:        "dpanel_setting_get",
			Description: "读取 DPanel 配置项。group_name 如 setting；name 如 DPanelInfo/docker",
			Route:       "/common/setting/get-setting",
			Params: []ParamDef{
			{Name: "group_name", Type: "string", Description: "", Required: true, PayloadKey: "groupName"},
			{Name: "name", Type: "string", Description: "", Required: true},
		},
			Minimum:     "read-only",
			Destructive: false,
		},
	})
}
