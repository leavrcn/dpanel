package main

import (
	"github.com/donknap/dpanel/app/mcp"
	"github.com/donknap/dpanel/common/accessor"
	"github.com/donknap/dpanel/common/dao"
	"github.com/donknap/dpanel/common/entity"
)

// mcpStore mcp.Store 的 settings 表实现（主仓侧）。
// MCP 配置存 ims_setting 表（group=mcp, name=config/service），Value 走
// accessor.SettingValueOption.Mcp 强类型字段（对齐上游配置风格）。
type mcpStore struct{}

const (
	mcpSettingGroup       = "mcp"
	mcpSettingNameConfig  = "config"
	mcpSettingNameService = "service"
)

func (mcpStore) save(group, name string, mutate func(opt *accessor.SettingValueOption)) error {
	row, _ := dao.Setting.
		Where(dao.Setting.GroupName.Eq(group)).
		Where(dao.Setting.Name.Eq(name)).
		First()
	if row == nil {
		row = &entity.Setting{GroupName: group, Name: name, Value: &accessor.SettingValueOption{}}
	}
	if row.Value == nil {
		row.Value = &accessor.SettingValueOption{}
	}
	mutate(row.Value)
	old, _ := dao.Setting.
		Where(dao.Setting.GroupName.Eq(group)).
		Where(dao.Setting.Name.Eq(name)).
		First()
	if old == nil {
		return dao.Setting.Create(row)
	}
	_, err := dao.Setting.
		Where(dao.Setting.GroupName.Eq(group)).
		Where(dao.Setting.Name.Eq(name)).
		Updates(row)
	return err
}

func (mcpStore) loadMcp() *accessor.McpConfig {
	row, err := dao.Setting.
		Where(dao.Setting.GroupName.Eq(mcpSettingGroup)).
		Where(dao.Setting.Name.Eq(mcpSettingNameConfig)).
		First()
	if err == nil && row != nil && row.Value != nil && row.Value.Mcp != nil {
		return row.Value.Mcp
	}
	return nil
}

func (s mcpStore) LoadMcpConfig() (enable bool, token, profile string) {
	if c := s.loadMcp(); c != nil {
		return c.Enable, c.Token, c.Profile
	}
	return false, "", ""
}

func (s mcpStore) SaveMcpConfig(enable bool, token, profile string) error {
	return s.save(mcpSettingGroup, mcpSettingNameConfig, func(opt *accessor.SettingValueOption) {
		opt.Mcp = &accessor.McpConfig{Enable: enable, Token: token, Profile: profile}
	})
}

func (s mcpStore) LoadServiceCredential() (username, password string) {
	row, err := dao.Setting.
		Where(dao.Setting.GroupName.Eq(mcpSettingGroup)).
		Where(dao.Setting.Name.Eq(mcpSettingNameService)).
		First()
	if err == nil && row != nil && row.Value != nil && row.Value.Mcp != nil &&
		row.Value.Mcp.ServiceCredential != nil {
		return row.Value.Mcp.ServiceCredential.Username, row.Value.Mcp.ServiceCredential.Password
	}
	return "", ""
}

func (s mcpStore) SaveServiceCredential(username, password string) error {
	return s.save(mcpSettingGroup, mcpSettingNameService, func(opt *accessor.SettingValueOption) {
		opt.Mcp = &accessor.McpConfig{
			ServiceCredential: &accessor.McpServiceCredential{Username: username, Password: password},
		}
	})
}

var _ mcp.Store = mcpStore{}
