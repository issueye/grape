package logic

import (
	"github.com/issueye/grape/internal/repository"
	"github.com/issueye/grape/internal/service"
)

type Plugin struct{}

func (p *Plugin) CreatePlugin(req *repository.CreatePlugin) error {
	srv := service.NewPlugin()
	err := srv.Create(req)

    return err
}
