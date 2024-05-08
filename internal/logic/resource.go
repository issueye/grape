package logic

import (
	"fmt"

	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/internal/repository"
	"github.com/issueye/grape/internal/service"
)

type Resource struct{}

func (Resource) Get(req *repository.QueryResource) ([]*model.ResourceInfo, error) {
	return service.NewResource().Query(req)
}

func (Resource) GetById(id string) (*model.ResourceInfo, error) {
	return service.NewResource().FindById(id)
}

// Modify
// 修改信息 不包含状态
func (Resource) Modify(req *repository.ModifyResource) error {
	ResourceService := service.NewResource()
	return ResourceService.Modify(req)
}

// Modify
// 修改信息 不包含状态
func (Resource) ModifyByMap(id string, datas map[string]any) error {
	ResourceService := service.NewResource()
	return ResourceService.ModifyByMap(id, datas)
}

// Create
// 创建数据
func (Resource) Create(req *repository.CreateResource) error {
	ResourceService := service.NewResource()

	// 创建数据
	err := ResourceService.Create(req)
	if err != nil {
		return fmt.Errorf("创建资源失败 %s", err.Error())
	}

	return nil
}

// Del
// 根据ID删除信息
func (Resource) Del(id string) error {
	ResourceService := service.NewResource()
	return ResourceService.Del(id)
}
