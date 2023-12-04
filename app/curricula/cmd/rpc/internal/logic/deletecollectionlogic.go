package logic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/pb/pb"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCollectionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCollectionLogic {
	return &DeleteCollectionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCollectionLogic) DeleteCollection(in *pb.DeleteCollectionRequest) (*pb.DeleteCollectionResponse, error) {
	// todo: test this method
	collection, err := l.svcCtx.CollectionModel.FindByInfos(l.ctx, model.CollectionInfo{
		Sid: in.UserId,
		Cid: in.DataId,
	})
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.CollectionModel.Delete(l.ctx, collection.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteCollectionResponse{}, nil
}
