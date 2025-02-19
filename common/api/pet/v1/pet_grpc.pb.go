// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: pet.proto

package pet_v1

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PetV1_Create_FullMethodName   = "/pet_v1.PetV1/Create"
	PetV1_Get_FullMethodName      = "/pet_v1.PetV1/Get"
	PetV1_ListPets_FullMethodName = "/pet_v1.PetV1/ListPets"
	PetV1_Update_FullMethodName   = "/pet_v1.PetV1/Update"
	PetV1_Delete_FullMethodName   = "/pet_v1.PetV1/Delete"
)

// PetV1Client is the client API for PetV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PetV1Client interface {
	// Creates a new pet
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Gets a pet by id
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// Gets a list of pets with filters
	ListPets(ctx context.Context, in *ListPetsRequest, opts ...grpc.CallOption) (*ListPetsResponse, error)
	// Updates a user's pet
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Deletes a user's pet
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type petV1Client struct {
	cc grpc.ClientConnInterface
}

func NewPetV1Client(cc grpc.ClientConnInterface) PetV1Client {
	return &petV1Client{cc}
}

func (c *petV1Client) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, PetV1_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petV1Client) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, PetV1_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petV1Client) ListPets(ctx context.Context, in *ListPetsRequest, opts ...grpc.CallOption) (*ListPetsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPetsResponse)
	err := c.cc.Invoke(ctx, PetV1_ListPets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petV1Client) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, PetV1_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petV1Client) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, PetV1_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PetV1Server is the server API for PetV1 service.
// All implementations must embed UnimplementedPetV1Server
// for forward compatibility.
type PetV1Server interface {
	// Creates a new pet
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Gets a pet by id
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// Gets a list of pets with filters
	ListPets(context.Context, *ListPetsRequest) (*ListPetsResponse, error)
	// Updates a user's pet
	Update(context.Context, *UpdateRequest) (*empty.Empty, error)
	// Deletes a user's pet
	Delete(context.Context, *DeleteRequest) (*empty.Empty, error)
	mustEmbedUnimplementedPetV1Server()
}

// UnimplementedPetV1Server must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPetV1Server struct{}

func (UnimplementedPetV1Server) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPetV1Server) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPetV1Server) ListPets(context.Context, *ListPetsRequest) (*ListPetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPets not implemented")
}
func (UnimplementedPetV1Server) Update(context.Context, *UpdateRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPetV1Server) Delete(context.Context, *DeleteRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPetV1Server) mustEmbedUnimplementedPetV1Server() {}
func (UnimplementedPetV1Server) testEmbeddedByValue()               {}

// UnsafePetV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PetV1Server will
// result in compilation errors.
type UnsafePetV1Server interface {
	mustEmbedUnimplementedPetV1Server()
}

func RegisterPetV1Server(s grpc.ServiceRegistrar, srv PetV1Server) {
	// If the following call pancis, it indicates UnimplementedPetV1Server was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PetV1_ServiceDesc, srv)
}

func _PetV1_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetV1Server).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetV1_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetV1Server).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetV1_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetV1Server).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetV1_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetV1Server).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetV1_ListPets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetV1Server).ListPets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetV1_ListPets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetV1Server).ListPets(ctx, req.(*ListPetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetV1_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetV1Server).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetV1_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetV1Server).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetV1_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetV1Server).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetV1_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetV1Server).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PetV1_ServiceDesc is the grpc.ServiceDesc for PetV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PetV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pet_v1.PetV1",
	HandlerType: (*PetV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PetV1_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _PetV1_Get_Handler,
		},
		{
			MethodName: "ListPets",
			Handler:    _PetV1_ListPets_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PetV1_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PetV1_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pet.proto",
}
