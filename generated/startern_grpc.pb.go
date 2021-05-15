// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package generated

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

// CompaniesClient is the client API for Companies service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompaniesClient interface {
	GetCompanies(ctx context.Context, in *GetCompaniesRequest, opts ...grpc.CallOption) (*GetCompaniesResponse, error)
}

type companiesClient struct {
	cc grpc.ClientConnInterface
}

func NewCompaniesClient(cc grpc.ClientConnInterface) CompaniesClient {
	return &companiesClient{cc}
}

func (c *companiesClient) GetCompanies(ctx context.Context, in *GetCompaniesRequest, opts ...grpc.CallOption) (*GetCompaniesResponse, error) {
	out := new(GetCompaniesResponse)
	err := c.cc.Invoke(ctx, "/Companies/GetCompanies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompaniesServer is the server API for Companies service.
// All implementations must embed UnimplementedCompaniesServer
// for forward compatibility
type CompaniesServer interface {
	GetCompanies(context.Context, *GetCompaniesRequest) (*GetCompaniesResponse, error)
	mustEmbedUnimplementedCompaniesServer()
}

// UnimplementedCompaniesServer must be embedded to have forward compatible implementations.
type UnimplementedCompaniesServer struct {
}

func (UnimplementedCompaniesServer) GetCompanies(context.Context, *GetCompaniesRequest) (*GetCompaniesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanies not implemented")
}
func (UnimplementedCompaniesServer) mustEmbedUnimplementedCompaniesServer() {}

// UnsafeCompaniesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompaniesServer will
// result in compilation errors.
type UnsafeCompaniesServer interface {
	mustEmbedUnimplementedCompaniesServer()
}

func RegisterCompaniesServer(s grpc.ServiceRegistrar, srv CompaniesServer) {
	s.RegisterService(&Companies_ServiceDesc, srv)
}

func _Companies_GetCompanies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompaniesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompaniesServer).GetCompanies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Companies/GetCompanies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompaniesServer).GetCompanies(ctx, req.(*GetCompaniesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Companies_ServiceDesc is the grpc.ServiceDesc for Companies service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Companies_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Companies",
	HandlerType: (*CompaniesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCompanies",
			Handler:    _Companies_GetCompanies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "startern.proto",
}

// VacanciesClient is the client API for Vacancies service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VacanciesClient interface {
	GetVacancies(ctx context.Context, in *GetVacanciesRequest, opts ...grpc.CallOption) (*GetVacanciesResponse, error)
}

type vacanciesClient struct {
	cc grpc.ClientConnInterface
}

func NewVacanciesClient(cc grpc.ClientConnInterface) VacanciesClient {
	return &vacanciesClient{cc}
}

func (c *vacanciesClient) GetVacancies(ctx context.Context, in *GetVacanciesRequest, opts ...grpc.CallOption) (*GetVacanciesResponse, error) {
	out := new(GetVacanciesResponse)
	err := c.cc.Invoke(ctx, "/Vacancies/GetVacancies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VacanciesServer is the server API for Vacancies service.
// All implementations must embed UnimplementedVacanciesServer
// for forward compatibility
type VacanciesServer interface {
	GetVacancies(context.Context, *GetVacanciesRequest) (*GetVacanciesResponse, error)
	mustEmbedUnimplementedVacanciesServer()
}

// UnimplementedVacanciesServer must be embedded to have forward compatible implementations.
type UnimplementedVacanciesServer struct {
}

func (UnimplementedVacanciesServer) GetVacancies(context.Context, *GetVacanciesRequest) (*GetVacanciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVacancies not implemented")
}
func (UnimplementedVacanciesServer) mustEmbedUnimplementedVacanciesServer() {}

// UnsafeVacanciesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VacanciesServer will
// result in compilation errors.
type UnsafeVacanciesServer interface {
	mustEmbedUnimplementedVacanciesServer()
}

func RegisterVacanciesServer(s grpc.ServiceRegistrar, srv VacanciesServer) {
	s.RegisterService(&Vacancies_ServiceDesc, srv)
}

func _Vacancies_GetVacancies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVacanciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacanciesServer).GetVacancies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Vacancies/GetVacancies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacanciesServer).GetVacancies(ctx, req.(*GetVacanciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Vacancies_ServiceDesc is the grpc.ServiceDesc for Vacancies service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Vacancies_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Vacancies",
	HandlerType: (*VacanciesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVacancies",
			Handler:    _Vacancies_GetVacancies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "startern.proto",
}
