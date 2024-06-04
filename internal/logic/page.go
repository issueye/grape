package logic

import (
	"errors"
	"fmt"
	"strings"

	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/internal/repository"
	"github.com/issueye/grape/internal/service"
)

type Page struct{}

func (Page) Get(req *repository.QueryPage) ([]*model.PageInfo, error) {
	return service.NewPage().Query(req)
}

func (Page) GetPageVersionList(productCode string) ([]*model.PageVersionInfo, error) {
	return service.NewPage().QueryVersion(&repository.QueryPageVersion{
		ProductCode: productCode,
	})
}

func (Page) GetById(id string) (*model.PageInfo, error) {
	return service.NewPage().FindById(id)
}

func (Page) PortCount(id string) (int64, error) {
	return service.NewPage().PortCount(id)
}

// Modify
// 修改信息 不包含状态
func (Page) Modify(req *repository.ModifyPage) error {
	PageService := service.NewPage()
	return PageService.Modify(req)
}

// Modify
// 修改信息 不包含状态
func (Page) CheckData(portId string, args ...any) error {
	PageService := service.NewPage()
	list, err := PageService.Query(&repository.QueryPage{
		PortId: portId,
	})

	if err != nil {
		return err
	}

	// 对比名字
	if len(args) > 0 {
		name := args[0].(string)
		for _, Page := range list {
			if strings.HasPrefix(name, fmt.Sprintf("/%s", Page.Name)) {
				return fmt.Errorf("[GIN匹配]类型的匹配规则与页面[%s]冲突，请修改为[MUX匹配]", Page.Name)
			}
		}
	}

	for _, Page := range list {
		// 节点里面去查询所有的节点
		ok, err := service.NewRule().FindLikeName(Page.Name)
		if err != nil {
			return err
		}

		// 查询到有相似的直接返回错误
		if ok {
			return fmt.Errorf("[GIN匹配]类型的匹配规则中发现了与页面[%s]冲突的路由，请修改为[MUX匹配]", Page.Name)
		}
	}

	return nil
}

// Modify
// 修改信息 不包含状态
func (Page) ModifyByMap(id string, datas map[string]any) error {
	PageService := service.NewPage()
	return PageService.ModifyByMap(id, datas)
}

// Create
// 创建数据
func (Page) Create(req *repository.CreatePage) error {
	pageSrv := service.NewPage()

	// 通过名称查询是否有相同名称的系统
	info, err := pageSrv.FindByName(req.Name, req.PortId)
	if err != nil {
		return fmt.Errorf("通过[%s]查找页面信息失败 %s", req.Name, err.Error())
	}

	if info.ID != "" {
		if info.ProductCode != req.ProductCode {
			return fmt.Errorf("不同产品不能使用相同路由名称 [%s-%s]", info.Title, info.Name)
		}
	}

	// 判断端口号在当前系统是否已经被使用
	info, err = pageSrv.FindByProductCode(req.PortId, req.ProductCode)
	if err != nil {
		return fmt.Errorf("查找页面信息失败 %s", err.Error())
	}

	pageSrv.OpenTx()
	defer func() {
		if err != nil {
			pageSrv.Rollback()
			return
		}

		pageSrv.Commit()
	}()

	if info.ID != "" {
		// 从版本中查看是否有相同版本
		versionInfo, err := pageSrv.FindByVersion(req.PortId, req.ProductCode, req.Version)
		if err != nil {
			return fmt.Errorf("查找页面版本信息失败 %s", err.Error())
		}

		if versionInfo.Version != "" {
			return fmt.Errorf("版本[%s]已经存在，请勿重复添加", req.Version)
		}

		err = pageSrv.CreatePageVersion(&model.PageVersionBase{PortId: req.PortId, ProductCode: req.ProductCode, Version: req.Version, PagePath: req.PagePath})
		if err != nil {
			return fmt.Errorf("创建版本失败 %s", err.Error())
		}

		// 将版本更新到当前页面
		err = pageSrv.ModifyByMap(info.ID, map[string]any{"version": req.Version})
		if err != nil {
			return fmt.Errorf("更新页面当前使用版本失败 %s", err.Error())
		}

		return nil
	}

	// 创建数据
	err = pageSrv.Create(req)
	if err != nil {
		return fmt.Errorf("创建信息失败 %s", err.Error())
	}

	err = pageSrv.CreatePageVersion(&model.PageVersionBase{PortId: req.PortId, ProductCode: req.ProductCode, Version: req.Version, PagePath: req.PagePath})
	if err != nil {
		return fmt.Errorf("创建版本失败 %s", err.Error())
	}

	portSrv := service.NewPort(pageSrv.GetContext())
	err = portSrv.StepCount(req.PortId, service.ST_PAGE, service.STT_PLUS)
	if err != nil {
		return fmt.Errorf("更新页面统计失败 %s", err.Error())
	}

	return nil
}

// Del
// 根据ID删除信息
func (Page) Del(id string) error {
	PageService := service.NewPage()

	// 检查使用状态，如果是正在使用则不允许删除
	info, err := PageService.FindById(id)
	if err != nil {
		return err
	}

	if info.ID == "" {
		return errors.New("未找到需要删除的数据")
	}

	if info.State == 1 {
		return errors.New("当前页面使用中，不能删除")
	}

	PageService.OpenTx()
	defer func() {
		if err != nil {
			PageService.Rollback()
			return
		}

		PageService.Commit()
	}()

	err = PageService.Del(id)
	if err != nil {
		return err
	}

	err = PageService.DelAllVersion(info.ProductCode)
	if err != nil {
		return err
	}

	portService := service.NewPort(PageService.GetContext())
	err = portService.StepCount(info.PortId, service.ST_RULE, service.STT_REDUCE)
	return err
}
