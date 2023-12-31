// Code generated by goctl. DO NOT EDIT.
// Source: evaluation.proto

package server

import (
	"context"

	evalikelogic "github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/internal/logic/eva_like"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/pb/pb"
)

type EvaLikeServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedEvaLikeServer
}

func NewEvaLikeServer(svcCtx *svc.ServiceContext) *EvaLikeServer {
	return &EvaLikeServer{
		svcCtx: svcCtx,
	}
}

func (s *EvaLikeServer) SetLike(ctx context.Context, in *pb.SetLikeRequest) (*pb.StatusResponse, error) {
	l := evalikelogic.NewSetLikeLogic(ctx, s.svcCtx)
	return l.SetLike(in)
}

func (s *EvaLikeServer) GetLike(ctx context.Context, in *pb.GetLikeRequest) (*pb.GetLikeResponse, error) {
	l := evalikelogic.NewGetLikeLogic(ctx, s.svcCtx)
	return l.GetLike(in)
}
