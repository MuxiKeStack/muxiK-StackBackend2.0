package logic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/pb/pb"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/common/ctxdata"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollectCurriculumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCollectCurriculumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectCurriculumLogic {
	return &CollectCurriculumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CollectCurriculumLogic) CollectCurriculum(req *types.CollectCurriculumRequest) (resp *types.CollectCurriculumResponse, err error) {
	// todo: Test this method
	sid := ctxdata.GetStudentIdFromCtx(l.ctx)
	_, err = l.svcCtx.CurriculaCenterRpc.CollectCurriculum(l.ctx, &pb.CollectCurriculumRequest{
		DataId: req.DataId,
		SID:    sid,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
