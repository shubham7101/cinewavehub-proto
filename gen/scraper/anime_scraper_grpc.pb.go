// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.1
// source: scraper/anime_scraper.proto

package scraper

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AnimeScraper_Info_FullMethodName                     = "/cinewavehub.scraper.AnimeScraper/Info"
	AnimeScraper_EpisodesAndMappings_FullMethodName      = "/cinewavehub.scraper.AnimeScraper/EpisodesAndMappings"
	AnimeScraper_EpisodesAndMappingsAsync_FullMethodName = "/cinewavehub.scraper.AnimeScraper/EpisodesAndMappingsAsync"
	AnimeScraper_Sources_FullMethodName                  = "/cinewavehub.scraper.AnimeScraper/Sources"
)

// AnimeScraperClient is the client API for AnimeScraper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnimeScraperClient interface {
	Info(ctx context.Context, in *AnimeInfoRequest, opts ...grpc.CallOption) (*AnimeInfoResponse, error)
	EpisodesAndMappings(ctx context.Context, in *AnimeEpisodesAndMappingsRequest, opts ...grpc.CallOption) (*AnimeEpisodesAndMappingsResponse, error)
	EpisodesAndMappingsAsync(ctx context.Context, in *AnimeEpisodesAndMappingsRequest, opts ...grpc.CallOption) (*ScraperResponse, error)
	Sources(ctx context.Context, in *SourcesRequest, opts ...grpc.CallOption) (*SourcesResponse, error)
}

type animeScraperClient struct {
	cc grpc.ClientConnInterface
}

func NewAnimeScraperClient(cc grpc.ClientConnInterface) AnimeScraperClient {
	return &animeScraperClient{cc}
}

func (c *animeScraperClient) Info(ctx context.Context, in *AnimeInfoRequest, opts ...grpc.CallOption) (*AnimeInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AnimeInfoResponse)
	err := c.cc.Invoke(ctx, AnimeScraper_Info_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *animeScraperClient) EpisodesAndMappings(ctx context.Context, in *AnimeEpisodesAndMappingsRequest, opts ...grpc.CallOption) (*AnimeEpisodesAndMappingsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AnimeEpisodesAndMappingsResponse)
	err := c.cc.Invoke(ctx, AnimeScraper_EpisodesAndMappings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *animeScraperClient) EpisodesAndMappingsAsync(ctx context.Context, in *AnimeEpisodesAndMappingsRequest, opts ...grpc.CallOption) (*ScraperResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ScraperResponse)
	err := c.cc.Invoke(ctx, AnimeScraper_EpisodesAndMappingsAsync_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *animeScraperClient) Sources(ctx context.Context, in *SourcesRequest, opts ...grpc.CallOption) (*SourcesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SourcesResponse)
	err := c.cc.Invoke(ctx, AnimeScraper_Sources_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnimeScraperServer is the server API for AnimeScraper service.
// All implementations must embed UnimplementedAnimeScraperServer
// for forward compatibility.
type AnimeScraperServer interface {
	Info(context.Context, *AnimeInfoRequest) (*AnimeInfoResponse, error)
	EpisodesAndMappings(context.Context, *AnimeEpisodesAndMappingsRequest) (*AnimeEpisodesAndMappingsResponse, error)
	EpisodesAndMappingsAsync(context.Context, *AnimeEpisodesAndMappingsRequest) (*ScraperResponse, error)
	Sources(context.Context, *SourcesRequest) (*SourcesResponse, error)
	mustEmbedUnimplementedAnimeScraperServer()
}

// UnimplementedAnimeScraperServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAnimeScraperServer struct{}

func (UnimplementedAnimeScraperServer) Info(context.Context, *AnimeInfoRequest) (*AnimeInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedAnimeScraperServer) EpisodesAndMappings(context.Context, *AnimeEpisodesAndMappingsRequest) (*AnimeEpisodesAndMappingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EpisodesAndMappings not implemented")
}
func (UnimplementedAnimeScraperServer) EpisodesAndMappingsAsync(context.Context, *AnimeEpisodesAndMappingsRequest) (*ScraperResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EpisodesAndMappingsAsync not implemented")
}
func (UnimplementedAnimeScraperServer) Sources(context.Context, *SourcesRequest) (*SourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sources not implemented")
}
func (UnimplementedAnimeScraperServer) mustEmbedUnimplementedAnimeScraperServer() {}
func (UnimplementedAnimeScraperServer) testEmbeddedByValue()                      {}

// UnsafeAnimeScraperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnimeScraperServer will
// result in compilation errors.
type UnsafeAnimeScraperServer interface {
	mustEmbedUnimplementedAnimeScraperServer()
}

func RegisterAnimeScraperServer(s grpc.ServiceRegistrar, srv AnimeScraperServer) {
	// If the following call pancis, it indicates UnimplementedAnimeScraperServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AnimeScraper_ServiceDesc, srv)
}

func _AnimeScraper_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnimeInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnimeScraperServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnimeScraper_Info_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnimeScraperServer).Info(ctx, req.(*AnimeInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnimeScraper_EpisodesAndMappings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnimeEpisodesAndMappingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnimeScraperServer).EpisodesAndMappings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnimeScraper_EpisodesAndMappings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnimeScraperServer).EpisodesAndMappings(ctx, req.(*AnimeEpisodesAndMappingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnimeScraper_EpisodesAndMappingsAsync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnimeEpisodesAndMappingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnimeScraperServer).EpisodesAndMappingsAsync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnimeScraper_EpisodesAndMappingsAsync_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnimeScraperServer).EpisodesAndMappingsAsync(ctx, req.(*AnimeEpisodesAndMappingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnimeScraper_Sources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnimeScraperServer).Sources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnimeScraper_Sources_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnimeScraperServer).Sources(ctx, req.(*SourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AnimeScraper_ServiceDesc is the grpc.ServiceDesc for AnimeScraper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AnimeScraper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cinewavehub.scraper.AnimeScraper",
	HandlerType: (*AnimeScraperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _AnimeScraper_Info_Handler,
		},
		{
			MethodName: "EpisodesAndMappings",
			Handler:    _AnimeScraper_EpisodesAndMappings_Handler,
		},
		{
			MethodName: "EpisodesAndMappingsAsync",
			Handler:    _AnimeScraper_EpisodesAndMappingsAsync_Handler,
		},
		{
			MethodName: "Sources",
			Handler:    _AnimeScraper_Sources_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scraper/anime_scraper.proto",
}
