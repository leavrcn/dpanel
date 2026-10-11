package tools

// cron 模块工具（对齐 DPanel v1.11.1）。
func init() {
	registerSlice([]ToolDef{
		{
			Name:        "dpanel_cron_list",
			Description: "计划任务列表",
			Route:       "/common/cron/get-list",
			Minimum:     "read-only",
			Destructive: false,
		},
		{
			Name:        "dpanel_cron_detail",
			Description: "计划任务详情",
			Route:       "/common/cron/get-detail",
			Params: []ParamDef{
			{Name: "id", Type: "number", Description: "", Required: true, Normalize: "int"},
		},
			Minimum:     "read-only",
			Destructive: false,
		},
		{
			Name:        "dpanel_cron_log_list",
			Description: "计划任务执行日志",
			Route:       "/common/cron/get-log-list",
			Params: []ParamDef{
			{Name: "id", Type: "number", Description: "", Required: true, Normalize: "int"},
		},
			Minimum:     "read-only",
			Destructive: false,
		},
		{
			Name:        "dpanel_cron_template",
			Description: "计划任务模板列表（10 个官方模板）",
			Route:       "/common/cron/template",
			Minimum:     "read-only",
			Destructive: false,
		},
		{
			Name:        "dpanel_cron_create",
			Description: "创建计划任务。trigger_type ∈ cron|event|manual。spec 为 cron 表达式（trigger=cron 时）",
			Route:       "/common/cron/create",
			Params: []ParamDef{
				{Name: "title", Type: "string", Description: "", Required: true},
				{Name: "spec", Type: "string", Description: "cron 表达式（6 字段：秒 分 时 日 月 周，如 '0 3 * * * *'）或 @manual", Required: true, PayloadKey: "expression", Normalize: "cron-expression"},
				{Name: "trigger_type", Type: "string", Description: "", Required: true, PayloadKey: "triggerType", Enum: []string{"cron", "event", "manual"}},
				{Name: "script", Type: "string", Description: "任务脚本（run-in-dpanel 时为 shell，否则为容器内命令）", Required: false},
				{Name: "environment", Type: "string", Description: "环境变量 JSON 数组 [{name, value}]", Required: false, Normalize: "json"},
				{Name: "container_name", Type: "string", Description: "目标容器名（容器内执行时）", Required: false, PayloadKey: "containerName"},
		},
			Minimum:     "read-write",
			Destructive: false,
		},
		{
			Name:        "dpanel_cron_run_once",
			Description: "立即执行一次计划任务",
			Route:       "/common/cron/run-once",
			Params: []ParamDef{
			{Name: "id", Type: "number", Description: "", Required: true, Normalize: "int"},
		},
			Minimum:     "read-write",
			Destructive: false,
		},
		{
			Name:        "dpanel_cron_delete",
			Description: "删除计划任务（破坏性，需 confirm=true）",
			Route:       "/common/cron/delete",
			Params: []ParamDef{
			{Name: "id", Type: "string", Description: "", Required: true, Normalize: "int-list"},
			{Name: "confirm", Type: "boolean", Description: "", Required: false},
		},
			Minimum:     "admin",
			Destructive: true,
		},
		{
			Name:        "dpanel_cron_prune_log",
			Description: "清理计划任务日志（破坏性，需 confirm=true）",
			Route:       "/common/cron/prune-log",
			Params: []ParamDef{
			{Name: "id", Type: "number", Description: "", Required: true, Normalize: "int"},
			{Name: "confirm", Type: "boolean", Description: "", Required: false},
		},
			Minimum:     "admin",
			Destructive: true,
		},
	})
}
