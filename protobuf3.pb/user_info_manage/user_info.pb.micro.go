// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user_info_manage/user_info.proto

package user_info_manage

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserService interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*RegisterResponse, error)
	Auth(ctx context.Context, in *AuthRequest, opts ...client.CallOption) (*AuthResponse, error)
	GetProfile(ctx context.Context, in *GetProfileRequest, opts ...client.CallOption) (*GetProfileResponse, error)
	GetHeadImage(ctx context.Context, in *GetHeadImageRequest, opts ...client.CallOption) (*GetHeadImageResponse, error)
	EditProfile(ctx context.Context, in *EditProfileRequest, opts ...client.CallOption) (*EditProfileResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "user.info.manage"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*RegisterResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.Register", in)
	out := new(RegisterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Auth(ctx context.Context, in *AuthRequest, opts ...client.CallOption) (*AuthResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.Auth", in)
	out := new(AuthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetProfile(ctx context.Context, in *GetProfileRequest, opts ...client.CallOption) (*GetProfileResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.GetProfile", in)
	out := new(GetProfileResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetHeadImage(ctx context.Context, in *GetHeadImageRequest, opts ...client.CallOption) (*GetHeadImageResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.GetHeadImage", in)
	out := new(GetHeadImageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) EditProfile(ctx context.Context, in *EditProfileRequest, opts ...client.CallOption) (*EditProfileResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.EditProfile", in)
	out := new(EditProfileResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	Register(context.Context, *RegisterRequest, *RegisterResponse) error
	Auth(context.Context, *AuthRequest, *AuthResponse) error
	GetProfile(context.Context, *GetProfileRequest, *GetProfileResponse) error
	GetHeadImage(context.Context, *GetHeadImageRequest, *GetHeadImageResponse) error
	EditProfile(context.Context, *EditProfileRequest, *EditProfileResponse) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		Register(ctx context.Context, in *RegisterRequest, out *RegisterResponse) error
		Auth(ctx context.Context, in *AuthRequest, out *AuthResponse) error
		GetProfile(ctx context.Context, in *GetProfileRequest, out *GetProfileResponse) error
		GetHeadImage(ctx context.Context, in *GetHeadImageRequest, out *GetHeadImageResponse) error
		EditProfile(ctx context.Context, in *EditProfileRequest, out *EditProfileResponse) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) Register(ctx context.Context, in *RegisterRequest, out *RegisterResponse) error {
	return h.UserServiceHandler.Register(ctx, in, out)
}

func (h *userServiceHandler) Auth(ctx context.Context, in *AuthRequest, out *AuthResponse) error {
	return h.UserServiceHandler.Auth(ctx, in, out)
}

func (h *userServiceHandler) GetProfile(ctx context.Context, in *GetProfileRequest, out *GetProfileResponse) error {
	return h.UserServiceHandler.GetProfile(ctx, in, out)
}

func (h *userServiceHandler) GetHeadImage(ctx context.Context, in *GetHeadImageRequest, out *GetHeadImageResponse) error {
	return h.UserServiceHandler.GetHeadImage(ctx, in, out)
}

func (h *userServiceHandler) EditProfile(ctx context.Context, in *EditProfileRequest, out *EditProfileResponse) error {
	return h.UserServiceHandler.EditProfile(ctx, in, out)
}