// Copyright (c) 2025 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: gitpod/experimental/v1/workspaces.proto

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

// WorkspacesServiceClient is the client API for WorkspacesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkspacesServiceClient interface {
	// ListWorkspaces enumerates all workspaces belonging to the authenticated user.
	ListWorkspaces(ctx context.Context, in *ListWorkspacesRequest, opts ...grpc.CallOption) (*ListWorkspacesResponse, error)
	// GetWorkspace returns a single workspace.
	GetWorkspace(ctx context.Context, in *GetWorkspaceRequest, opts ...grpc.CallOption) (*GetWorkspaceResponse, error)
	// StreamWorkspaceStatus returns workspace status once it changed.
	StreamWorkspaceStatus(ctx context.Context, in *StreamWorkspaceStatusRequest, opts ...grpc.CallOption) (WorkspacesService_StreamWorkspaceStatusClient, error)
	// GetOwnerToken returns an owner token.
	GetOwnerToken(ctx context.Context, in *GetOwnerTokenRequest, opts ...grpc.CallOption) (*GetOwnerTokenResponse, error)
	// CreateAndStartWorkspace creates a new workspace and starts it.
	CreateAndStartWorkspace(ctx context.Context, in *CreateAndStartWorkspaceRequest, opts ...grpc.CallOption) (*CreateAndStartWorkspaceResponse, error)
	// StartWorkspace starts an existing workspace.
	StartWorkspace(ctx context.Context, in *StartWorkspaceRequest, opts ...grpc.CallOption) (*StartWorkspaceResponse, error)
	// StopWorkspace stops a running workspace (instance).
	// Errors:
	//
	//	NOT_FOUND:           the workspace_id is unkown
	//	FAILED_PRECONDITION: if there's no running instance
	StopWorkspace(ctx context.Context, in *StopWorkspaceRequest, opts ...grpc.CallOption) (*StopWorkspaceResponse, error)
	// DeleteWorkspace deletes a workspace.
	// When the workspace is running, it will be stopped as well.
	// Deleted workspaces cannot be started again.
	DeleteWorkspace(ctx context.Context, in *DeleteWorkspaceRequest, opts ...grpc.CallOption) (*DeleteWorkspaceResponse, error)
	UpdatePort(ctx context.Context, in *UpdatePortRequest, opts ...grpc.CallOption) (*UpdatePortResponse, error)
	// ListWorkspaceClasses enumerates all available workspace classes.
	ListWorkspaceClasses(ctx context.Context, in *ListWorkspaceClassesRequest, opts ...grpc.CallOption) (*ListWorkspaceClassesResponse, error)
	// GetDefaultWorkspaceImage returns the default workspace image from different sources.
	GetDefaultWorkspaceImage(ctx context.Context, in *GetDefaultWorkspaceImageRequest, opts ...grpc.CallOption) (*GetDefaultWorkspaceImageResponse, error)
}

type workspacesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkspacesServiceClient(cc grpc.ClientConnInterface) WorkspacesServiceClient {
	return &workspacesServiceClient{cc}
}

