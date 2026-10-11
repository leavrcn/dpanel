package tools

// env 模块工具（对齐 DPanel v1.11.1）。
func init() {
	registerSlice([]ToolDef{
		{
			Name:        "dpanel_env_list",
			Description: "列出 Docker 环境（多主机），含当前环境",
			Route:       "/common/env/get-list",
			Minimum:     "read-only",
			Destructive: false,
		},
		{
			Name:        "dpanel_env_detail",
			Description: "Docker 环境详情",
			Route:       "/common/env/get-detail",
			Params: []ParamDef{
			{Name: "name", Type: "string", Description: "", Required: true},
		},
			Minimum:     "read-only",
			Destructive: false,
		},
		{
			Name:        "dpanel_env_switch",
			Description: "切换当前 Docker 环境（多主机）",
			Route:       "/common/env/switch",
			Params: []ParamDef{
			{Name: "env_name", Type: "string", Description: "", Required: true, PayloadKey: "name"},
		},
			Minimum:     "read-write",
			Destructive: false,
		},
		{
			Name:        "dpanel_env_create",
			Description: "创建 Docker 环境（连接远程 Docker host）",
			Route:       "/common/env/create",
			Params: []ParamDef{
			{Name: "name", Type: "string", Description: "", Required: true},
			{Name: "title", Type: "string", Description: "", Required: true},
			{Name: "address", Type: "string", Description: "", Required: true},
		},
			Minimum:     "read-write",
			Destructive: false,
		},
		{
			Name:        "dpanel_env_delete",
			Description: "删除 Docker 环境（破坏性，需 confirm=true）",
			Route:       "/common/env/delete",
			Params: []ParamDef{
			{Name: "env_name", Type: "string", Description: "", Required: true, PayloadKey: "name", Normalize: "str-list"},
			{Name: "confirm", Type: "boolean", Description: "", Required: false},
		},
			Minimum:     "admin",
			Destructive: true,
		},
	})
}
