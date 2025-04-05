// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.6
// source: v1/menu/menu.proto

package menuv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	MenuService_CreateCategory_FullMethodName       = "/menu.MenuService/CreateCategory"
	MenuService_UpdateCategory_FullMethodName       = "/menu.MenuService/UpdateCategory"
	MenuService_ListCategories_FullMethodName       = "/menu.MenuService/ListCategories"
	MenuService_DeleteCategory_FullMethodName       = "/menu.MenuService/DeleteCategory"
	MenuService_CreateDish_FullMethodName           = "/menu.MenuService/CreateDish"
	MenuService_UpdateDish_FullMethodName           = "/menu.MenuService/UpdateDish"
	MenuService_GetDish_FullMethodName              = "/menu.MenuService/GetDish"
	MenuService_ListDishes_FullMethodName           = "/menu.MenuService/ListDishes"
	MenuService_CreateMenu_FullMethodName           = "/menu.MenuService/CreateMenu"
	MenuService_GetActiveMenu_FullMethodName        = "/menu.MenuService/GetActiveMenu"
	MenuService_UpdateMenu_FullMethodName           = "/menu.MenuService/UpdateMenu"
	MenuService_CreatePromotion_FullMethodName      = "/menu.MenuService/CreatePromotion"
	MenuService_UpdatePromotion_FullMethodName      = "/menu.MenuService/UpdatePromotion"
	MenuService_ListActivePromotions_FullMethodName = "/menu.MenuService/ListActivePromotions"
	MenuService_GenerateUploadUrl_FullMethodName    = "/menu.MenuService/GenerateUploadUrl"
)

// MenuServiceClient is the client API for MenuService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MenuServiceClient interface {
	// Категории
	CreateCategory(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error)
	UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error)
	ListCategories(ctx context.Context, in *ListCategoriesRequest, opts ...grpc.CallOption) (*ListCategoriesResponse, error)
	DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Блюда
	CreateDish(ctx context.Context, in *DishRequest, opts ...grpc.CallOption) (*DishResponse, error)
	UpdateDish(ctx context.Context, in *UpdateDishRequest, opts ...grpc.CallOption) (*DishResponse, error)
	GetDish(ctx context.Context, in *GetDishRequest, opts ...grpc.CallOption) (*DishResponse, error)
	ListDishes(ctx context.Context, in *ListDishesRequest, opts ...grpc.CallOption) (*ListDishesResponse, error)
	// Меню
	CreateMenu(ctx context.Context, in *MenuRequest, opts ...grpc.CallOption) (*MenuResponse, error)
	GetActiveMenu(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*MenuResponse, error)
	UpdateMenu(ctx context.Context, in *UpdateMenuRequest, opts ...grpc.CallOption) (*MenuResponse, error)
	// Акции
	CreatePromotion(ctx context.Context, in *PromotionRequest, opts ...grpc.CallOption) (*PromotionResponse, error)
	UpdatePromotion(ctx context.Context, in *UpdatePromotionRequest, opts ...grpc.CallOption) (*PromotionResponse, error)
	ListActivePromotions(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListPromotionsResponse, error)
	// Изображения
	GenerateUploadUrl(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageResponse, error)
}

type menuServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMenuServiceClient(cc grpc.ClientConnInterface) MenuServiceClient {
	return &menuServiceClient{cc}
}

