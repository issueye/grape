package repository

// http 报文查询请求结构体
type HttpRequest struct {
	Url    string            `json:"url"`
	Method string            `json:"method"`
	Header map[string]string `json:"header"`
	Body   string            `json:"body"`
}

// http 报文查询响应结构体
type HttpResponse struct {
	StatusCode int               `json:"statusCode"`
	Header     map[string]string `json:"header"`
	Body       string            `json:"body"`
}
