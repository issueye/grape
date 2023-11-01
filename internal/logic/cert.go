package logic

import (
	"fmt"

	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/internal/repository"
	"github.com/issueye/grape/internal/service"
)

type Cert struct{}

func (Cert) Get(req *repository.QueryCert) ([]*model.CertInfo, error) {
	return service.NewCert().Query(req)
}

func (Cert) GetById(id string) (*model.CertInfo, error) {
	return service.NewCert().FindById(id)
}

// Modify
// 修改信息 不包含状态
func (Cert) Modify(req *repository.ModifyCert) error {
	CertService := service.NewCert()
	return CertService.Modify(req)
}

// Create
// 创建数据
func (Cert) Create(req *repository.CreateCert) error {
	// 判断端口号在当前系统是否已经被使用
	CertService := service.NewCert()
	info, err := CertService.FindByName(req.Name)
	if err != nil {
		return fmt.Errorf("检查证书失败 %s", err.Error())
	}

	if info.ID != "" {
		return fmt.Errorf("目标证书[%s]已经添加，请勿重复添加", info.Name)
	}

	// 创建数据
	err = CertService.Create(req)
	if err != nil {
		return fmt.Errorf("创建信息失败 %s", err.Error())
	}

	return nil
}

// Del
// 根据ID删除信息
func (Cert) Del(id string) error {
	CertService := service.NewCert()

	// 检查使用状态，如果是正在使用则不允许删除
	_, err := CertService.FindById(id)
	if err != nil {
		return err
	}

	return CertService.Del(id)
}
