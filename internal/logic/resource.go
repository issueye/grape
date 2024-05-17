package logic

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/internal/global"
	"github.com/issueye/grape/internal/repository"
	"github.com/issueye/grape/internal/service"
	"github.com/issueye/grape/pkg/utils"
)

var Rsse = make(map[string]*ResourceSSE)

type ResourceSSE struct {
	Id string `json:"id"` // 编码
}

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
func (Resource) Upload(data *repository.UploadData) (map[string]string, error) {

	ext := filepath.Ext(data.UploadKey.Filename)
	filename := strings.TrimSuffix(path.Base(data.UploadKey.Filename), ext)
	// 生成一个sha字符串
	filename = utils.Sha256_2Str(fmt.Sprintf("%s-%s", filename, utils.GetUUID()))

	Rsse[filename] = &ResourceSSE{Id: filename}

	go func() {
		Resource{}.fileParse(data, filename, ext)
	}()

	return map[string]string{
		"name": filename,
		"ext":  ext,
	}, nil
}

type Progress struct {
	Code     int    `json:"code"`     // 状态码
	Message  string `json:"message"`  // 消息
	Progress int    `json:"progress"` // 进度
}

func getProgress(code int, progress int, msg string) string {
	data := &Progress{
		Code:     code,
		Progress: progress,
		Message:  msg,
	}

	return utils.Struct2Json(data)
}

type Send struct {
	event string // 编码
}

func NewSend(event string) *Send {
	return &Send{
		event: event,
	}
}

func (s *Send) Success(progress int, msg string) {
	global.SSE.SendEventMessage(getProgress(1, progress, msg), s.event, "")
	time.Sleep(1 * time.Second)
}

func (s *Send) Fail(progress int, msg string) {
	global.SSE.SendEventMessage(getProgress(0, progress, msg), s.event, "")
}

func (Resource) fileParse(data *repository.UploadData, filename string, ext string) {
	send := NewSend(data.Id)

	path := filepath.Join(global.GetResourcePathByType(ext), fmt.Sprintf("%s%s", filename, ext))
	send.Success(30, "开始创建文件")

	// 创建上传文件
	out, err := os.Create(path)
	if err != nil {
		global.Log.Errorf("创建文件失败 %s", err.Error())
		send.Fail(30, "创建文件失败...")
		return
	}

	global.Log.Info("创建文件副本成功")
	send.Success(40, "创建文件副本成功...")

	src, err := data.UploadKey.Open()
	if err != nil {
		global.Log.Errorf("打开上传的文件失败 %s", err.Error())
		send.Fail(40, "打开上传的文件失败...")
		return
	}

	global.Log.Info("打开上传的文件成功")
	send.Success(60, "打开上传的文件成功...")

	global.Log.Info("开始拷贝文件流")
	send.Success(70, "开始拷贝文件流...")

	defer out.Close()
	//将读取的文件流写到文件中
	_, err = io.Copy(out, src)
	if err != nil {
		global.Log.Errorf("读取失败 %s", err.Error())
		send.Fail(70, "读取失败...")
		return
	}

	global.Log.Info("完成文件流拷贝")
	send.Success(75, "完成文件流拷贝...")

	switch strings.ToLower(data.Type) {
	case "page":
		{
			send.Success(77, "开始解压文件...")

			tempPath := global.GetTempPath()
			err = utils.Unzip(path, tempPath)
			if err != nil {
				global.Log.Errorf("解压文件失败 %s", err.Error())
				send.Fail(77, "解压文件失败...")
				return
			}

			configFile := filepath.Join(tempPath, "pageConfig.toml")
			info, err := os.ReadFile(configFile)
			if err != nil {
				global.Log.Errorf("读取页面配置信息失败 %s", err.Error())
				send.Fail(80, "读取页面配置信息失败...")
				return
			}

			fmt.Println("info", string(info))
		}
	}

	send.Success(100, "上传成功")
}
