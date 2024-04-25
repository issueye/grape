package service

import (
	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/internal/common/service"
	"github.com/issueye/grape/internal/global"
	"github.com/issueye/grape/internal/repository"
)

type Plugin struct {
	*service.BaseService
}

func NewPlugin() *Plugin {
	return &Plugin{
		BaseService: service.NewBaseService(global.DB),
	}
}

// Create
// 创建信息
func (s *Plugin) Create(data *repository.CreatePlugin) error {
	info := model.PluginInfo{}.New()
	info.Name = data.Name
	info.Path = data.Path
	info.Version = data.Version
	info.Key = data.Key
	info.Value = data.Value

	return s.Db.Model(info).Create(info).Error
}
