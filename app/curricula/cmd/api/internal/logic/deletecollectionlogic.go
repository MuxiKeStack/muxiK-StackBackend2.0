package logic

import (
	"context"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCollectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCollectionLogic {
	return &DeleteCollectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCollectionLogic) DeleteCollection(req *types.DeleteCollectionRequest) (resp *types.DeleteCollectionResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