func (c *menuServiceClient) CreateCategory(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CategoryResponse)
	err := c.cc.Invoke(ctx, MenuService_CreateCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) UpdateCategory(ctx context.Context, in *UpdateCategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CategoryResponse)
	err := c.cc.Invoke(ctx, MenuService_UpdateCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) ListCategories(ctx context.Context, in *ListCategoriesRequest, opts ...grpc.CallOption) (*ListCategoriesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCategoriesResponse)
	err := c.cc.Invoke(ctx, MenuService_ListCategories_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MenuService_DeleteCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) CreateDish(ctx context.Context, in *DishRequest, opts ...grpc.CallOption) (*DishResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DishResponse)
	err := c.cc.Invoke(ctx, MenuService_CreateDish_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) UpdateDish(ctx context.Context, in *UpdateDishRequest, opts ...grpc.CallOption) (*DishResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DishResponse)
	err := c.cc.Invoke(ctx, MenuService_UpdateDish_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) GetDish(ctx context.Context, in *GetDishRequest, opts ...grpc.CallOption) (*DishResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DishResponse)
	err := c.cc.Invoke(ctx, MenuService_GetDish_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) ListDishes(ctx context.Context, in *ListDishesRequest, opts ...grpc.CallOption) (*ListDishesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListDishesResponse)
	err := c.cc.Invoke(ctx, MenuService_ListDishes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) CreateMenu(ctx context.Context, in *MenuRequest, opts ...grpc.CallOption) (*MenuResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MenuResponse)
	err := c.cc.Invoke(ctx, MenuService_CreateMenu_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) GetActiveMenu(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*MenuResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MenuResponse)
	err := c.cc.Invoke(ctx, MenuService_GetActiveMenu_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) UpdateMenu(ctx context.Context, in *UpdateMenuRequest, opts ...grpc.CallOption) (*MenuResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MenuResponse)
	err := c.cc.Invoke(ctx, MenuService_UpdateMenu_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) CreatePromotion(ctx context.Context, in *PromotionRequest, opts ...grpc.CallOption) (*PromotionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PromotionResponse)
	err := c.cc.Invoke(ctx, MenuService_CreatePromotion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) UpdatePromotion(ctx context.Context, in *UpdatePromotionRequest, opts ...grpc.CallOption) (*PromotionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PromotionResponse)
	err := c.cc.Invoke(ctx, MenuService_UpdatePromotion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) ListActivePromotions(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListPromotionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPromotionsResponse)
	err := c.cc.Invoke(ctx, MenuService_ListActivePromotions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) GenerateUploadUrl(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ImageResponse)
	err := c.cc.Invoke(ctx, MenuService_GenerateUploadUrl_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MenuServiceServer is the server API for MenuService service.
// All implementations must embed UnimplementedMenuServiceServer
// for forward compatibility
type MenuServiceServer interface {
	// Категории
	CreateCategory(context.Context, *CategoryRequest) (*CategoryResponse, error)
	UpdateCategory(context.Context, *UpdateCategoryRequest) (*CategoryResponse, error)
	ListCategories(context.Context, *ListCategoriesRequest) (*ListCategoriesResponse, error)
	DeleteCategory(context.Context, *DeleteCategoryRequest) (*emptypb.Empty, error)
	// Блюда
	CreateDish(context.Context, *DishRequest) (*DishResponse, error)
	UpdateDish(context.Context, *UpdateDishRequest) (*DishResponse, error)
	GetDish(context.Context, *GetDishRequest) (*DishResponse, error)
	ListDishes(context.Context, *ListDishesRequest) (*ListDishesResponse, error)
	// Меню
	CreateMenu(context.Context, *MenuRequest) (*MenuResponse, error)
	GetActiveMenu(context.Context, *emptypb.Empty) (*MenuResponse, error)
	UpdateMenu(context.Context, *UpdateMenuRequest) (*MenuResponse, error)
	// Акции
	CreatePromotion(context.Context, *PromotionRequest) (*PromotionResponse, error)
	UpdatePromotion(context.Context, *UpdatePromotionRequest) (*PromotionResponse, error)
	ListActivePromotions(context.Context, *emptypb.Empty) (*ListPromotionsResponse, error)
	// Изображения
	GenerateUploadUrl(context.Context, *ImageRequest) (*ImageResponse, error)
	mustEmbedUnimplementedMenuServiceServer()
}

// UnimplementedMenuServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMenuServiceServer struct {
}

func (UnimplementedMenuServiceServer) CreateCategory(context.Context, *CategoryRequest) (*CategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}
func (UnimplementedMenuServiceServer) UpdateCategory(context.Context, *UpdateCategoryRequest) (*CategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCategory not implemented")
}
func (UnimplementedMenuServiceServer) ListCategories(context.Context, *ListCategoriesRequest) (*ListCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCategories not implemented")
}
func (UnimplementedMenuServiceServer) DeleteCategory(context.Context, *DeleteCategoryRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategory not implemented")
}
func (UnimplementedMenuServiceServer) CreateDish(context.Context, *DishRequest) (*DishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDish not implemented")
}
func (UnimplementedMenuServiceServer) UpdateDish(context.Context, *UpdateDishRequest) (*DishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDish not implemented")
}
func (UnimplementedMenuServiceServer) GetDish(context.Context, *GetDishRequest) (*DishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDish not implemented")
}
func (UnimplementedMenuServiceServer) ListDishes(context.Context, *ListDishesRequest) (*ListDishesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDishes not implemented")
}
func (UnimplementedMenuServiceServer) CreateMenu(context.Context, *MenuRequest) (*MenuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMenu not implemented")
}
func (UnimplementedMenuServiceServer) GetActiveMenu(context.Context, *emptypb.Empty) (*MenuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActiveMenu not implemented")
}
func (UnimplementedMenuServiceServer) UpdateMenu(context.Context, *UpdateMenuRequest) (*MenuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMenu not implemented")
}
func (UnimplementedMenuServiceServer) CreatePromotion(context.Context, *PromotionRequest) (*PromotionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePromotion not implemented")
}
func (UnimplementedMenuServiceServer) UpdatePromotion(context.Context, *UpdatePromotionRequest) (*PromotionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePromotion not implemented")
}
func (UnimplementedMenuServiceServer) ListActivePromotions(context.Context, *emptypb.Empty) (*ListPromotionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListActivePromotions not implemented")
}
func (UnimplementedMenuServiceServer) GenerateUploadUrl(context.Context, *ImageRequest) (*ImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateUploadUrl not implemented")
}
func (UnimplementedMenuServiceServer) mustEmbedUnimplementedMenuServiceServer() {}

// UnsafeMenuServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MenuServiceServer will
// result in compilation errors.
type UnsafeMenuServiceServer interface {
	mustEmbedUnimplementedMenuServiceServer()
}

func RegisterMenuServiceServer(s grpc.ServiceRegistrar, srv MenuServiceServer) {
	s.RegisterService(&MenuService_ServiceDesc, srv)
}

func _MenuService_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MenuService_CreateCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).CreateCategory(ctx, req.(*CategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_UpdateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).UpdateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MenuService_UpdateCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).UpdateCategory(ctx, req.(*UpdateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_ListCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).ListCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MenuService_ListCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).ListCategories(ctx, req.(*ListCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_DeleteCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).DeleteCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MenuService_DeleteCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).DeleteCategory(ctx, req.(*DeleteCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_CreateDish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).CreateDish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MenuService_CreateDish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).CreateDish(ctx, req.(*DishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_UpdateDish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).UpdateDish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MenuService_UpdateDish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).UpdateDish(ctx, req.(*UpdateDishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_GetDish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).GetDish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MenuService_GetDish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).GetDish(ctx, req.(*GetDishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_ListDishes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDishesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).ListDishes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MenuService_ListDishes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).ListDishes(ctx, req.(*ListDishesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_CreateMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).CreateMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MenuService_CreateMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).CreateMenu(ctx, req.(*MenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_GetActiveMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).GetActiveMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MenuService_GetActiveMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).GetActiveMenu(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_UpdateMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).UpdateMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MenuService_UpdateMenu_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).UpdateMenu(ctx, req.(*UpdateMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_CreatePromotion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromotionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).CreatePromotion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MenuService_CreatePromotion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).CreatePromotion(ctx, req.(*PromotionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_UpdatePromotion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePromotionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).UpdatePromotion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MenuService_UpdatePromotion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).UpdatePromotion(ctx, req.(*UpdatePromotionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_ListActivePromotions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).ListActivePromotions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MenuService_ListActivePromotions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).ListActivePromotions(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_GenerateUploadUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).GenerateUploadUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MenuService_GenerateUploadUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).GenerateUploadUrl(ctx, req.(*ImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MenuService_ServiceDesc is the grpc.ServiceDesc for MenuService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MenuService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "menu.MenuService",
	HandlerType: (*MenuServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCategory",
			Handler:    _MenuService_CreateCategory_Handler,
		},
		{
			MethodName: "UpdateCategory",
			Handler:    _MenuService_UpdateCategory_Handler,
		},
		{
			MethodName: "ListCategories",
			Handler:    _MenuService_ListCategories_Handler,
		},
		{
			MethodName: "DeleteCategory",
			Handler:    _MenuService_DeleteCategory_Handler,
		},
		{
			MethodName: "CreateDish",
			Handler:    _MenuService_CreateDish_Handler,
		},
		{
			MethodName: "UpdateDish",
			Handler:    _MenuService_UpdateDish_Handler,
		},
		{
			MethodName: "GetDish",
			Handler:    _MenuService_GetDish_Handler,
		},
		{
			MethodName: "ListDishes",
			Handler:    _MenuService_ListDishes_Handler,
		},
		{
			MethodName: "CreateMenu",
			Handler:    _MenuService_CreateMenu_Handler,
		},
		{
			MethodName: "GetActiveMenu",
			Handler:    _MenuService_GetActiveMenu_Handler,
		},
		{
			MethodName: "UpdateMenu",
			Handler:    _MenuService_UpdateMenu_Handler,
		},
		{
			MethodName: "CreatePromotion",
			Handler:    _MenuService_CreatePromotion_Handler,
		},
		{
			MethodName: "UpdatePromotion",
			Handler:    _MenuService_UpdatePromotion_Handler,
		},
		{
			MethodName: "ListActivePromotions",
			Handler:    _MenuService_ListActivePromotions_Handler,
		},
		{
			MethodName: "GenerateUploadUrl",
			Handler:    _MenuService_GenerateUploadUrl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/menu/menu.proto",
}
