// Code generated by protoc-gen-go.
// source: dl.proto
// DO NOT EDIT!

/*
Package dlpb is a generated protocol buffer package.

It is generated from these files:
	dl.proto

It has these top-level messages:
	Down
	Link
	DownloadRequest
	DownloadResponse
	GetAllRequest
	GetAllResponse
	GetRequest
	GetResponse
	DelRequest
	DelResponse
	HookNamesRequest
	HookNamesResponse
*/
package dlpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Status int32

const (
	Status_QUEUED  Status = 0
	Status_RUNNING Status = 1
	Status_SUCCESS Status = 2
	Status_ERROR   Status = 3
)

var Status_name = map[int32]string{
	0: "QUEUED",
	1: "RUNNING",
	2: "SUCCESS",
	3: "ERROR",
}
var Status_value = map[string]int32{
	"QUEUED":  0,
	"RUNNING": 1,
	"SUCCESS": 2,
	"ERROR":   3,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Down struct {
	Id          int64    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Status      Status   `protobuf:"varint,3,opt,name=status,enum=dlpb.Status" json:"status,omitempty"`
	Errors      []string `protobuf:"bytes,4,rep,name=errors" json:"errors,omitempty"`
	Posthook    []string `protobuf:"bytes,5,rep,name=posthook" json:"posthook,omitempty"`
	Destination string   `protobuf:"bytes,6,opt,name=destination" json:"destination,omitempty"`
	CreatedAt   int64    `protobuf:"varint,7,opt,name=created_at" json:"created_at,omitempty"`
	ModifiedAt  int64    `protobuf:"varint,8,opt,name=modified_at" json:"modified_at,omitempty"`
	Links       []*Link  `protobuf:"bytes,9,rep,name=links" json:"links,omitempty"`
	StatusText  string   `protobuf:"bytes,10,opt,name=status_text" json:"status_text,omitempty"`
}

func (m *Down) Reset()                    { *m = Down{} }
func (m *Down) String() string            { return proto.CompactTextString(m) }
func (*Down) ProtoMessage()               {}
func (*Down) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Down) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Down) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Down) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_QUEUED
}

func (m *Down) GetErrors() []string {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *Down) GetPosthook() []string {
	if m != nil {
		return m.Posthook
	}
	return nil
}

func (m *Down) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *Down) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Down) GetModifiedAt() int64 {
	if m != nil {
		return m.ModifiedAt
	}
	return 0
}

func (m *Down) GetLinks() []*Link {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *Down) GetStatusText() string {
	if m != nil {
		return m.StatusText
	}
	return ""
}

type Link struct {
	Id         int64   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Url        string  `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	Status     Status  `protobuf:"varint,3,opt,name=status,enum=dlpb.Status" json:"status,omitempty"`
	CreatedAt  int64   `protobuf:"varint,4,opt,name=created_at" json:"created_at,omitempty"`
	ModifiedAt int64   `protobuf:"varint,5,opt,name=modified_at" json:"modified_at,omitempty"`
	Filename   string  `protobuf:"bytes,6,opt,name=filename" json:"filename,omitempty"`
	Percent    float64 `protobuf:"fixed64,7,opt,name=percent" json:"percent,omitempty"`
	StatusText string  `protobuf:"bytes,8,opt,name=status_text" json:"status_text,omitempty"`
}

func (m *Link) Reset()                    { *m = Link{} }
func (m *Link) String() string            { return proto.CompactTextString(m) }
func (*Link) ProtoMessage()               {}
func (*Link) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Link) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Link) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Link) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_QUEUED
}

func (m *Link) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Link) GetModifiedAt() int64 {
	if m != nil {
		return m.ModifiedAt
	}
	return 0
}

func (m *Link) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *Link) GetPercent() float64 {
	if m != nil {
		return m.Percent
	}
	return 0
}

func (m *Link) GetStatusText() string {
	if m != nil {
		return m.StatusText
	}
	return ""
}

type DownloadRequest struct {
	Down *Down `protobuf:"bytes,1,opt,name=down" json:"down,omitempty"`
}

func (m *DownloadRequest) Reset()                    { *m = DownloadRequest{} }
func (m *DownloadRequest) String() string            { return proto.CompactTextString(m) }
func (*DownloadRequest) ProtoMessage()               {}
func (*DownloadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DownloadRequest) GetDown() *Down {
	if m != nil {
		return m.Down
	}
	return nil
}

type DownloadResponse struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *DownloadResponse) Reset()                    { *m = DownloadResponse{} }
func (m *DownloadResponse) String() string            { return proto.CompactTextString(m) }
func (*DownloadResponse) ProtoMessage()               {}
func (*DownloadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DownloadResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetAllRequest struct {
	Statuses []Status `protobuf:"varint,1,rep,packed,name=statuses,enum=dlpb.Status" json:"statuses,omitempty"`
}

func (m *GetAllRequest) Reset()                    { *m = GetAllRequest{} }
func (m *GetAllRequest) String() string            { return proto.CompactTextString(m) }
func (*GetAllRequest) ProtoMessage()               {}
func (*GetAllRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetAllRequest) GetStatuses() []Status {
	if m != nil {
		return m.Statuses
	}
	return nil
}

type GetAllResponse struct {
	Downs []*Down `protobuf:"bytes,1,rep,name=downs" json:"downs,omitempty"`
}

func (m *GetAllResponse) Reset()                    { *m = GetAllResponse{} }
func (m *GetAllResponse) String() string            { return proto.CompactTextString(m) }
func (*GetAllResponse) ProtoMessage()               {}
func (*GetAllResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetAllResponse) GetDowns() []*Down {
	if m != nil {
		return m.Downs
	}
	return nil
}

type GetRequest struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetResponse struct {
	Down *Down `protobuf:"bytes,1,opt,name=down" json:"down,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GetResponse) GetDown() *Down {
	if m != nil {
		return m.Down
	}
	return nil
}

