package engine

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

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
	traffic := TrafficStatistics{}

	headerSize := int64(0)
	traffic.InHttpMessages = append(traffic.InHttpMessages, "\n================请求报文==================\n")
	url := req.URL.Path
	if req.URL.RawQuery != "" {
		url += fmt.Sprintf("?%s", req.URL.RawQuery)
	}
	traffic.InHttpMessages = append(traffic.InHttpMessages, fmt.Sprintf("%s\n", url))
	traffic.InHttpMessages = append(traffic.InHttpMessages, fmt.Sprintf("%s\n", req.Method))
	for key, value := range req.Header {
		headerSize += int64(len(fmt.Sprintf("%s: %s\r\n", key, value)))
		traffic.InHttpMessages = append(traffic.InHttpMessages, fmt.Sprintf("%s: %s\n", key, value))
	}

	traffic.InHttpMessages = append(traffic.InHttpMessages, "\r\n")

	traffic.InHeaderBytes = headerSize

	if req.Body != nil {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			global.Log.Errorf("读取请求内容失败 %s", err.Error())
			return m.RoundTripper.RoundTrip(req)
		}

		bodyBuf := bytes.NewBuffer(body)

		traffic.InBodyBytes = int64(bodyBuf.Len())
		traffic.InHttpMessages = append(traffic.InHttpMessages, bodyBuf.String())

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

type TrafficStatistics struct {
	// 请求信息
	InHeaderBytes  int64
	InBodyBytes    int64
	InHttpMessages []string

	// 返回信息
	OutHeaderBytes  int64
	OutBodyBytes    int64
	OutHttpMessages []string
}

func director(target *url.URL) func(*http.Request) {
	return func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
	}
}

func modifyResponse(resp *http.Response) error {
	var (
		// totalInBytes  int64
		// totalOutBytes int64
		traffic *TrafficStatistics
		ok      bool
	)

	if resp != nil {
		// 读取请求主体和头部的大小
		ctx := resp.Request.Context()
		data := ctx.Value("traffic")
		if data != nil {
			traffic, ok = data.(*TrafficStatistics)
			if ok {
				// totalInBytes = traffic.InBodyBytes + traffic.InHeaderBytes
				fmt.Println(strings.Join(traffic.InHttpMessages, ""))

			} else {
				traffic = &TrafficStatistics{
					InBodyBytes:   0,
					InHeaderBytes: 0,
				}
			}
		}

		str := make([]string, 0)
		str = append(str, "\n================响应报文==================\n")
		// headerSize := int64(len([]byte(resp.Proto + "\r\n")))
		for key, value := range resp.Header {
			str = append(str, fmt.Sprintf("%s: %s\n", key, value))
			// headerSize += int64(len(fmt.Sprintf("%s: %s\r\n", key, value)))
		}

		// 读取响应头部的大小
		respDump, err := httputil.DumpResponse(resp, false)
		if err != nil {
			return err
		}
		// totalOutBytes += int64(len(respDump))

		str = append(str, "\r\n")

		// headerSize += int64(len("\r\n"))
		traffic.OutHeaderBytes = int64(len(respDump))
		// totalOutBytes += headerSize

		if resp.Body != nil {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				global.Log.Errorf("读取返回内容 %s", err.Error())
				return err
			}
			bodyBuf := bytes.NewBuffer(body)
			str = append(str, bodyBuf.String())
			// totalOutBytes += int64(bodyBuf.Len())
			traffic.OutBodyBytes = int64(bodyBuf.Len())

			resp.Body = io.NopCloser(bytes.NewBuffer(body))
		}

		fmt.Println(strings.Join(str, ""))

		fmt.Printf("in header %d bytes\n", traffic.InHeaderBytes)
		fmt.Printf("in body %d bytes\n", traffic.InBodyBytes)
		fmt.Printf("out header %d bytes\n", traffic.OutHeaderBytes)
		fmt.Printf("out body %d bytes\n", traffic.OutBodyBytes)
	}

	fmt.Printf("Total in: %d bytes\n", traffic.InBodyBytes+traffic.InHeaderBytes)
	fmt.Printf("Total out: %d bytes\n", traffic.OutBodyBytes+traffic.OutHeaderBytes)
	return nil
}
