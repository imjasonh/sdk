// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: policy_results.platform.proto

package v1

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
	PolicyResults_List_FullMethodName = "/chainguard.platform.tenant.PolicyResults/List"
)

// PolicyResultsClient is the client API for PolicyResults service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PolicyResultsClient interface {
	List(ctx context.Context, in *PolicyResultFilter, opts ...grpc.CallOption) (*PolicyResultList, error)
}

type policyResultsClient struct {
	cc grpc.ClientConnInterface
}

func NewPolicyResultsClient(cc grpc.ClientConnInterface) PolicyResultsClient {
	return &policyResultsClient{cc}
}

func (c *policyResultsClient) List(ctx context.Context, in *PolicyResultFilter, opts ...grpc.CallOption) (*PolicyResultList, error) {
	out := new(PolicyResultList)
	err := c.cc.Invoke(ctx, PolicyResults_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PolicyResultsServer is the server API for PolicyResults service.
// All implementations must embed UnimplementedPolicyResultsServer
// for forward compatibility
type PolicyResultsServer interface {
	List(context.Context, *PolicyResultFilter) (*PolicyResultList, error)
	mustEmbedUnimplementedPolicyResultsServer()
}

// UnimplementedPolicyResultsServer must be embedded to have forward compatible implementations.
type UnimplementedPolicyResultsServer struct {
}

func (UnimplementedPolicyResultsServer) List(context.Context, *PolicyResultFilter) (*PolicyResultList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedPolicyResultsServer) mustEmbedUnimplementedPolicyResultsServer() {}

// UnsafePolicyResultsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PolicyResultsServer will
// result in compilation errors.
type UnsafePolicyResultsServer interface {
	mustEmbedUnimplementedPolicyResultsServer()
}

func RegisterPolicyResultsServer(s grpc.ServiceRegistrar, srv PolicyResultsServer) {
	s.RegisterService(&PolicyResults_ServiceDesc, srv)
}

func _PolicyResults_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolicyResultFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyResultsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PolicyResults_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyResultsServer).List(ctx, req.(*PolicyResultFilter))
	}
	return interceptor(ctx, in, info, handler)
}

// PolicyResults_ServiceDesc is the grpc.ServiceDesc for PolicyResults service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PolicyResults_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chainguard.platform.tenant.PolicyResults",
	HandlerType: (*PolicyResultsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _PolicyResults_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "policy_results.platform.proto",
}