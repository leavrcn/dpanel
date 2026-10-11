package tools

// store 模块工具（对齐 DPanel v1.11.1）。
func init() {
	registerSlice([]ToolDef{
		{
			Name:        "dpanel_store_list",
			Description: "应用商店应用列表",
			Route:       "/common/store/get-list",
			Minimum:     "read-only",
			Destructive: false,
		},
		{
			Name:        "dpanel_store_sync",
			Description: "同步应用商店数据",
			Route:       "/common/store/sync",
			Params: []ParamDef{
			{Name: "name", Type: "string", Description: "", Required: true},
			{Name: "store_type", Type: "string", Description: "", Required: true, PayloadKey: "type"},
			{Name: "url", Type: "string", Description: "", Required: true},
		},
			Minimum:     "read-write",
			Destructive: false,
		},
		{
			Name:        "dpanel_store_deploy",
			Description: "从应用商店部署应用。store_id 为商店应用 id",
			Route:       "/common/store/deploy",
			Params: []ParamDef{
			{Name: "store_id", Type: "number", Description: "", Required: true, PayloadKey: "storeId", Normalize: "int"},
			{Name: "option", Type: "string", Description: "", Required: true, Normalize: "json"},
		},
			Minimum:     "read-write",
			Destructive: false,
		},
		{
			Name:        "dpanel_store_create",
			Description: "添加应用商店源",
			Route:       "/common/store/create",
			Params: []ParamDef{
			{Name: "name", Type: "string", Description: "", Required: true},
			{Name: "title", Type: "string", Description: "", Required: true},
			{Name: "store_type", Type: "string", Description: "", Required: true, PayloadKey: "type"},
			{Name: "url", Type: "string", Description: "", Required: true},
			{Name: "apps", Type: "string", Description: "", Required: true, Normalize: "json"},
		},
			Minimum:     "read-write",
			Destructive: false,
		},
		{
			Name:        "dpanel_store_delete",
			Description: "删除应用商店源（破坏性，需 confirm=true）",
			Route:       "/common/store/delete",
			Params: []ParamDef{
			{Name: "id", Type: "string", Description: "", Required: true, Normalize: "int-list"},
			{Name: "confirm", Type: "boolean", Description: "", Required: false},
		},
			Minimum:     "admin",
			Destructive: true,
		},
	})
}
