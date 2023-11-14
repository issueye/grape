package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// trafficMiddleware 是自定义的中间件函数，用于统计流量
func TrafficMiddleware(log *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求的入流量
		start := time.Now()

		// 创建一个 ResponseWriter 的代理，用于捕获响应的出流量
		w := &responseWriter{
			ResponseWriter: c.Writer,
			body:           make([]byte, 0),
		}

		// 设置代理，保证后续处理链路中使用的是代理的 ResponseWriter
		c.Writer = w

		// 继续处理链路
		c.Next()

		// 计算请求的入流量和请求头的大小
		elapsed := time.Since(start)
		requestSize := calculateRequestSize(c.Request)

		size := w.Size()
		headerSize := w.HeaderSize()

		reqIP := c.ClientIP()
		if reqIP == "::1" {
			reqIP = "127.0.0.1"
		}

		// 打印流量信息
		log.Infof("\nIP:%s, 请求: %s %s, code: %d, 请求Header: %d bytes, 请求Body: %d bytes, 总流量: %d bytes, 响应Header: %d, 响应Body: %d, 响应总流量：%d, 响应时间: %s",
			reqIP,
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			requestSize.HeaderSize,
			requestSize.BodySize,
			requestSize.TotalSize,
			headerSize,
			size,
			size+headerSize,
			elapsed,
		)
	}
}

// RequestSize 包含请求大小和请求头大小的结构体
type RequestSize struct {
	TotalSize  int // 请求大小（包括请求体和请求头）
	BodySize   int // 请求体大小
	HeaderSize int // 请求头大小
}

// calculateRequestSize 计算请求的大小和请求头的大小
func calculateRequestSize(r *http.Request) RequestSize {
	// 序列化请求头
	headerBytes, err := json.Marshal(r.Header)
	if err != nil {
		log.Println("序列化请求头错误:", err)
	}

	// 计算请求大小和请求头大小
	bodySize := int(r.ContentLength)
	totalSize := len(headerBytes) + int(r.ContentLength)
	headerSize := len(headerBytes)

	return RequestSize{
		TotalSize:  totalSize,
		HeaderSize: headerSize,
		BodySize:   bodySize,
	}
}

// responseWriter 是 ResponseWriter 的代理，用于捕获响应的出流量
type responseWriter struct {
	gin.ResponseWriter
	body []byte
}

// Write 是实现 ResponseWriter 接口的方法，用于捕获响应的出流量
func (w *responseWriter) Write(b []byte) (int, error) {
	w.body = append(w.body, b...)
	return w.ResponseWriter.Write(b)
}

// Size 返回响应的大小
func (w *responseWriter) Size() int {
	return len(w.body)
}

func (w *responseWriter) HeaderSize() int {
	headerSize := 0
	for key, values := range w.ResponseWriter.Header() {
		for _, value := range values {
			headerSize += len(key) + len(value)
		}
	}
	return headerSize
}
