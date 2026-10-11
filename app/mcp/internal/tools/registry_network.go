package tools

// network 模块工具（对齐 DPanel v1.11.1）。
func init() {
	registerSlice([]ToolDef{
		{
			Name:        "dpanel_network_list",
			Description: "网络列表",
			Route:       "/app/network/get-list",
			Minimum:     "read-only",
			Destructive: false,
		},
		{
			Name:        "dpanel_network_detail",
			Description: "网络详情",
			Route:       "/app/network/get-detail",
			Params: []ParamDef{
			{Name: "name", Type: "string", Description: "", Required: true},
		},
			Minimum:     "read-only",
			Destructive: false,
		},
		{
			Name:        "dpanel_network_container_list",
			Description: "网络上的容器列表",
			Route:       "/app/network/get-container-list",
			Params: []ParamDef{
			{Name: "name", Type: "string", Description: "", Required: true, Normalize: "str-list"},
		},
			Minimum:     "read-only",
			Destructive: false,
		},
		{
			Name:        "dpanel_network_create",
			Description: "创建网络。driver ∈ bridge|macvlan|ipvlan|overlay",
			Route:       "/app/network/create",
			Params: []ParamDef{
			{Name: "name", Type: "string", Description: "", Required: true},
			{Name: "driver", Type: "string", Description: "", Required: false, Enum: []string{"bridge", "macvlan", "ipvlan", "overlay"}, Default: "bridge"},
		},
			Minimum:     "read-write",
			Destructive: false,
		},
		{
			Name:        "dpanel_network_connect",
			Description: "将容器接入网络",
			Route:       "/app/network/connect",
			Params: []ParamDef{
			{Name: "name", Type: "string", Description: "", Required: true},
			{Name: "container_name", Type: "string", Description: "", Required: true, PayloadKey: "containerName"},
		},
			Minimum:     "read-write",
			Destructive: false,
		},
		{
			Name:        "dpanel_network_disconnect",
			Description: "将容器从网络断开",
			Route:       "/app/network/disconnect",
			Params: []ParamDef{
			{Name: "name", Type: "string", Description: "", Required: true},
			{Name: "container_name", Type: "string", Description: "", Required: true, PayloadKey: "containerName"},
		},
			Minimum:     "read-write",
			Destructive: false,
		},
		{
			Name:        "dpanel_network_delete",
			Description: "删除网络（破坏性，需 confirm=true）",
			Route:       "/app/network/delete",
			Params: []ParamDef{
			{Name: "name", Type: "string", Description: "", Required: true, Normalize: "str-list"},
			{Name: "confirm", Type: "boolean", Description: "", Required: false},
		},
			Minimum:     "admin",
			Destructive: true,
		},
		{
			Name:        "dpanel_network_prune",
			Description: "清理未使用网络（破坏性，需 confirm=true）",
			Route:       "/app/network/prune",
			Params: []ParamDef{
			{Name: "confirm", Type: "boolean", Description: "", Required: false},
		},
			Minimum:     "admin",
			Destructive: true,
		},
	})
}
