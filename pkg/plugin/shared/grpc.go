package shared

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"github.com/issueye/grape/pkg/plugin/shared/proto/proto"
	"google.golang.org/grpc"
)

// 宿主程序
type GRPCClient struct {
	broker  *plugin.GRPCBroker       // broker
	GServer *grpc.Server             // GRPC 服务
	client  proto.CommonHelperClient // 公共接口客户端
}

// 初始化
func (m *GRPCClient) Init(cfg IConfig) error {
	helperServer := &GRPCHostHelperServer{Impl: cfg}

	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		s := grpc.NewServer(opts...)
		proto.RegisterHostHelperServer(s, helperServer)
		m.GServer = s
		return s
	}

	brokerID := m.broker.NextId()
	go m.broker.AcceptAndServe(brokerID, serverFunc)

	_, err := m.client.Init(context.Background(), &proto.InitRequest{
		AddServer: brokerID,
	})
	if err != nil {
		return err
	}

	return err
}

func (m *GRPCClient) Heartbeat() (string, error) {
	resp, err := m.client.Heartbeat(context.Background(), &proto.Empty{})
	if err != nil {
		return "", err
	}

	return resp.Value, nil
}

// 插件程序 SERVER
type GRPCServer struct {
	Impl   CommonHelper
	broker *plugin.GRPCBroker
	conn   *grpc.ClientConn
}

func (m *GRPCServer) Init(ctx context.Context, req *proto.InitRequest) (*proto.Empty, error) {
	conn, err := m.broker.Dial(req.AddServer)
	if err != nil {
		return nil, err
	}

	m.conn = conn

	host := &GRPCHostHelperClient{proto.NewHostHelperClient(conn)}
	return &proto.Empty{}, m.Impl.Init(host)
}

func (m *GRPCServer) Heartbeat(ctx context.Context, req *proto.Empty) (*proto.HeartbeatResponse, error) {
	result, err := m.Impl.Heartbeat()
	if err != nil {
		return nil, err
	}

	return &proto.HeartbeatResponse{
		Value: result,
	}, nil
}

// ----------------------------------hostHelper------------------------------------------------

type GRPCHostHelperClient struct{ client proto.HostHelperClient }

func (m *GRPCHostHelperClient) GetParam(secens string, params string, defValue string) (string, error) {
	resp, err := m.client.GetParam(context.Background(), &proto.ParamRequest{
		Secens:   secens,
		Params:   params,
		DefValue: defValue,
	})

	if err != nil {
		return "", err
	}

	return resp.Value, nil
}

type GRPCHostHelperServer struct{ Impl IHostHelper }

func (m *GRPCHostHelperServer) GetParam(ctx context.Context, req *proto.ParamRequest) (resp *proto.ResultString, err error) {
	result, err := m.Impl.GetParam(req.Secens, req.Params, req.DefValue)
	if err != nil {
		return nil, err
	}

	return &proto.ResultString{Value: result}, nil
}
