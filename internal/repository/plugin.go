package repository

import "github.com/issueye/grape/internal/common/model"

type CreatePlugin struct {
	model.PluginBase
}

type ModifyPlugin struct {
	model.PluginInfo
}
