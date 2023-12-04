package logic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/pb/pb"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/common/ctxdata"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/api/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllCollectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllCollectionLogic {
	return &GetAllCollectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllCollectionLogic) GetAllCollection(req *types.GetAllCollectionRequest) (resp *types.GetAllCollectionResponse, err error) {
	// todo: test this method
	sid := ctxdata.GetStudentIdFromCtx(l.ctx)
	response, err := l.svcCtx.CurriculaCenterRpc.GetAllCollection(l.ctx, &pb.GetAllCollectionRequest{UserId: sid})
	if err != nil {
		return nil, err
	}
	for _, info := range response.Info {
		resp.Info = append(resp.Info, types.CurriculaInfo{
			DataId:        info.DataId,
			CurriculaId:   info.CurriculaId,
			CurriculaName: info.CurriculaName,
			Teacher:       info.Teacher,
			Type:          uint8(info.Type),
			Rate:          info.Rate,
		})
	}
	return resp, err
}
