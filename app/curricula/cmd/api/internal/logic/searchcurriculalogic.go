package logic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/pb/pb"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchCurriculaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchCurriculaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchCurriculaLogic {
	return &SearchCurriculaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchCurriculaLogic) SearchCurricula(req *types.SearchCurriculaRequest) (resp *types.SearchCurriculaResponse, err error) {
	// todo: Test this method
	result, err := l.svcCtx.CurriculaCenterRpc.SearchCurricula(l.ctx, &pb.SearchCurriculaRequest{
		CurriculaId:   req.CurriculaId,
		CurriculaName: req.CurriculaName,
		Teacher:       req.Teacher,
		Type:          uint32(req.Type),
	})
	if err != nil {
		return &types.SearchCurriculaResponse{}, err
	}
	info := result.GetInfo()
	resp.Info[0] = types.CurriculaInfo{
		DataId:        info[0].DataId,
		CurriculaId:   info[0].CurriculaId,
		CurriculaName: info[0].CurriculaName,
		Teacher:       info[0].Teacher,
		Type:          info[0].Type,
		Rate:          info[0].Rate,
	}
	return resp, nil
}
