// Code generated by goctl. DO NOT EDIT.
// Source: evaluation.proto

package eva_like

import (
	"context"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateEvaluationRequest = pb.CreateEvaluationRequest
	DeleteEvaluationRequest = pb.DeleteEvaluationRequest
	Evaluation              = pb.Evaluation
	EvaluationResponse      = pb.EvaluationResponse
	GetEvaluationRequest    = pb.GetEvaluationRequest
	GetLikeRequest          = pb.GetLikeRequest
	GetLikeResponse         = pb.GetLikeResponse
	GetReportsRequest       = pb.GetReportsRequest
	GetReportsResponse      = pb.GetReportsResponse
	GetTheReportRequest     = pb.GetTheReportRequest
	GetTheReportResponse    = pb.GetTheReportResponse
	Like                    = pb.Like
	Report                  = pb.Report
	SendReportRequest       = pb.SendReportRequest
	SetLikeRequest          = pb.SetLikeRequest
	SetReportRequest        = pb.SetReportRequest
	StatusResponse          = pb.StatusResponse
	UpdateEvaluationRequest = pb.UpdateEvaluationRequest

	EvaLike interface {
		SetLike(ctx context.Context, in *SetLikeRequest, opts ...grpc.CallOption) (*StatusResponse, error)
		GetLike(ctx context.Context, in *GetLikeRequest, opts ...grpc.CallOption) (*GetLikeResponse, error)
	}

	defaultEvaLike struct {
		cli zrpc.Client
	}
)

func NewEvaLike(cli zrpc.Client) EvaLike {
	return &defaultEvaLike{
		cli: cli,
	}
}

func (m *defaultEvaLike) SetLike(ctx context.Context, in *SetLikeRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	client := pb.NewEvaLikeClient(m.cli.Conn())
	return client.SetLike(ctx, in, opts...)
}

func (m *defaultEvaLike) GetLike(ctx context.Context, in *GetLikeRequest, opts ...grpc.CallOption) (*GetLikeResponse, error) {
	client := pb.NewEvaLikeClient(m.cli.Conn())
	return client.GetLike(ctx, in, opts...)
}
