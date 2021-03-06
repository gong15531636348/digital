// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getArea.proto

package message

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Requset struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Requset) Reset()         { *m = Requset{} }
func (m *Requset) String() string { return proto.CompactTextString(m) }
func (*Requset) ProtoMessage()    {}
func (*Requset) Descriptor() ([]byte, []int) {
	return fileDescriptor_420050f84ccd68e8, []int{0}
}

func (m *Requset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Requset.Unmarshal(m, b)
}
func (m *Requset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Requset.Marshal(b, m, deterministic)
}
func (m *Requset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Requset.Merge(m, src)
}
func (m *Requset) XXX_Size() int {
	return xxx_messageInfo_Requset.Size(m)
}
func (m *Requset) XXX_DiscardUnknown() {
	xxx_messageInfo_Requset.DiscardUnknown(m)
}

var xxx_messageInfo_Requset proto.InternalMessageInfo

type Response struct {
	Errno                string   `protobuf:"bytes,1,opt,name=errno,proto3" json:"errno,omitempty"`
	Errmsg               string   `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	Data                 []*Area  `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_420050f84ccd68e8, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *Response) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *Response) GetData() []*Area {
	if m != nil {
		return m.Data
	}
	return nil
}

type Area struct {
	Aid                  int32    `protobuf:"varint,1,opt,name=aid,proto3" json:"aid,omitempty"`
	Aname                string   `protobuf:"bytes,2,opt,name=aname,proto3" json:"aname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Area) Reset()         { *m = Area{} }
func (m *Area) String() string { return proto.CompactTextString(m) }
func (*Area) ProtoMessage()    {}
func (*Area) Descriptor() ([]byte, []int) {
	return fileDescriptor_420050f84ccd68e8, []int{2}
}

func (m *Area) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Area.Unmarshal(m, b)
}
func (m *Area) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Area.Marshal(b, m, deterministic)
}
func (m *Area) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Area.Merge(m, src)
}
func (m *Area) XXX_Size() int {
	return xxx_messageInfo_Area.Size(m)
}
func (m *Area) XXX_DiscardUnknown() {
	xxx_messageInfo_Area.DiscardUnknown(m)
}

var xxx_messageInfo_Area proto.InternalMessageInfo

func (m *Area) GetAid() int32 {
	if m != nil {
		return m.Aid
	}
	return 0
}

func (m *Area) GetAname() string {
	if m != nil {
		return m.Aname
	}
	return ""
}

func init() {
	proto.RegisterType((*Requset)(nil), "message.Requset")
	proto.RegisterType((*Response)(nil), "message.Response")
	proto.RegisterType((*Area)(nil), "message.Area")
}

func init() { proto.RegisterFile("getArea.proto", fileDescriptor_420050f84ccd68e8) }

var fileDescriptor_420050f84ccd68e8 = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x3b, 0x8b, 0xc3, 0x30,
	0x10, 0x84, 0xf1, 0xf9, 0x75, 0xde, 0xc3, 0xe0, 0x13, 0xc7, 0x61, 0x52, 0x39, 0xae, 0x5c, 0xa9,
	0x70, 0xda, 0x34, 0xa9, 0xd2, 0x2b, 0x65, 0xaa, 0x0d, 0x5e, 0x4c, 0x0a, 0x5b, 0xce, 0x4a, 0xf9,
	0xff, 0x41, 0x0f, 0x42, 0x3a, 0x7d, 0x23, 0xf8, 0x66, 0x16, 0xea, 0x99, 0xec, 0x89, 0x09, 0xe5,
	0xc6, 0xda, 0x6a, 0x51, 0x2e, 0x64, 0x0c, 0xce, 0xd4, 0x57, 0x50, 0x2a, 0x7a, 0x3c, 0x0d, 0xd9,
	0xfe, 0x0a, 0xdf, 0x8a, 0xcc, 0xa6, 0x57, 0x43, 0xe2, 0x0f, 0x72, 0x62, 0x5e, 0x75, 0x9b, 0x74,
	0xc9, 0x50, 0xa9, 0x00, 0xe2, 0x1f, 0x0a, 0x62, 0x5e, 0xcc, 0xdc, 0x7e, 0xf9, 0x38, 0x92, 0xd8,
	0x43, 0x36, 0xa1, 0xc5, 0x36, 0xed, 0xd2, 0xe1, 0x67, 0xac, 0x65, 0x94, 0x4b, 0x57, 0xa8, 0xfc,
	0x57, 0x2f, 0x21, 0x73, 0x24, 0x1a, 0x48, 0xf1, 0x3e, 0x79, 0x6d, 0xae, 0xdc, 0xd3, 0x55, 0xe1,
	0x8a, 0x0b, 0x45, 0x67, 0x80, 0xf1, 0x08, 0x70, 0x0e, 0x8b, 0x2f, 0xc4, 0x42, 0x42, 0x19, 0x49,
	0x34, 0x6f, 0x7b, 0xdc, 0xbd, 0xfb, 0xfd, 0x48, 0xc2, 0xfc, 0x5b, 0xe1, 0xaf, 0x3c, 0xbc, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xb3, 0x45, 0x20, 0x56, 0xf6, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GetAreaSerClient is the client API for GetAreaSer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GetAreaSerClient interface {
	GetArea(ctx context.Context, in *Requset, opts ...grpc.CallOption) (*Response, error)
}

type getAreaSerClient struct {
	cc *grpc.ClientConn
}

func NewGetAreaSerClient(cc *grpc.ClientConn) GetAreaSerClient {
	return &getAreaSerClient{cc}
}

func (c *getAreaSerClient) GetArea(ctx context.Context, in *Requset, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/message.GetAreaSer/GetArea", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetAreaSerServer is the server API for GetAreaSer service.
type GetAreaSerServer interface {
	GetArea(context.Context, *Requset) (*Response, error)
}

// UnimplementedGetAreaSerServer can be embedded to have forward compatible implementations.
type UnimplementedGetAreaSerServer struct {
}

func (*UnimplementedGetAreaSerServer) GetArea(ctx context.Context, req *Requset) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArea not implemented")
}

func RegisterGetAreaSerServer(s *grpc.Server, srv GetAreaSerServer) {
	s.RegisterService(&_GetAreaSer_serviceDesc, srv)
}

func _GetAreaSer_GetArea_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Requset)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetAreaSerServer).GetArea(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.GetAreaSer/GetArea",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetAreaSerServer).GetArea(ctx, req.(*Requset))
	}
	return interceptor(ctx, in, info, handler)
}

var _GetAreaSer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message.GetAreaSer",
	HandlerType: (*GetAreaSerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetArea",
			Handler:    _GetAreaSer_GetArea_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "getArea.proto",
}
