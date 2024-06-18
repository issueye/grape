package engine

import (
	"bytes"
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

// 定义反向代理处理器
func ReverseProxyHttpHandler(targetURL string) *httputil.ReverseProxy {
	// 解析目标URL
	target, err := url.Parse(targetURL)
	if err != nil {
		global.Log.Errorf("解析目标URL失败 %s", err.Error())
	}

	// 创建反向代理实例
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ModifyResponse = modifyResponse

	return proxy
}

func modifyResponse(resp *http.Response) error {
	var (
		totalInBytes  int64
		totalOutBytes int64
	)

	if resp != nil && resp.Body != nil {
		// 读取响应主体的大小
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		totalOutBytes += int64(len(body))
		resp.Body = io.NopCloser(bytes.NewBuffer(body))
	}

	// 读取请求主体和头部的大小
	if resp != nil && resp.Request != nil {
		if resp.Request.Body != nil {
			// 请求主体
			body, err := io.ReadAll(resp.Request.Body)
			if err != nil {
				return err
			}
			totalInBytes += int64(len(body))
			resp.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		}

		// 请求头部
		reqDump, err := httputil.DumpRequest(resp.Request, false)
		if err != nil {
			return err
		}
		totalInBytes += int64(len(reqDump))
	}

	// 读取响应头部的大小
	respDump, err := httputil.DumpResponse(resp, false)
	if err != nil {
		return err
	}
	totalOutBytes += int64(len(respDump))

	fmt.Printf("Total in: %d bytes\n", totalInBytes)
	fmt.Printf("Total out: %d bytes\n", totalOutBytes)

	return nil
}
