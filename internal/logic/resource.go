package logic

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/internal/global"
	"github.com/issueye/grape/internal/repository"
	"github.com/issueye/grape/internal/service"
	"github.com/issueye/grape/pkg/utils"
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

// Upload
// 上传文件
func (Resource) Upload(data *repository.UploadData) error {

	ext := filepath.Ext(data.UploadKey.Filename)
	filename := strings.TrimSuffix(path.Base(data.UploadKey.Filename), ext)
	// 生成一个sha字符串
	filename = utils.Sha256_2Str(fmt.Sprintf("%s-%s", filename, utils.GetUUID()))

	// 获取文件名，并创建新的文件存储
	path := filepath.Join(global.GetResourceRootPath(ext), fmt.Sprintf("%s%s", filename, ext))
	// 如果是页面的就按照页面的逻辑进行处理
	if data.Type == "page" {
		// 获取文件名，并创建新的文件存储
		path = filepath.Join(global.GetResourcePagePath(ext), fmt.Sprintf("%s%s", filename, ext))
	}

	// 创建上传文件
	out, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("创建文件失败 %s", err.Error())
	}

	src, err := data.UploadKey.Open()
	if err != nil {
		return fmt.Errorf("打开上传的文件失败 %s", err.Error())
	}

	defer out.Close()
	//将读取的文件流写到文件中
	_, err = io.Copy(out, src)
	if err != nil {
		return fmt.Errorf("读取失败 %s", err.Error())
	}

	return nil
}
