package global

import (
	"embed"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gopkg.in/antage/eventsource.v1"
	"gorm.io/gorm"
)

var (
	DB         *gorm.DB
	Log        *zap.SugaredLogger
	Logger     *zap.Logger
	Router     *gin.Engine
	HttpServer *http.Server
	Auth       *jwt.GinJWTMiddleware
	SSE        eventsource.EventSource
	PageStatic embed.FS
)

type ActionType int

const (
	AT_START  ActionType = iota // 启动
	AT_STOP                     // 停用
	AT_RELOAD                   // 重载
)

type Port struct {
	Id     string
	Port   int
	Action ActionType
}

var (
	PortChan = make(chan *Port, 10) // 创建一个管道
)
