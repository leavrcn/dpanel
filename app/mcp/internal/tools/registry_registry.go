package tools

// registry 模块工具（对齐 DPanel v1.11.1）。
func init() {
	registerSlice([]ToolDef{
		{
			Name:        "dpanel_registry_list",
			Description: "镜像仓库列表",
			Route:       "/common/registry/get-list",
			Minimum:     "read-only",
			Destructive: false,
		},
		{
			Name:        "dpanel_registry_detail",
			Description: "镜像仓库详情",
			Route:       "/common/registry/get-detail",
			Params: []ParamDef{
			{Name: "id", Type: "number", Description: "", Required: true, Normalize: "int"},
		},
			Minimum:     "read-only",
			Destructive: false,
		},
		{
			Name:        "dpanel_registry_create",
			Description: "添加镜像仓库",
			Route:       "/common/registry/create",
			Params: []ParamDef{
			{Name: "title", Type: "string", Description: "", Required: true},
			{Name: "server_address", Type: "string", Description: "", Required: true, PayloadKey: "serverAddress"},
			{Name: "option", Type: "string", Description: "", Required: false, Normalize: "json", Merge: true},
		},
			Minimum:     "read-write",
			Destructive: false,
		},
		{
			Name:        "dpanel_registry_delete",
			Description: "删除镜像仓库（破坏性，需 confirm=true）",
			Route:       "/common/registry/delete",
			Params: []ParamDef{
			{Name: "id", Type: "string", Description: "", Required: true, Normalize: "int-list"},
			{Name: "confirm", Type: "boolean", Description: "", Required: false},
		},
			Minimum:     "admin",
			Destructive: true,
		},
	})
}
