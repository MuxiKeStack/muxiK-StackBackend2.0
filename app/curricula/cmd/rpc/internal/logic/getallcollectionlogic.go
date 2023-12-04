package logic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllCollectionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllCollectionLogic {
	return &GetAllCollectionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllCollectionLogic) GetAllCollection(in *pb.GetAllCollectionRequest) (*pb.GetAllCollectionResponse, error) {
	// todo: test this method
	collections, err := l.svcCtx.CollectionModel.FindManyByUid(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	var resp *pb.GetAllCollectionResponse
	for _, collection := range collections {
		v, err := l.svcCtx.CurriculaModel.FindOne(l.ctx, collection.CDataId)
		if err != nil {
			return nil, err
		}
		curricula := &pb.CurriculaInfo{
			DataId:        v.Id,
			CurriculaId:   uint32(v.Cid),
			CurriculaName: v.CurriculaName,
			Teacher:       v.Teacher,
			Type:          uint32(v.Type),
			Rate:          float32(v.Rate.Float64),
		}
		resp.Info = append(resp.Info, curricula)
	}
	return resp, nil
}
