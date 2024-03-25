// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: alignedlayer/verify/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgUpdateParams is the Msg/UpdateParams request type.
type MsgUpdateParams struct {
	// authority is the address that controls the module (defaults to x/gov unless overwritten).
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// NOTE: All parameters must be supplied.
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *MsgUpdateParams) Reset()         { *m = MsgUpdateParams{} }
func (m *MsgUpdateParams) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParams) ProtoMessage()    {}
func (*MsgUpdateParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_a665612caf78653b, []int{0}
}
func (m *MsgUpdateParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParams.Merge(m, src)
}
func (m *MsgUpdateParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParams proto.InternalMessageInfo

func (m *MsgUpdateParams) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgUpdateParams) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// MsgUpdateParamsResponse defines the response structure for executing a
// MsgUpdateParams message.
type MsgUpdateParamsResponse struct {
}

func (m *MsgUpdateParamsResponse) Reset()         { *m = MsgUpdateParamsResponse{} }
func (m *MsgUpdateParamsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParamsResponse) ProtoMessage()    {}
func (*MsgUpdateParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a665612caf78653b, []int{1}
}
func (m *MsgUpdateParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParamsResponse.Merge(m, src)
}
func (m *MsgUpdateParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParamsResponse proto.InternalMessageInfo

type MsgGnarkPlonk struct {
	Creator      string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Proof        string `protobuf:"bytes,2,opt,name=proof,proto3" json:"proof,omitempty"`
	PublicInputs string `protobuf:"bytes,3,opt,name=publicInputs,proto3" json:"publicInputs,omitempty"`
	VerifyingKey string `protobuf:"bytes,4,opt,name=verifyingKey,proto3" json:"verifyingKey,omitempty"`
}

func (m *MsgGnarkPlonk) Reset()         { *m = MsgGnarkPlonk{} }
func (m *MsgGnarkPlonk) String() string { return proto.CompactTextString(m) }
func (*MsgGnarkPlonk) ProtoMessage()    {}
func (*MsgGnarkPlonk) Descriptor() ([]byte, []int) {
	return fileDescriptor_a665612caf78653b, []int{2}
}
func (m *MsgGnarkPlonk) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGnarkPlonk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGnarkPlonk.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGnarkPlonk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGnarkPlonk.Merge(m, src)
}
func (m *MsgGnarkPlonk) XXX_Size() int {
	return m.Size()
}
func (m *MsgGnarkPlonk) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGnarkPlonk.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGnarkPlonk proto.InternalMessageInfo

func (m *MsgGnarkPlonk) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgGnarkPlonk) GetProof() string {
	if m != nil {
		return m.Proof
	}
	return ""
}

func (m *MsgGnarkPlonk) GetPublicInputs() string {
	if m != nil {
		return m.PublicInputs
	}
	return ""
}

func (m *MsgGnarkPlonk) GetVerifyingKey() string {
	if m != nil {
		return m.VerifyingKey
	}
	return ""
}

type MsgGnarkPlonkResponse struct {
}

