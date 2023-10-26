package global

import (
	"net/http"
	"sync"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB         *gorm.DB
	Log        *zap.SugaredLogger
	Logger     *zap.Logger
	Router     *gin.Engine
	HttpServer *http.Server
	Auth       *jwt.GinJWTMiddleware

	// 存放 gin 对象
	GinEngines = new(sync.Map)
)
