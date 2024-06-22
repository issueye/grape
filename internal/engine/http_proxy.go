package engine

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/issueye/grape/internal/common/model"
	"github.com/issueye/grape/internal/global"
)

// http 代理
type HttpProxy struct {
	IsTSL bool // 是否需要证书
	Port  uint // 端口号
}

type myRoundTripper struct {
	http.RoundTripper
}

func (m *myRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	traffic := model.NewTrafficStatistics()

	headerSize := int64(0)
	url := req.URL.Path
	if req.URL.RawQuery != "" {
		url += fmt.Sprintf("?%s", req.URL.RawQuery)
	}

	traffic.Request.Path = url
	traffic.Request.Method = req.Method

	for key, value := range req.Header {
		headerSize += int64(len(fmt.Sprintf("%s: %s\r\n", key, value)))
		traffic.Request.Header[key] = value
	}

	traffic.Request.InHeaderBytes = headerSize

	if req.Body != nil {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			global.Log.Errorf("读取请求内容失败 %s", err.Error())
			return m.RoundTripper.RoundTrip(req)
		}

		bodyBuf := bytes.NewBuffer(body)

		traffic.Request.InBodyBytes = int64(bodyBuf.Len())
		traffic.Request.Body = bodyBuf.String()

		req.Body = io.NopCloser(bytes.NewBuffer(body))
	}

	ctx := req.Context()
	ctx = context.WithValue(ctx, "traffic", &traffic)
	req = req.WithContext(ctx)

	resp, err := m.RoundTripper.RoundTrip(req)
	if err != nil {
		global.Log.Errorf("转发失败 %s", err.Error())
	}

	return resp, err
}

// 定义反向代理处理器
func ReverseProxyHttpHandler(targetURL string) *httputil.ReverseProxy {
	// 解析目标URL
	target, err := url.Parse(targetURL)
	if err != nil {
		global.Log.Errorf("解析目标URL失败 %s", err.Error())
	}

	// 创建反向代理实例
	proxy := &httputil.ReverseProxy{
		Director:       director(target),
		ModifyResponse: modifyResponse,
		Transport:      &myRoundTripper{http.DefaultTransport},
	}

	// proxy := httputil.NewSingleHostReverseProxy(target)
	// proxy.Director = director(target)
	// proxy.ModifyResponse = modifyResponse

	return proxy
}

func director(target *url.URL) func(*http.Request) {
	return func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
	}
}

func modifyResponse(resp *http.Response) error {
	var (
		traffic *model.TrafficStatistics
		ok      bool
	)

	if resp != nil {
		// 读取请求主体和头部的大小
		ctx := resp.Request.Context()
		data := ctx.Value("traffic")
		if data != nil {
			traffic, ok = data.(*model.TrafficStatistics)
			if !ok {
				traffic = model.NewTrafficStatistics()
			}
		}

		for key, value := range resp.Header {
			traffic.Response.Header[key] = value
		}

		// 读取响应头部的大小
		respDump, err := httputil.DumpResponse(resp, false)
		if err != nil {
			return err
		}
		traffic.Response.OutHeaderBytes = int64(len(respDump))

		if resp.Body != nil {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				global.Log.Errorf("读取返回内容 %s", err.Error())
				return err
			}
			bodyBuf := bytes.NewBuffer(body)
			traffic.Response.Body = bodyBuf.String()
			traffic.Response.OutBodyBytes = int64(bodyBuf.Len())

			resp.Body = io.NopCloser(bytes.NewBuffer(body))
		}

		// 将数据通过管道的方式传入 global.Index
		global.IndexDB <- traffic
		return nil
	}

	return nil
}
