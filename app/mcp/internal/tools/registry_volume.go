package tools

// volume 模块工具（对齐 DPanel v1.11.1）。
func init() {
	registerSlice([]ToolDef{
		{
			Name:        "dpanel_volume_list",
			Description: "卷列表（含使用状态）",
			Route:       "/app/volume/get-list",
			Minimum:     "read-only",
			Destructive: false,
		},
		{
			Name:        "dpanel_volume_detail",
			Description: "卷详情",
			Route:       "/app/volume/get-detail",
			Params: []ParamDef{
			{Name: "name", Type: "string", Description: "", Required: true},
		},
			Minimum:     "read-only",
			Destructive: false,
		},
		{
			Name:        "dpanel_volume_create",
			Description: "创建卷",
			Route:       "/app/volume/create",
			Params: []ParamDef{
			{Name: "name", Type: "string", Description: "", Required: true},
			{Name: "driver", Type: "string", Description: "", Required: false, Default: "local"},
		},
			Minimum:     "read-write",
			Destructive: false,
		},
		{
			Name:        "dpanel_volume_delete",
			Description: "删除卷（破坏性，需 confirm=true）。name 为卷名数组",
			Route:       "/app/volume/delete",
			Params: []ParamDef{
			{Name: "name", Type: "string", Description: "", Required: true, Normalize: "str-list"},
			{Name: "confirm", Type: "boolean", Description: "", Required: false},
		},
			Minimum:     "admin",
			Destructive: true,
		},
		{
			Name:        "dpanel_volume_prune",
			Description: "清理未使用卷（破坏性，需 confirm=true）",
			Route:       "/app/volume/prune",
			Params: []ParamDef{
			{Name: "confirm", Type: "boolean", Description: "", Required: false},
		},
			Minimum:     "admin",
			Destructive: true,
		},
	})
}
