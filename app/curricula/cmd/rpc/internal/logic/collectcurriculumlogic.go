package logic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/model"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollectCurriculumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCollectCurriculumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectCurriculumLogic {
	return &CollectCurriculumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CollectCurriculumLogic) CollectCurriculum(in *pb.CollectCurriculumRequest) (*pb.CollectCurriculumResponse, error) {
	// todo: Test this method
	_, err := l.svcCtx.CollectionModel.Insert(l.ctx, &model.Collections{
		Sid:     in.SID,
		CDataId: in.DataId,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CollectCurriculumResponse{}, nil
}