func (m *MsgGnarkPlonkResponse) Reset()         { *m = MsgGnarkPlonkResponse{} }
func (m *MsgGnarkPlonkResponse) String() string { return proto.CompactTextString(m) }
func (*MsgGnarkPlonkResponse) ProtoMessage()    {}
func (*MsgGnarkPlonkResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a665612caf78653b, []int{3}
}
func (m *MsgGnarkPlonkResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGnarkPlonkResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGnarkPlonkResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGnarkPlonkResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGnarkPlonkResponse.Merge(m, src)
}
func (m *MsgGnarkPlonkResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgGnarkPlonkResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGnarkPlonkResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGnarkPlonkResponse proto.InternalMessageInfo

type MsgCairoPlatinum struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Proof   string `protobuf:"bytes,2,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (m *MsgCairoPlatinum) Reset()         { *m = MsgCairoPlatinum{} }
func (m *MsgCairoPlatinum) String() string { return proto.CompactTextString(m) }
func (*MsgCairoPlatinum) ProtoMessage()    {}
func (*MsgCairoPlatinum) Descriptor() ([]byte, []int) {
	return fileDescriptor_a665612caf78653b, []int{4}
}
func (m *MsgCairoPlatinum) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCairoPlatinum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCairoPlatinum.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCairoPlatinum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCairoPlatinum.Merge(m, src)
}
func (m *MsgCairoPlatinum) XXX_Size() int {
	return m.Size()
}
func (m *MsgCairoPlatinum) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCairoPlatinum.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCairoPlatinum proto.InternalMessageInfo

func (m *MsgCairoPlatinum) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCairoPlatinum) GetProof() string {
	if m != nil {
		return m.Proof
	}
	return ""
}

type MsgCairoPlatinumResponse struct {
}

func (m *MsgCairoPlatinumResponse) Reset()         { *m = MsgCairoPlatinumResponse{} }
func (m *MsgCairoPlatinumResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCairoPlatinumResponse) ProtoMessage()    {}
func (*MsgCairoPlatinumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a665612caf78653b, []int{5}
}
func (m *MsgCairoPlatinumResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCairoPlatinumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCairoPlatinumResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCairoPlatinumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCairoPlatinumResponse.Merge(m, src)
}
func (m *MsgCairoPlatinumResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCairoPlatinumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCairoPlatinumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCairoPlatinumResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgUpdateParams)(nil), "alignedlayer.verify.MsgUpdateParams")
	proto.RegisterType((*MsgUpdateParamsResponse)(nil), "alignedlayer.verify.MsgUpdateParamsResponse")
	proto.RegisterType((*MsgGnarkPlonk)(nil), "alignedlayer.verify.MsgGnarkPlonk")
	proto.RegisterType((*MsgGnarkPlonkResponse)(nil), "alignedlayer.verify.MsgGnarkPlonkResponse")
	proto.RegisterType((*MsgCairoPlatinum)(nil), "alignedlayer.verify.MsgCairoPlatinum")
	proto.RegisterType((*MsgCairoPlatinumResponse)(nil), "alignedlayer.verify.MsgCairoPlatinumResponse")
}

func init() { proto.RegisterFile("alignedlayer/verify/tx.proto", fileDescriptor_a665612caf78653b) }

var fileDescriptor_a665612caf78653b = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xce, 0xb6, 0xb6, 0x92, 0x67, 0x8a, 0xba, 0x46, 0xb2, 0xdd, 0xca, 0x1a, 0x16, 0x0b, 0x25,
	0xd8, 0x2c, 0x56, 0x14, 0xe9, 0x41, 0x30, 0x1e, 0x44, 0x24, 0x10, 0x56, 0xbc, 0x88, 0x20, 0x93,
	0x64, 0x3a, 0x0e, 0xdd, 0x9d, 0x59, 0x66, 0x26, 0xa5, 0x7b, 0x13, 0x8f, 0x9e, 0xbc, 0xf8, 0x1f,
	0x3c, 0xe6, 0xa0, 0xbf, 0xc0, 0x4b, 0x8f, 0xc5, 0x93, 0x27, 0x91, 0xe4, 0x90, 0xbf, 0x21, 0xbb,
	0xb3, 0xdb, 0x64, 0x97, 0x14, 0xdb, 0xcb, 0xee, 0xbe, 0xf7, 0x7d, 0xef, 0x7d, 0xdf, 0x9b, 0x79,
	0x0b, 0x77, 0x50, 0x40, 0x09, 0xc3, 0xc3, 0x00, 0xc5, 0x58, 0x78, 0x47, 0x58, 0xd0, 0x83, 0xd8,
	0x53, 0xc7, 0xed, 0x48, 0x70, 0xc5, 0xcd, 0x5b, 0x8b, 0x68, 0x5b, 0xa3, 0xf6, 0x4d, 0x14, 0x52,
	0xc6, 0xbd, 0xf4, 0xa9, 0x79, 0x76, 0x63, 0xc0, 0x65, 0xc8, 0xa5, 0x17, 0x4a, 0xe2, 0x1d, 0x3d,
	0x48, 0x5e, 0x19, 0xb0, 0xa9, 0x81, 0xf7, 0x69, 0xe4, 0xe9, 0x20, 0x83, 0xea, 0x84, 0x13, 0xae,
	0xf3, 0xc9, 0x57, 0x96, 0x6d, 0x2e, 0xf3, 0x13, 0x21, 0x81, 0xc2, 0xac, 0xce, 0xfd, 0x69, 0xc0,
	0xf5, 0xae, 0x24, 0x6f, 0xa2, 0x21, 0x52, 0xb8, 0x97, 0x22, 0xe6, 0x63, 0xa8, 0xa2, 0x91, 0xfa,
	0xc0, 0x05, 0x55, 0xb1, 0x65, 0x34, 0x8d, 0x9d, 0x6a, 0xc7, 0xfa, 0xf5, 0x7d, 0xb7, 0x9e, 0x09,
	0x3e, 0x1b, 0x0e, 0x05, 0x96, 0xf2, 0xb5, 0x12, 0x94, 0x11, 0x7f, 0x4e, 0x35, 0x9f, 0xc2, 0xba,
	0xee, 0x6d, 0xad, 0x34, 0x8d, 0x9d, 0x6b, 0x7b, 0x5b, 0xed, 0x25, 0x03, 0xb7, 0xb5, 0x48, 0xa7,
	0x7a, 0xf2, 0xe7, 0x6e, 0xe5, 0xdb, 0x6c, 0xdc, 0x32, 0xfc, 0xac, 0x6a, 0xff, 0xc9, 0xa7, 0xd9,
	0xb8, 0x35, 0xef, 0xf7, 0x79, 0x36, 0x6e, 0x6d, 0x17, 0x06, 0x38, 0xce, 0x47, 0x28, 0x39, 0x76,
	0x37, 0xa1, 0x51, 0x4a, 0xf9, 0x58, 0x46, 0x9c, 0x49, 0xec, 0x7e, 0x35, 0x60, 0xa3, 0x2b, 0xc9,
	0x0b, 0x86, 0xc4, 0x61, 0x2f, 0xe0, 0xec, 0xd0, 0xb4, 0xe0, 0xea, 0x40, 0x60, 0xa4, 0xb8, 0xd0,
	0xc3, 0xf9, 0x79, 0x68, 0xd6, 0x61, 0x2d, 0x12, 0x9c, 0x1f, 0xa4, 0xfe, 0xab, 0xbe, 0x0e, 0x4c,
	0x17, 0x6a, 0xd1, 0xa8, 0x1f, 0xd0, 0xc1, 0x4b, 0x16, 0x8d, 0x94, 0xb4, 0x56, 0x53, 0xb0, 0x90,
	0x4b, 0x38, 0xda, 0x1a, 0x65, 0xe4, 0x15, 0x8e, 0xad, 0x2b, 0x9a, 0xb3, 0x98, 0xdb, 0xaf, 0x25,
	0xe3, 0xe5, 0x5a, 0x6e, 0x03, 0x6e, 0x17, 0x6c, 0x9d, 0x19, 0xee, 0xc1, 0x8d, 0xae, 0x24, 0xcf,
	0x11, 0x15, 0xbc, 0x17, 0x20, 0x45, 0xd9, 0x28, 0xbc, 0xac, 0xe5, 0x92, 0x94, 0x0d, 0x56, 0xb9,
	0x63, 0xae, 0xb6, 0xf7, 0x63, 0x05, 0x56, 0xbb, 0x92, 0x98, 0x7d, 0xa8, 0x15, 0x76, 0xe0, 0xde,
	0xd2, 0xbb, 0x2b, 0x1d, 0xb2, 0x7d, 0xff, 0x22, 0xac, 0x5c, 0xcb, 0x7c, 0x07, 0xb0, 0x70, 0x0d,
	0xee, 0x79, 0xb5, 0x73, 0x8e, 0xdd, 0xfa, 0x3f, 0xe7, 0xac, 0x3b, 0x86, 0x8d, 0xe2, 0xa1, 0x6d,
	0x9f, 0x57, 0x5c, 0xa0, 0xd9, 0xbb, 0x17, 0xa2, 0xe5, 0x32, 0xf6, 0xda, 0xc7, 0x64, 0x67, 0x3b,
	0x8f, 0x4e, 0x26, 0x8e, 0x71, 0x3a, 0x71, 0x8c, 0xbf, 0x13, 0xc7, 0xf8, 0x32, 0x75, 0x2a, 0xa7,
	0x53, 0xa7, 0xf2, 0x7b, 0xea, 0x54, 0xde, 0x6e, 0x2d, 0x5f, 0x59, 0x15, 0x47, 0x58, 0xf6, 0xd7,
	0xd3, 0xbf, 0xee, 0xe1, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfb, 0xd1, 0x2e, 0xf2, 0x29, 0x04,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	GnarkPlonk(ctx context.Context, in *MsgGnarkPlonk, opts ...grpc.CallOption) (*MsgGnarkPlonkResponse, error)
	CairoPlatinum(ctx context.Context, in *MsgCairoPlatinum, opts ...grpc.CallOption) (*MsgCairoPlatinumResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, "/alignedlayer.verify.Msg/UpdateParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GnarkPlonk(ctx context.Context, in *MsgGnarkPlonk, opts ...grpc.CallOption) (*MsgGnarkPlonkResponse, error) {
	out := new(MsgGnarkPlonkResponse)
	err := c.cc.Invoke(ctx, "/alignedlayer.verify.Msg/GnarkPlonk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CairoPlatinum(ctx context.Context, in *MsgCairoPlatinum, opts ...grpc.CallOption) (*MsgCairoPlatinumResponse, error) {
	out := new(MsgCairoPlatinumResponse)
	err := c.cc.Invoke(ctx, "/alignedlayer.verify.Msg/CairoPlatinum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	GnarkPlonk(context.Context, *MsgGnarkPlonk) (*MsgGnarkPlonkResponse, error)
	CairoPlatinum(context.Context, *MsgCairoPlatinum) (*MsgCairoPlatinumResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) UpdateParams(ctx context.Context, req *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (*UnimplementedMsgServer) GnarkPlonk(ctx context.Context, req *MsgGnarkPlonk) (*MsgGnarkPlonkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GnarkPlonk not implemented")
}
func (*UnimplementedMsgServer) CairoPlatinum(ctx context.Context, req *MsgCairoPlatinum) (*MsgCairoPlatinumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CairoPlatinum not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alignedlayer.verify.Msg/UpdateParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_GnarkPlonk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGnarkPlonk)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GnarkPlonk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alignedlayer.verify.Msg/GnarkPlonk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GnarkPlonk(ctx, req.(*MsgGnarkPlonk))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CairoPlatinum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCairoPlatinum)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CairoPlatinum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alignedlayer.verify.Msg/CairoPlatinum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CairoPlatinum(ctx, req.(*MsgCairoPlatinum))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "alignedlayer.verify.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "GnarkPlonk",
			Handler:    _Msg_GnarkPlonk_Handler,
		},
		{
			MethodName: "CairoPlatinum",
			Handler:    _Msg_CairoPlatinum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "alignedlayer/verify/tx.proto",
}

func (m *MsgUpdateParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgGnarkPlonk) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGnarkPlonk) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGnarkPlonk) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VerifyingKey) > 0 {
		i -= len(m.VerifyingKey)
		copy(dAtA[i:], m.VerifyingKey)
		i = encodeVarintTx(dAtA, i, uint64(len(m.VerifyingKey)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.PublicInputs) > 0 {
		i -= len(m.PublicInputs)
		copy(dAtA[i:], m.PublicInputs)
		i = encodeVarintTx(dAtA, i, uint64(len(m.PublicInputs)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Proof) > 0 {
		i -= len(m.Proof)
		copy(dAtA[i:], m.Proof)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Proof)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgGnarkPlonkResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGnarkPlonkResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGnarkPlonkResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgCairoPlatinum) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCairoPlatinum) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCairoPlatinum) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Proof) > 0 {
		i -= len(m.Proof)
		copy(dAtA[i:], m.Proof)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Proof)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCairoPlatinumResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCairoPlatinumResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCairoPlatinumResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgUpdateParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Params.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgUpdateParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgGnarkPlonk) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Proof)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.PublicInputs)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.VerifyingKey)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgGnarkPlonkResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgCairoPlatinum) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Proof)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCairoPlatinumResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgUpdateParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgGnarkPlonk) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgGnarkPlonk: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGnarkPlonk: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proof = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicInputs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicInputs = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerifyingKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VerifyingKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgGnarkPlonkResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgGnarkPlonkResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGnarkPlonkResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCairoPlatinum) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCairoPlatinum: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCairoPlatinum: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proof = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCairoPlatinumResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCairoPlatinumResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCairoPlatinumResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)