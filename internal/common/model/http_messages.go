package model

type TrafficStatistics struct {
	// 请求信息
	// InHeaderBytes  int64
	// InBodyBytes    int64
	// InHttpMessages []string

	// 响应信息
	Request  HttpRequest
	Response HttpResponse

	// 返回信息
	// OutHeaderBytes  int64
	// OutBodyBytes    int64
	// OutHttpMessages []string
}

func NewTrafficStatistics() *TrafficStatistics {
	return &TrafficStatistics{
		Request: HttpRequest{
			Header: make(map[string][]string),
		},
		Response: HttpResponse{
			Header: make(map[string][]string),
		},
	}
}

type HttpRequest struct {
	Method        string
	Path          string
	Query         string
	Header        map[string][]string
	Body          string
	InHeaderBytes int64
	InBodyBytes   int64
}

type HttpResponse struct {
	Header         map[string][]string
	Body           string
	StatusCode     int
	OutHeaderBytes int64
	OutBodyBytes   int64
}