func (c *workspacesServiceClient) ListWorkspaces(ctx context.Context, in *ListWorkspacesRequest, opts ...grpc.CallOption) (*ListWorkspacesResponse, error) {
	out := new(ListWorkspacesResponse)
	err := c.cc.Invoke(ctx, "/gitpod.experimental.v1.WorkspacesService/ListWorkspaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspacesServiceClient) GetWorkspace(ctx context.Context, in *GetWorkspaceRequest, opts ...grpc.CallOption) (*GetWorkspaceResponse, error) {
	out := new(GetWorkspaceResponse)
	err := c.cc.Invoke(ctx, "/gitpod.experimental.v1.WorkspacesService/GetWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspacesServiceClient) StreamWorkspaceStatus(ctx context.Context, in *StreamWorkspaceStatusRequest, opts ...grpc.CallOption) (WorkspacesService_StreamWorkspaceStatusClient, error) {
	stream, err := c.cc.NewStream(ctx, &WorkspacesService_ServiceDesc.Streams[0], "/gitpod.experimental.v1.WorkspacesService/StreamWorkspaceStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &workspacesServiceStreamWorkspaceStatusClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WorkspacesService_StreamWorkspaceStatusClient interface {
	Recv() (*StreamWorkspaceStatusResponse, error)
	grpc.ClientStream
}

type workspacesServiceStreamWorkspaceStatusClient struct {
	grpc.ClientStream
}

func (x *workspacesServiceStreamWorkspaceStatusClient) Recv() (*StreamWorkspaceStatusResponse, error) {
	m := new(StreamWorkspaceStatusResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *workspacesServiceClient) GetOwnerToken(ctx context.Context, in *GetOwnerTokenRequest, opts ...grpc.CallOption) (*GetOwnerTokenResponse, error) {
	out := new(GetOwnerTokenResponse)
	err := c.cc.Invoke(ctx, "/gitpod.experimental.v1.WorkspacesService/GetOwnerToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspacesServiceClient) CreateAndStartWorkspace(ctx context.Context, in *CreateAndStartWorkspaceRequest, opts ...grpc.CallOption) (*CreateAndStartWorkspaceResponse, error) {
	out := new(CreateAndStartWorkspaceResponse)
	err := c.cc.Invoke(ctx, "/gitpod.experimental.v1.WorkspacesService/CreateAndStartWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspacesServiceClient) StartWorkspace(ctx context.Context, in *StartWorkspaceRequest, opts ...grpc.CallOption) (*StartWorkspaceResponse, error) {
	out := new(StartWorkspaceResponse)
	err := c.cc.Invoke(ctx, "/gitpod.experimental.v1.WorkspacesService/StartWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspacesServiceClient) StopWorkspace(ctx context.Context, in *StopWorkspaceRequest, opts ...grpc.CallOption) (*StopWorkspaceResponse, error) {
	out := new(StopWorkspaceResponse)
	err := c.cc.Invoke(ctx, "/gitpod.experimental.v1.WorkspacesService/StopWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspacesServiceClient) DeleteWorkspace(ctx context.Context, in *DeleteWorkspaceRequest, opts ...grpc.CallOption) (*DeleteWorkspaceResponse, error) {
	out := new(DeleteWorkspaceResponse)
	err := c.cc.Invoke(ctx, "/gitpod.experimental.v1.WorkspacesService/DeleteWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspacesServiceClient) UpdatePort(ctx context.Context, in *UpdatePortRequest, opts ...grpc.CallOption) (*UpdatePortResponse, error) {
	out := new(UpdatePortResponse)
	err := c.cc.Invoke(ctx, "/gitpod.experimental.v1.WorkspacesService/UpdatePort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspacesServiceClient) ListWorkspaceClasses(ctx context.Context, in *ListWorkspaceClassesRequest, opts ...grpc.CallOption) (*ListWorkspaceClassesResponse, error) {
	out := new(ListWorkspaceClassesResponse)
	err := c.cc.Invoke(ctx, "/gitpod.experimental.v1.WorkspacesService/ListWorkspaceClasses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workspacesServiceClient) GetDefaultWorkspaceImage(ctx context.Context, in *GetDefaultWorkspaceImageRequest, opts ...grpc.CallOption) (*GetDefaultWorkspaceImageResponse, error) {
	out := new(GetDefaultWorkspaceImageResponse)
	err := c.cc.Invoke(ctx, "/gitpod.experimental.v1.WorkspacesService/GetDefaultWorkspaceImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkspacesServiceServer is the server API for WorkspacesService service.
// All implementations must embed UnimplementedWorkspacesServiceServer
// for forward compatibility
type WorkspacesServiceServer interface {
	// ListWorkspaces enumerates all workspaces belonging to the authenticated user.
	ListWorkspaces(context.Context, *ListWorkspacesRequest) (*ListWorkspacesResponse, error)
	// GetWorkspace returns a single workspace.
	GetWorkspace(context.Context, *GetWorkspaceRequest) (*GetWorkspaceResponse, error)
	// StreamWorkspaceStatus returns workspace status once it changed.
	StreamWorkspaceStatus(*StreamWorkspaceStatusRequest, WorkspacesService_StreamWorkspaceStatusServer) error
	// GetOwnerToken returns an owner token.
	GetOwnerToken(context.Context, *GetOwnerTokenRequest) (*GetOwnerTokenResponse, error)
	// CreateAndStartWorkspace creates a new workspace and starts it.
	CreateAndStartWorkspace(context.Context, *CreateAndStartWorkspaceRequest) (*CreateAndStartWorkspaceResponse, error)
	// StartWorkspace starts an existing workspace.
	StartWorkspace(context.Context, *StartWorkspaceRequest) (*StartWorkspaceResponse, error)
	// StopWorkspace stops a running workspace (instance).
	// Errors:
	//
	//	NOT_FOUND:           the workspace_id is unkown
	//	FAILED_PRECONDITION: if there's no running instance
	StopWorkspace(context.Context, *StopWorkspaceRequest) (*StopWorkspaceResponse, error)
	// DeleteWorkspace deletes a workspace.
	// When the workspace is running, it will be stopped as well.
	// Deleted workspaces cannot be started again.
	DeleteWorkspace(context.Context, *DeleteWorkspaceRequest) (*DeleteWorkspaceResponse, error)
	UpdatePort(context.Context, *UpdatePortRequest) (*UpdatePortResponse, error)
	// ListWorkspaceClasses enumerates all available workspace classes.
	ListWorkspaceClasses(context.Context, *ListWorkspaceClassesRequest) (*ListWorkspaceClassesResponse, error)
	// GetDefaultWorkspaceImage returns the default workspace image from different sources.
	GetDefaultWorkspaceImage(context.Context, *GetDefaultWorkspaceImageRequest) (*GetDefaultWorkspaceImageResponse, error)
	mustEmbedUnimplementedWorkspacesServiceServer()
}

// UnimplementedWorkspacesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWorkspacesServiceServer struct {
}

func (UnimplementedWorkspacesServiceServer) ListWorkspaces(context.Context, *ListWorkspacesRequest) (*ListWorkspacesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkspaces not implemented")
}
func (UnimplementedWorkspacesServiceServer) GetWorkspace(context.Context, *GetWorkspaceRequest) (*GetWorkspaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkspace not implemented")
}
func (UnimplementedWorkspacesServiceServer) StreamWorkspaceStatus(*StreamWorkspaceStatusRequest, WorkspacesService_StreamWorkspaceStatusServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamWorkspaceStatus not implemented")
}
func (UnimplementedWorkspacesServiceServer) GetOwnerToken(context.Context, *GetOwnerTokenRequest) (*GetOwnerTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOwnerToken not implemented")
}
func (UnimplementedWorkspacesServiceServer) CreateAndStartWorkspace(context.Context, *CreateAndStartWorkspaceRequest) (*CreateAndStartWorkspaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAndStartWorkspace not implemented")
}
func (UnimplementedWorkspacesServiceServer) StartWorkspace(context.Context, *StartWorkspaceRequest) (*StartWorkspaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartWorkspace not implemented")
}
func (UnimplementedWorkspacesServiceServer) StopWorkspace(context.Context, *StopWorkspaceRequest) (*StopWorkspaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopWorkspace not implemented")
}
func (UnimplementedWorkspacesServiceServer) DeleteWorkspace(context.Context, *DeleteWorkspaceRequest) (*DeleteWorkspaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorkspace not implemented")
}
func (UnimplementedWorkspacesServiceServer) UpdatePort(context.Context, *UpdatePortRequest) (*UpdatePortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePort not implemented")
}
func (UnimplementedWorkspacesServiceServer) ListWorkspaceClasses(context.Context, *ListWorkspaceClassesRequest) (*ListWorkspaceClassesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkspaceClasses not implemented")
}
func (UnimplementedWorkspacesServiceServer) GetDefaultWorkspaceImage(context.Context, *GetDefaultWorkspaceImageRequest) (*GetDefaultWorkspaceImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDefaultWorkspaceImage not implemented")
}
func (UnimplementedWorkspacesServiceServer) mustEmbedUnimplementedWorkspacesServiceServer() {}

// UnsafeWorkspacesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkspacesServiceServer will
// result in compilation errors.
type UnsafeWorkspacesServiceServer interface {
	mustEmbedUnimplementedWorkspacesServiceServer()
}

func RegisterWorkspacesServiceServer(s grpc.ServiceRegistrar, srv WorkspacesServiceServer) {
	s.RegisterService(&WorkspacesService_ServiceDesc, srv)
}

func _WorkspacesService_ListWorkspaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkspacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspacesServiceServer).ListWorkspaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.experimental.v1.WorkspacesService/ListWorkspaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspacesServiceServer).ListWorkspaces(ctx, req.(*ListWorkspacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspacesService_GetWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspacesServiceServer).GetWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.experimental.v1.WorkspacesService/GetWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspacesServiceServer).GetWorkspace(ctx, req.(*GetWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspacesService_StreamWorkspaceStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamWorkspaceStatusRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WorkspacesServiceServer).StreamWorkspaceStatus(m, &workspacesServiceStreamWorkspaceStatusServer{stream})
}

type WorkspacesService_StreamWorkspaceStatusServer interface {
	Send(*StreamWorkspaceStatusResponse) error
	grpc.ServerStream
}

type workspacesServiceStreamWorkspaceStatusServer struct {
	grpc.ServerStream
}

func (x *workspacesServiceStreamWorkspaceStatusServer) Send(m *StreamWorkspaceStatusResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _WorkspacesService_GetOwnerToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOwnerTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspacesServiceServer).GetOwnerToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.experimental.v1.WorkspacesService/GetOwnerToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspacesServiceServer).GetOwnerToken(ctx, req.(*GetOwnerTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspacesService_CreateAndStartWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAndStartWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspacesServiceServer).CreateAndStartWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.experimental.v1.WorkspacesService/CreateAndStartWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspacesServiceServer).CreateAndStartWorkspace(ctx, req.(*CreateAndStartWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspacesService_StartWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspacesServiceServer).StartWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.experimental.v1.WorkspacesService/StartWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspacesServiceServer).StartWorkspace(ctx, req.(*StartWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspacesService_StopWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspacesServiceServer).StopWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.experimental.v1.WorkspacesService/StopWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspacesServiceServer).StopWorkspace(ctx, req.(*StopWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspacesService_DeleteWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWorkspaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspacesServiceServer).DeleteWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.experimental.v1.WorkspacesService/DeleteWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspacesServiceServer).DeleteWorkspace(ctx, req.(*DeleteWorkspaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspacesService_UpdatePort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspacesServiceServer).UpdatePort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.experimental.v1.WorkspacesService/UpdatePort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspacesServiceServer).UpdatePort(ctx, req.(*UpdatePortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspacesService_ListWorkspaceClasses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkspaceClassesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspacesServiceServer).ListWorkspaceClasses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.experimental.v1.WorkspacesService/ListWorkspaceClasses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspacesServiceServer).ListWorkspaceClasses(ctx, req.(*ListWorkspaceClassesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkspacesService_GetDefaultWorkspaceImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDefaultWorkspaceImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspacesServiceServer).GetDefaultWorkspaceImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitpod.experimental.v1.WorkspacesService/GetDefaultWorkspaceImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspacesServiceServer).GetDefaultWorkspaceImage(ctx, req.(*GetDefaultWorkspaceImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkspacesService_ServiceDesc is the grpc.ServiceDesc for WorkspacesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkspacesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gitpod.experimental.v1.WorkspacesService",
	HandlerType: (*WorkspacesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListWorkspaces",
			Handler:    _WorkspacesService_ListWorkspaces_Handler,
		},
		{
			MethodName: "GetWorkspace",
			Handler:    _WorkspacesService_GetWorkspace_Handler,
		},
		{
			MethodName: "GetOwnerToken",
			Handler:    _WorkspacesService_GetOwnerToken_Handler,
		},
		{
			MethodName: "CreateAndStartWorkspace",
			Handler:    _WorkspacesService_CreateAndStartWorkspace_Handler,
		},
		{
			MethodName: "StartWorkspace",
			Handler:    _WorkspacesService_StartWorkspace_Handler,
		},
		{
			MethodName: "StopWorkspace",
			Handler:    _WorkspacesService_StopWorkspace_Handler,
		},
		{
			MethodName: "DeleteWorkspace",
			Handler:    _WorkspacesService_DeleteWorkspace_Handler,
		},
		{
			MethodName: "UpdatePort",
			Handler:    _WorkspacesService_UpdatePort_Handler,
		},
		{
			MethodName: "ListWorkspaceClasses",
			Handler:    _WorkspacesService_ListWorkspaceClasses_Handler,
		},
		{
			MethodName: "GetDefaultWorkspaceImage",
			Handler:    _WorkspacesService_GetDefaultWorkspaceImage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamWorkspaceStatus",
			Handler:       _WorkspacesService_StreamWorkspaceStatus_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "gitpod/experimental/v1/workspaces.proto",
}
