package logic

import (
	"context"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchCurriculaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchCurriculaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchCurriculaLogic {
	return &SearchCurriculaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchCurriculaLogic) SearchCurricula(in *pb.SearchCurriculaRequest) (*pb.SearchCurriculaResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchCurriculaResponse{}, nil
}