type DelRequest struct {
	Down *Down `protobuf:"bytes,1,opt,name=down" json:"down,omitempty"`
}

func (m *DelRequest) Reset()                    { *m = DelRequest{} }
func (m *DelRequest) String() string            { return proto.CompactTextString(m) }
func (*DelRequest) ProtoMessage()               {}
func (*DelRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *DelRequest) GetDown() *Down {
	if m != nil {
		return m.Down
	}
	return nil
}

type DelResponse struct {
}

func (m *DelResponse) Reset()                    { *m = DelResponse{} }
func (m *DelResponse) String() string            { return proto.CompactTextString(m) }
func (*DelResponse) ProtoMessage()               {}
func (*DelResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type HookNamesRequest struct {
}

func (m *HookNamesRequest) Reset()                    { *m = HookNamesRequest{} }
func (m *HookNamesRequest) String() string            { return proto.CompactTextString(m) }
func (*HookNamesRequest) ProtoMessage()               {}
func (*HookNamesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type HookNamesResponse struct {
	Names []string `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
}

func (m *HookNamesResponse) Reset()                    { *m = HookNamesResponse{} }
func (m *HookNamesResponse) String() string            { return proto.CompactTextString(m) }
func (*HookNamesResponse) ProtoMessage()               {}
func (*HookNamesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *HookNamesResponse) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func init() {
	proto.RegisterType((*Down)(nil), "dlpb.Down")
	proto.RegisterType((*Link)(nil), "dlpb.Link")
	proto.RegisterType((*DownloadRequest)(nil), "dlpb.DownloadRequest")
	proto.RegisterType((*DownloadResponse)(nil), "dlpb.DownloadResponse")
	proto.RegisterType((*GetAllRequest)(nil), "dlpb.GetAllRequest")
	proto.RegisterType((*GetAllResponse)(nil), "dlpb.GetAllResponse")
	proto.RegisterType((*GetRequest)(nil), "dlpb.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "dlpb.GetResponse")
	proto.RegisterType((*DelRequest)(nil), "dlpb.DelRequest")
	proto.RegisterType((*DelResponse)(nil), "dlpb.DelResponse")
	proto.RegisterType((*HookNamesRequest)(nil), "dlpb.HookNamesRequest")
	proto.RegisterType((*HookNamesResponse)(nil), "dlpb.HookNamesResponse")
	proto.RegisterEnum("dlpb.Status", Status_name, Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Dl service

type DlClient interface {
	Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (*DownloadResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Del(ctx context.Context, in *DelRequest, opts ...grpc.CallOption) (*DelResponse, error)
	HookNames(ctx context.Context, in *HookNamesRequest, opts ...grpc.CallOption) (*HookNamesResponse, error)
}

type dlClient struct {
	cc *grpc.ClientConn
}

func NewDlClient(cc *grpc.ClientConn) DlClient {
	return &dlClient{cc}
}

func (c *dlClient) Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (*DownloadResponse, error) {
	out := new(DownloadResponse)
	err := grpc.Invoke(ctx, "/dlpb.Dl/Download", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dlClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := grpc.Invoke(ctx, "/dlpb.Dl/GetAll", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dlClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/dlpb.Dl/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dlClient) Del(ctx context.Context, in *DelRequest, opts ...grpc.CallOption) (*DelResponse, error) {
	out := new(DelResponse)
	err := grpc.Invoke(ctx, "/dlpb.Dl/Del", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dlClient) HookNames(ctx context.Context, in *HookNamesRequest, opts ...grpc.CallOption) (*HookNamesResponse, error) {
	out := new(HookNamesResponse)
	err := grpc.Invoke(ctx, "/dlpb.Dl/HookNames", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Dl service

type DlServer interface {
	Download(context.Context, *DownloadRequest) (*DownloadResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Del(context.Context, *DelRequest) (*DelResponse, error)
	HookNames(context.Context, *HookNamesRequest) (*HookNamesResponse, error)
}

func RegisterDlServer(s *grpc.Server, srv DlServer) {
	s.RegisterService(&_Dl_serviceDesc, srv)
}

func _Dl_Download_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DlServer).Download(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dlpb.Dl/Download",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DlServer).Download(ctx, req.(*DownloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dl_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DlServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dlpb.Dl/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DlServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dl_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DlServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dlpb.Dl/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DlServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dl_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DlServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dlpb.Dl/Del",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DlServer).Del(ctx, req.(*DelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dl_HookNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HookNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DlServer).HookNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dlpb.Dl/HookNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DlServer).HookNames(ctx, req.(*HookNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dl_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dlpb.Dl",
	HandlerType: (*DlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Download",
			Handler:    _Dl_Download_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _Dl_GetAll_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Dl_Get_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _Dl_Del_Handler,
		},
		{
			MethodName: "HookNames",
			Handler:    _Dl_HookNames_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dl.proto",
}

func init() { proto.RegisterFile("dl.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 525 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x5d, 0x8b, 0xd3, 0x4c,
	0x18, 0xdd, 0x34, 0x1f, 0x9b, 0x9c, 0x6c, 0xbb, 0xd9, 0xe9, 0xfb, 0xea, 0x6c, 0x91, 0x25, 0xe4,
	0x42, 0x8b, 0x2b, 0x15, 0x2a, 0x5e, 0x88, 0x20, 0xc8, 0xb6, 0x54, 0x41, 0x2a, 0xb6, 0xf4, 0x7a,
	0xc9, 0x6e, 0x66, 0x31, 0x34, 0xcd, 0xd4, 0xcc, 0x94, 0xf5, 0x87, 0x78, 0xed, 0x4f, 0xf2, 0x37,
	0x49, 0x66, 0x92, 0x36, 0xad, 0x8a, 0x5e, 0xe6, 0xcc, 0x79, 0x9e, 0xf3, 0xf1, 0x40, 0xe0, 0x26,
	0xd9, 0x60, 0x5d, 0x70, 0xc9, 0x89, 0x95, 0x64, 0xeb, 0x9b, 0xe8, 0x87, 0x01, 0x6b, 0xc4, 0xef,
	0x73, 0x02, 0xb4, 0xd2, 0x84, 0x1a, 0xa1, 0xd1, 0x37, 0xc9, 0x09, 0xac, 0x3c, 0x5e, 0x31, 0xda,
	0x0a, 0x8d, 0xbe, 0x47, 0x1e, 0xc1, 0x11, 0x32, 0x96, 0x1b, 0x41, 0xcd, 0xd0, 0xe8, 0x77, 0x86,
	0x27, 0x83, 0x72, 0x72, 0x30, 0x57, 0x18, 0xe9, 0xc0, 0x61, 0x45, 0xc1, 0x0b, 0x41, 0xad, 0xd0,
	0xec, 0x7b, 0x24, 0x80, 0xbb, 0xe6, 0x42, 0x7e, 0xe6, 0x7c, 0x49, 0x6d, 0x85, 0x74, 0xe1, 0x27,
	0x4c, 0xc8, 0x34, 0x8f, 0x65, 0xca, 0x73, 0xea, 0xa8, 0xa5, 0x04, 0xb8, 0x2d, 0x58, 0x2c, 0x59,
	0x72, 0x1d, 0x4b, 0x7a, 0xac, 0x64, 0xbb, 0xf0, 0x57, 0x3c, 0x49, 0xef, 0x52, 0x0d, 0xba, 0x0a,
	0x3c, 0x87, 0x9d, 0xa5, 0xf9, 0x52, 0x50, 0x2f, 0x34, 0xfb, 0xfe, 0x10, 0x5a, 0xfc, 0x43, 0x9a,
	0x2f, 0x4b, 0xbe, 0x36, 0x76, 0x2d, 0xd9, 0x57, 0x49, 0x51, 0x2e, 0x8e, 0xbe, 0x1b, 0xb0, 0xd4,
	0x6b, 0x33, 0x90, 0x0f, 0x73, 0x53, 0x64, 0xff, 0x94, 0x67, 0xdf, 0x98, 0xf5, 0x3b, 0x63, 0xb6,
	0x02, 0x03, 0xb8, 0x77, 0x69, 0xc6, 0x54, 0x51, 0x3a, 0xd3, 0x29, 0x8e, 0xd7, 0xac, 0xb8, 0x65,
	0xb9, 0x0e, 0x64, 0x1c, 0x1a, 0x74, 0x95, 0xc1, 0x4b, 0x9c, 0x96, 0x85, 0x67, 0x3c, 0x4e, 0x66,
	0xec, 0xcb, 0x86, 0x09, 0x49, 0x28, 0xac, 0x84, 0xdf, 0xe7, 0xca, 0xec, 0x36, 0x62, 0x49, 0x8a,
	0x2e, 0x10, 0xec, 0xc8, 0x62, 0xcd, 0x73, 0xc1, 0x9a, 0xc1, 0xa2, 0xe7, 0x68, 0x4f, 0x98, 0x7c,
	0x9b, 0x65, 0xf5, 0xaa, 0x0b, 0xb8, 0x5a, 0x92, 0x09, 0x6a, 0x84, 0xe6, 0x61, 0xbc, 0xe8, 0x12,
	0x9d, 0x7a, 0xa0, 0x5a, 0x77, 0x0e, 0xbb, 0x14, 0xd7, 0xf4, 0x7d, 0x75, 0x0a, 0x4c, 0x98, 0xac,
	0x57, 0x37, 0x75, 0x9f, 0xc0, 0x57, 0x2f, 0xd5, 0x8e, 0x3f, 0x07, 0x78, 0x0c, 0x8c, 0x58, 0xf6,
	0xf7, 0xa0, 0x6d, 0xf8, 0x8a, 0xa7, 0x17, 0x46, 0x04, 0xc1, 0x3b, 0xce, 0x97, 0xd3, 0x78, 0xc5,
	0x44, 0x35, 0x1c, 0x45, 0x38, 0x6b, 0x60, 0x95, 0x72, 0x1b, 0x76, 0x79, 0x01, 0xed, 0xde, 0x7b,
	0xfa, 0x0a, 0x4e, 0x75, 0x47, 0xc0, 0xf9, 0xb4, 0x18, 0x2f, 0xc6, 0xa3, 0xe0, 0x88, 0xf8, 0x38,
	0x9e, 0x2d, 0xa6, 0xd3, 0xf7, 0xd3, 0x49, 0x60, 0x94, 0x1f, 0xf3, 0xc5, 0xd5, 0xd5, 0x78, 0x3e,
	0x0f, 0x5a, 0xc4, 0x83, 0x3d, 0x9e, 0xcd, 0x3e, 0xce, 0x02, 0x73, 0xf8, 0xad, 0x85, 0xd6, 0x28,
	0x23, 0xaf, 0xe1, 0xd6, 0x8d, 0x93, 0xff, 0x77, 0x06, 0x1b, 0xe7, 0xea, 0x3d, 0x38, 0x84, 0x2b,
	0xd3, 0x47, 0xe4, 0x25, 0x1c, 0xdd, 0x2e, 0xe9, 0x6a, 0xce, 0xde, 0x71, 0x7a, 0xff, 0xed, 0x83,
	0xdb, 0xb1, 0x67, 0x30, 0x27, 0x4c, 0x92, 0x60, 0xfb, 0x5c, 0x0f, 0x9c, 0x35, 0x90, 0x26, 0x7b,
	0xc4, 0xb2, 0x9a, 0xbd, 0x6b, 0xb7, 0x66, 0x37, 0x7b, 0x3c, 0x22, 0x6f, 0xe0, 0x6d, 0x5b, 0x23,
	0x95, 0xf3, 0xc3, 0x6a, 0x7b, 0x0f, 0x7f, 0xc1, 0xeb, 0xf9, 0x1b, 0x47, 0xfd, 0x2d, 0x5e, 0xfc,
	0x0c, 0x00, 0x00, 0xff, 0xff, 0x43, 0x62, 0xbf, 0xfc, 0x39, 0x04, 0x00, 0x00,
}
