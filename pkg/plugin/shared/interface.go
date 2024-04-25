package shared

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"github.com/issueye/grape/pkg/plugin/shared/proto/proto"
	"google.golang.org/grpc"
)

var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "YH_LINE_PLUGIN",
	MagicCookieValue: "host",
}

var PluginMap = map[string]plugin.Plugin{
	"common": &CommonPlugin{},
}

type CommonPlugin struct {
	plugin.NetRPCUnsupportedPlugin

	Impl CommonHelper
}

func (cp *CommonPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterCommonHelperServer(s, &GRPCServer{
		Impl:   cp.Impl,
		broker: broker,
	})
	return nil
}

func (cp *CommonPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{
		client: proto.NewCommonHelperClient(c),
		broker: broker,
	}, nil
}

var _ plugin.GRPCPlugin = &CommonPlugin{}

// 宿主程序接口
type IHostHelper interface {
	IConfig
}

// 配置信息
type IConfig interface {
	GetParam(secens string, params string, defValue string) (string, error) // 读取配置参数
	// SetParam(secens string, params string, value string, explanation string) error // 添加\修改配置参数
}

type CommonHelper interface {
	Init(IConfig) error
	Heartbeat() (string, error)
}
