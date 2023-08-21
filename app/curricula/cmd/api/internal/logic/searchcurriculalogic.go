package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
