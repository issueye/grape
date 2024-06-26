// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.1
// source: host.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// message Get
type ParamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Secens   string `protobuf:"bytes,1,opt,name=secens,proto3" json:"secens,omitempty"`
	Params   string `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	DefValue string `protobuf:"bytes,3,opt,name=defValue,proto3" json:"defValue,omitempty"`
}

func (x *ParamRequest) Reset() {
	*x = ParamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_host_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParamRequest) ProtoMessage() {}

func (x *ParamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_host_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParamRequest.ProtoReflect.Descriptor instead.
func (*ParamRequest) Descriptor() ([]byte, []int) {
	return file_host_proto_rawDescGZIP(), []int{0}
}

func (x *ParamRequest) GetSecens() string {
	if x != nil {
		return x.Secens
	}
	return ""
}

func (x *ParamRequest) GetParams() string {
	if x != nil {
		return x.Params
	}
	return ""
}

func (x *ParamRequest) GetDefValue() string {
	if x != nil {
		return x.DefValue
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_host_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_host_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_host_proto_rawDescGZIP(), []int{1}
}

type ResultString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ResultString) Reset() {
	*x = ResultString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_host_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultString) ProtoMessage() {}

func (x *ResultString) ProtoReflect() protoreflect.Message {
	mi := &file_host_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultString.ProtoReflect.Descriptor instead.
func (*ResultString) Descriptor() ([]byte, []int) {
	return file_host_proto_rawDescGZIP(), []int{2}
}

func (x *ResultString) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type InitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddServer uint32 `protobuf:"varint,1,opt,name=add_server,json=addServer,proto3" json:"add_server,omitempty"`
}

func (x *InitRequest) Reset() {
	*x = InitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_host_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitRequest) ProtoMessage() {}

func (x *InitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_host_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitRequest.ProtoReflect.Descriptor instead.
func (*InitRequest) Descriptor() ([]byte, []int) {
	return file_host_proto_rawDescGZIP(), []int{3}
}

func (x *InitRequest) GetAddServer() uint32 {
	if x != nil {
		return x.AddServer
	}
	return 0
}

type HeartbeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *HeartbeatResponse) Reset() {
	*x = HeartbeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_host_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatResponse) ProtoMessage() {}

func (x *HeartbeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_host_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatResponse.ProtoReflect.Descriptor instead.
func (*HeartbeatResponse) Descriptor() ([]byte, []int) {
	return file_host_proto_rawDescGZIP(), []int{4}
}

func (x *HeartbeatResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_host_proto protoreflect.FileDescriptor

var file_host_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x0c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x65, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x66, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x66, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x24, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2c,
	0x0a, 0x0b, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x64, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x61, 0x64, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x29, 0x0a, 0x11,
	0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x42, 0x0a, 0x0a, 0x48, 0x6f, 0x73, 0x74, 0x48,
	0x65, 0x6c, 0x70, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x6d, 0x0a, 0x0c, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x04, 0x49,
	0x6e, 0x69, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x69, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x33, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65,
	0x61, 0x74, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65,
	0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_host_proto_rawDescOnce sync.Once
	file_host_proto_rawDescData = file_host_proto_rawDesc
)

func file_host_proto_rawDescGZIP() []byte {
	file_host_proto_rawDescOnce.Do(func() {
		file_host_proto_rawDescData = protoimpl.X.CompressGZIP(file_host_proto_rawDescData)
	})
	return file_host_proto_rawDescData
}

var file_host_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_host_proto_goTypes = []interface{}{
	(*ParamRequest)(nil),      // 0: proto.ParamRequest
	(*Empty)(nil),             // 1: proto.Empty
	(*ResultString)(nil),      // 2: proto.ResultString
	(*InitRequest)(nil),       // 3: proto.InitRequest
	(*HeartbeatResponse)(nil), // 4: proto.HeartbeatResponse
}
var file_host_proto_depIdxs = []int32{
	0, // 0: proto.HostHelper.GetParam:input_type -> proto.ParamRequest
	3, // 1: proto.CommonHelper.Init:input_type -> proto.InitRequest
	1, // 2: proto.CommonHelper.Heartbeat:input_type -> proto.Empty
	2, // 3: proto.HostHelper.GetParam:output_type -> proto.ResultString
	1, // 4: proto.CommonHelper.Init:output_type -> proto.Empty
	4, // 5: proto.CommonHelper.Heartbeat:output_type -> proto.HeartbeatResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_host_proto_init() }
func file_host_proto_init() {
	if File_host_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_host_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParamRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_host_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_host_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultString); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_host_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_host_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_host_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_host_proto_goTypes,
		DependencyIndexes: file_host_proto_depIdxs,
		MessageInfos:      file_host_proto_msgTypes,
	}.Build()
	File_host_proto = out.File
	file_host_proto_rawDesc = nil
	file_host_proto_goTypes = nil
	file_host_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HostHelperClient is the client API for HostHelper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HostHelperClient interface {
	GetParam(ctx context.Context, in *ParamRequest, opts ...grpc.CallOption) (*ResultString, error)
}

type hostHelperClient struct {
	cc grpc.ClientConnInterface
}

func NewHostHelperClient(cc grpc.ClientConnInterface) HostHelperClient {
	return &hostHelperClient{cc}
}

func (c *hostHelperClient) GetParam(ctx context.Context, in *ParamRequest, opts ...grpc.CallOption) (*ResultString, error) {
	out := new(ResultString)
	err := c.cc.Invoke(ctx, "/proto.HostHelper/GetParam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HostHelperServer is the server API for HostHelper service.
type HostHelperServer interface {
	GetParam(context.Context, *ParamRequest) (*ResultString, error)
}

// UnimplementedHostHelperServer can be embedded to have forward compatible implementations.
type UnimplementedHostHelperServer struct {
}

func (*UnimplementedHostHelperServer) GetParam(context.Context, *ParamRequest) (*ResultString, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParam not implemented")
}

func RegisterHostHelperServer(s *grpc.Server, srv HostHelperServer) {
	s.RegisterService(&_HostHelper_serviceDesc, srv)
}

func _HostHelper_GetParam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostHelperServer).GetParam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.HostHelper/GetParam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostHelperServer).GetParam(ctx, req.(*ParamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HostHelper_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.HostHelper",
	HandlerType: (*HostHelperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetParam",
			Handler:    _HostHelper_GetParam_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "host.proto",
}

// CommonHelperClient is the client API for CommonHelper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommonHelperClient interface {
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*Empty, error)
	Heartbeat(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HeartbeatResponse, error)
}

type commonHelperClient struct {
	cc grpc.ClientConnInterface
}

func NewCommonHelperClient(cc grpc.ClientConnInterface) CommonHelperClient {
	return &commonHelperClient{cc}
}

func (c *commonHelperClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.CommonHelper/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonHelperClient) Heartbeat(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HeartbeatResponse, error) {
	out := new(HeartbeatResponse)
	err := c.cc.Invoke(ctx, "/proto.CommonHelper/Heartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommonHelperServer is the server API for CommonHelper service.
type CommonHelperServer interface {
	Init(context.Context, *InitRequest) (*Empty, error)
	Heartbeat(context.Context, *Empty) (*HeartbeatResponse, error)
}

// UnimplementedCommonHelperServer can be embedded to have forward compatible implementations.
type UnimplementedCommonHelperServer struct {
}

func (*UnimplementedCommonHelperServer) Init(context.Context, *InitRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (*UnimplementedCommonHelperServer) Heartbeat(context.Context, *Empty) (*HeartbeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}

func RegisterCommonHelperServer(s *grpc.Server, srv CommonHelperServer) {
	s.RegisterService(&_CommonHelper_serviceDesc, srv)
}

func _CommonHelper_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonHelperServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CommonHelper/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonHelperServer).Init(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommonHelper_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonHelperServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CommonHelper/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonHelperServer).Heartbeat(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommonHelper_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CommonHelper",
	HandlerType: (*CommonHelperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _CommonHelper_Init_Handler,
		},
		{
			MethodName: "Heartbeat",
			Handler:    _CommonHelper_Heartbeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "host.proto",
}
