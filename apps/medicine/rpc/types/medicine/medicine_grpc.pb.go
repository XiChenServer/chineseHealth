// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.11.2
// source: rpc/medicine.proto

package medicine

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Medicine_MedicineCreate_FullMethodName = "/medicine.Medicine/MedicineCreate"
	Medicine_MedicineDel_FullMethodName    = "/medicine.Medicine/MedicineDel"
	Medicine_MedicineMod_FullMethodName    = "/medicine.Medicine/MedicineMod"
	Medicine_MedicineFind_FullMethodName   = "/medicine.Medicine/MedicineFind"
	Medicine_ImageCreate_FullMethodName    = "/medicine.Medicine/ImageCreate"
	Medicine_ImageDelete_FullMethodName    = "/medicine.Medicine/ImageDelete"
)

// MedicineClient is the client API for Medicine service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MedicineClient interface {
	MedicineCreate(ctx context.Context, in *MedicineCreateRequest, opts ...grpc.CallOption) (*MedicineCreateResponse, error)
	MedicineDel(ctx context.Context, in *MedicineDelRequest, opts ...grpc.CallOption) (*MedicineDelResponse, error)
	MedicineMod(ctx context.Context, in *MedicineModRequest, opts ...grpc.CallOption) (*MedicineModResponse, error)
	MedicineFind(ctx context.Context, in *MedicineFindRequest, opts ...grpc.CallOption) (*MedicineFindResponse, error)
	ImageCreate(ctx context.Context, in *ImageCreateRequest, opts ...grpc.CallOption) (*ImageCreteResponse, error)
	ImageDelete(ctx context.Context, in *ImageDelRequest, opts ...grpc.CallOption) (*ImageDelResponse, error)
}

type medicineClient struct {
	cc grpc.ClientConnInterface
}

func NewMedicineClient(cc grpc.ClientConnInterface) MedicineClient {
	return &medicineClient{cc}
}

func (c *medicineClient) MedicineCreate(ctx context.Context, in *MedicineCreateRequest, opts ...grpc.CallOption) (*MedicineCreateResponse, error) {
	out := new(MedicineCreateResponse)
	err := c.cc.Invoke(ctx, Medicine_MedicineCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medicineClient) MedicineDel(ctx context.Context, in *MedicineDelRequest, opts ...grpc.CallOption) (*MedicineDelResponse, error) {
	out := new(MedicineDelResponse)
	err := c.cc.Invoke(ctx, Medicine_MedicineDel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medicineClient) MedicineMod(ctx context.Context, in *MedicineModRequest, opts ...grpc.CallOption) (*MedicineModResponse, error) {
	out := new(MedicineModResponse)
	err := c.cc.Invoke(ctx, Medicine_MedicineMod_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medicineClient) MedicineFind(ctx context.Context, in *MedicineFindRequest, opts ...grpc.CallOption) (*MedicineFindResponse, error) {
	out := new(MedicineFindResponse)
	err := c.cc.Invoke(ctx, Medicine_MedicineFind_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medicineClient) ImageCreate(ctx context.Context, in *ImageCreateRequest, opts ...grpc.CallOption) (*ImageCreteResponse, error) {
	out := new(ImageCreteResponse)
	err := c.cc.Invoke(ctx, Medicine_ImageCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medicineClient) ImageDelete(ctx context.Context, in *ImageDelRequest, opts ...grpc.CallOption) (*ImageDelResponse, error) {
	out := new(ImageDelResponse)
	err := c.cc.Invoke(ctx, Medicine_ImageDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MedicineServer is the server API for Medicine service.
// All implementations must embed UnimplementedMedicineServer
// for forward compatibility
type MedicineServer interface {
	MedicineCreate(context.Context, *MedicineCreateRequest) (*MedicineCreateResponse, error)
	MedicineDel(context.Context, *MedicineDelRequest) (*MedicineDelResponse, error)
	MedicineMod(context.Context, *MedicineModRequest) (*MedicineModResponse, error)
	MedicineFind(context.Context, *MedicineFindRequest) (*MedicineFindResponse, error)
	ImageCreate(context.Context, *ImageCreateRequest) (*ImageCreteResponse, error)
	ImageDelete(context.Context, *ImageDelRequest) (*ImageDelResponse, error)
	mustEmbedUnimplementedMedicineServer()
}

// UnimplementedMedicineServer must be embedded to have forward compatible implementations.
type UnimplementedMedicineServer struct {
}

func (UnimplementedMedicineServer) MedicineCreate(context.Context, *MedicineCreateRequest) (*MedicineCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MedicineCreate not implemented")
}
func (UnimplementedMedicineServer) MedicineDel(context.Context, *MedicineDelRequest) (*MedicineDelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MedicineDel not implemented")
}
func (UnimplementedMedicineServer) MedicineMod(context.Context, *MedicineModRequest) (*MedicineModResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MedicineMod not implemented")
}
func (UnimplementedMedicineServer) MedicineFind(context.Context, *MedicineFindRequest) (*MedicineFindResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MedicineFind not implemented")
}
func (UnimplementedMedicineServer) ImageCreate(context.Context, *ImageCreateRequest) (*ImageCreteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImageCreate not implemented")
}
func (UnimplementedMedicineServer) ImageDelete(context.Context, *ImageDelRequest) (*ImageDelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImageDelete not implemented")
}
func (UnimplementedMedicineServer) mustEmbedUnimplementedMedicineServer() {}

// UnsafeMedicineServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MedicineServer will
// result in compilation errors.
type UnsafeMedicineServer interface {
	mustEmbedUnimplementedMedicineServer()
}

func RegisterMedicineServer(s grpc.ServiceRegistrar, srv MedicineServer) {
	s.RegisterService(&Medicine_ServiceDesc, srv)
}

func _Medicine_MedicineCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MedicineCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicineServer).MedicineCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Medicine_MedicineCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicineServer).MedicineCreate(ctx, req.(*MedicineCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Medicine_MedicineDel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MedicineDelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicineServer).MedicineDel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Medicine_MedicineDel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicineServer).MedicineDel(ctx, req.(*MedicineDelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Medicine_MedicineMod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MedicineModRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicineServer).MedicineMod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Medicine_MedicineMod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicineServer).MedicineMod(ctx, req.(*MedicineModRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Medicine_MedicineFind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MedicineFindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicineServer).MedicineFind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Medicine_MedicineFind_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicineServer).MedicineFind(ctx, req.(*MedicineFindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Medicine_ImageCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicineServer).ImageCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Medicine_ImageCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicineServer).ImageCreate(ctx, req.(*ImageCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Medicine_ImageDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageDelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicineServer).ImageDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Medicine_ImageDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicineServer).ImageDelete(ctx, req.(*ImageDelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Medicine_ServiceDesc is the grpc.ServiceDesc for Medicine service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Medicine_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "medicine.Medicine",
	HandlerType: (*MedicineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MedicineCreate",
			Handler:    _Medicine_MedicineCreate_Handler,
		},
		{
			MethodName: "MedicineDel",
			Handler:    _Medicine_MedicineDel_Handler,
		},
		{
			MethodName: "MedicineMod",
			Handler:    _Medicine_MedicineMod_Handler,
		},
		{
			MethodName: "MedicineFind",
			Handler:    _Medicine_MedicineFind_Handler,
		},
		{
			MethodName: "ImageCreate",
			Handler:    _Medicine_ImageCreate_Handler,
		},
		{
			MethodName: "ImageDelete",
			Handler:    _Medicine_ImageDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/medicine.proto",
}