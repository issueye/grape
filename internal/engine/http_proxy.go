package engine

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/issueye/grape/internal/global"
)

// http 代理
type HttpProxy struct {
	IsTSL bool // 是否需要证书
	Port  uint // 端口号
}

type TrafficStats struct {
	RequestBytes  int64
	ResponseBytes int64
}

// TrafficCountingReader 用于统计读取的字节数
type TrafficCountingReader struct {
	io.ReadCloser
	Stats *TrafficStats
}

func (r *TrafficCountingReader) Read(p []byte) (int, error) {
	n, err := r.ReadCloser.Read(p)
	r.Stats.RequestBytes += int64(n)
	r.Stats.ResponseBytes += int64(n)
	return n, err
}

func (r *TrafficCountingReader) Close() error {
	return r.ReadCloser.Close()
}

func getStatsFromContext(req *http.Request) *TrafficStats {
	statsIface := req.Context().Value("traffic_stats")
	if statsIface == nil {
		return &TrafficStats{}
	}
	stats, ok := statsIface.(*TrafficStats)
	if !ok {
		return &TrafficStats{}
	}
	return stats
}

// 定义反向代理处理器
func ReverseProxyHttpHandler(targetURL string) *httputil.ReverseProxy {
	// 解析目标URL
	target, err := url.Parse(targetURL)
	if err != nil {
		global.Log.Errorf("解析目标URL失败 %s", err.Error())
	}

	// 创建反向代理实例
	proxy := httputil.NewSingleHostReverseProxy(target)

	proxy.Director = func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.Host = target.Host

		// 创建 TrafficStats 对象
		stats := &TrafficStats{}
		// 将 TrafficStats 对象添加到请求上下文中
		ctx := context.WithValue(req.Context(), "traffic_stats", stats)
		req = req.WithContext(ctx)

		if req.Body != nil {
			req.Body = &TrafficCountingReader{
				ReadCloser: req.Body,
				Stats:      stats,
			}
		}
	}

	proxy.ModifyResponse = func(resp *http.Response) error {
		// 获取请求的 TrafficStats 对象
		stats := getStatsFromContext(resp.Request)

		// 使用 TrafficCountingReader 包装响应 Body
		resp.Body = &TrafficCountingReader{
			ReadCloser: resp.Body,
			Stats:      stats,
		}

		fmt.Printf("Forwarded %d bytes (request), %d bytes (response)\n", stats.RequestBytes, stats.ResponseBytes)
		return nil
	}

	return proxy
}
