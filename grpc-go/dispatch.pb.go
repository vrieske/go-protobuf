// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.6
// source: protobuf/dispatch.proto

package grpc_go

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

type DispatchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Ip     string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *DispatchReq) Reset() {
	*x = DispatchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_dispatch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DispatchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DispatchReq) ProtoMessage() {}

func (x *DispatchReq) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_dispatch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DispatchReq.ProtoReflect.Descriptor instead.
func (*DispatchReq) Descriptor() ([]byte, []int) {
	return file_protobuf_dispatch_proto_rawDescGZIP(), []int{0}
}

func (x *DispatchReq) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *DispatchReq) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type DispatchRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ips     []string `protobuf:"bytes,1,rep,name=ips,proto3" json:"ips,omitempty"`
	Cname   string   `protobuf:"bytes,2,opt,name=cname,proto3" json:"cname,omitempty"`
	Message string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DispatchRes) Reset() {
	*x = DispatchRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_dispatch_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DispatchRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DispatchRes) ProtoMessage() {}

func (x *DispatchRes) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_dispatch_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DispatchRes.ProtoReflect.Descriptor instead.
func (*DispatchRes) Descriptor() ([]byte, []int) {
	return file_protobuf_dispatch_proto_rawDescGZIP(), []int{1}
}

func (x *DispatchRes) GetIps() []string {
	if x != nil {
		return x.Ips
	}
	return nil
}

func (x *DispatchRes) GetCname() string {
	if x != nil {
		return x.Cname
	}
	return ""
}

func (x *DispatchRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_protobuf_dispatch_proto protoreflect.FileDescriptor

var file_protobuf_dispatch_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x69, 0x73, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x64, 0x69, 0x73, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x22, 0x35, 0x0a, 0x0b, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x4f, 0x0a, 0x0b, 0x44, 0x69,
	0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x70, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x70, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x47, 0x0a, 0x08, 0x44,
	0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x12, 0x3b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x44, 0x69,
	0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x12, 0x15, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63,
	0x68, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e,
	0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2d, 0x67, 0x6f, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_dispatch_proto_rawDescOnce sync.Once
	file_protobuf_dispatch_proto_rawDescData = file_protobuf_dispatch_proto_rawDesc
)

func file_protobuf_dispatch_proto_rawDescGZIP() []byte {
	file_protobuf_dispatch_proto_rawDescOnce.Do(func() {
		file_protobuf_dispatch_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_dispatch_proto_rawDescData)
	})
	return file_protobuf_dispatch_proto_rawDescData
}

var file_protobuf_dispatch_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protobuf_dispatch_proto_goTypes = []interface{}{
	(*DispatchReq)(nil), // 0: dispatch.DispatchReq
	(*DispatchRes)(nil), // 1: dispatch.DispatchRes
}
var file_protobuf_dispatch_proto_depIdxs = []int32{
	0, // 0: dispatch.Dispatch.GetDispatch:input_type -> dispatch.DispatchReq
	1, // 1: dispatch.Dispatch.GetDispatch:output_type -> dispatch.DispatchRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_dispatch_proto_init() }
func file_protobuf_dispatch_proto_init() {
	if File_protobuf_dispatch_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_dispatch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DispatchReq); i {
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
		file_protobuf_dispatch_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DispatchRes); i {
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
			RawDescriptor: file_protobuf_dispatch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_dispatch_proto_goTypes,
		DependencyIndexes: file_protobuf_dispatch_proto_depIdxs,
		MessageInfos:      file_protobuf_dispatch_proto_msgTypes,
	}.Build()
	File_protobuf_dispatch_proto = out.File
	file_protobuf_dispatch_proto_rawDesc = nil
	file_protobuf_dispatch_proto_goTypes = nil
	file_protobuf_dispatch_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DispatchClient is the client API for Dispatch service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DispatchClient interface {
	GetDispatch(ctx context.Context, in *DispatchReq, opts ...grpc.CallOption) (*DispatchRes, error)
}

type dispatchClient struct {
	cc grpc.ClientConnInterface
}

func NewDispatchClient(cc grpc.ClientConnInterface) DispatchClient {
	return &dispatchClient{cc}
}

func (c *dispatchClient) GetDispatch(ctx context.Context, in *DispatchReq, opts ...grpc.CallOption) (*DispatchRes, error) {
	out := new(DispatchRes)
	err := c.cc.Invoke(ctx, "/dispatch.Dispatch/GetDispatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DispatchServer is the server API for Dispatch service.
type DispatchServer interface {
	GetDispatch(context.Context, *DispatchReq) (*DispatchRes, error)
}

// UnimplementedDispatchServer can be embedded to have forward compatible implementations.
type UnimplementedDispatchServer struct {
}

func (*UnimplementedDispatchServer) GetDispatch(context.Context, *DispatchReq) (*DispatchRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDispatch not implemented")
}

func RegisterDispatchServer(s *grpc.Server, srv DispatchServer) {
	s.RegisterService(&_Dispatch_serviceDesc, srv)
}

func _Dispatch_GetDispatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DispatchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatchServer).GetDispatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dispatch.Dispatch/GetDispatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatchServer).GetDispatch(ctx, req.(*DispatchReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dispatch_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dispatch.Dispatch",
	HandlerType: (*DispatchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDispatch",
			Handler:    _Dispatch_GetDispatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/dispatch.proto",
}
