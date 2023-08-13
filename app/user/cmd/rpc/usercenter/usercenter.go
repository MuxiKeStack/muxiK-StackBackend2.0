// Code generated by goctl. DO NOT EDIT.
// Source: usercenter.proto

package usercenter

import (
	"context"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GenerateTokenRequest     = pb.GenerateTokenRequest
	GenerateTokenResponse    = pb.GenerateTokenResponse
	GetInfoRequest           = pb.GetInfoRequest
	GetInfoResponse          = pb.GetInfoResponse
	IncreaseIntegralRequest  = pb.IncreaseIntegralRequest
	IncreaseIntegralResponse = pb.IncreaseIntegralResponse
	LoginRequest             = pb.LoginRequest
	LoginResponse            = pb.LoginResponse
	SharingRequest           = pb.SharingRequest
	SharingResponse          = pb.SharingResponse
	UpdateInfoRequest        = pb.UpdateInfoRequest
	UpdateInfoResponse       = pb.UpdateInfoResponse

	Usercenter interface {
		Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		GenerateToken(ctx context.Context, in *GenerateTokenRequest, opts ...grpc.CallOption) (*GenerateTokenResponse, error)
		UpdateInfo(ctx context.Context, in *UpdateInfoRequest, opts ...grpc.CallOption) (*UpdateInfoResponse, error)
		SharingPlan(ctx context.Context, in *SharingRequest, opts ...grpc.CallOption) (*SharingResponse, error)
		GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error)
		IncreaseIntegral(ctx context.Context, in *IncreaseIntegralRequest, opts ...grpc.CallOption) (*IncreaseIntegralResponse, error)
	}

	defaultUsercenter struct {
		cli zrpc.Client
	}
)

func NewUsercenter(cli zrpc.Client) Usercenter {
	return &defaultUsercenter{
		cli: cli,
	}
}

func (m *defaultUsercenter) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUsercenter) GenerateToken(ctx context.Context, in *GenerateTokenRequest, opts ...grpc.CallOption) (*GenerateTokenResponse, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.GenerateToken(ctx, in, opts...)
}

func (m *defaultUsercenter) UpdateInfo(ctx context.Context, in *UpdateInfoRequest, opts ...grpc.CallOption) (*UpdateInfoResponse, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.UpdateInfo(ctx, in, opts...)
}

func (m *defaultUsercenter) SharingPlan(ctx context.Context, in *SharingRequest, opts ...grpc.CallOption) (*SharingResponse, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.SharingPlan(ctx, in, opts...)
}

func (m *defaultUsercenter) GetInfo(ctx context.Context, in *GetInfoRequest, opts ...grpc.CallOption) (*GetInfoResponse, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.GetInfo(ctx, in, opts...)
}

func (m *defaultUsercenter) IncreaseIntegral(ctx context.Context, in *IncreaseIntegralRequest, opts ...grpc.CallOption) (*IncreaseIntegralResponse, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.IncreaseIntegral(ctx, in, opts...)
}