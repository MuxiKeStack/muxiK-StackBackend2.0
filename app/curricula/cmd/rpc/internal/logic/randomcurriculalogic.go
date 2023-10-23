package logic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/cmd/rpc/pb/pb"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/curricula/model"
	"math/rand"

	"github.com/zeromicro/go-zero/core/logx"
)

type RandomCurriculaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRandomCurriculaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RandomCurriculaLogic {
	return &RandomCurriculaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RandomCurriculaLogic) RandomCurricula(in *pb.RandomRequest) (*pb.RandomResponse, error) {
	// todo: test this method
	curriculas, err := l.svcCtx.CurriculaModel.FindMany(l.ctx, model.Cinfo{})
	if err != nil {
		return nil, err
	}
	n := rand.Intn(len(curriculas) - 5)
	var res []*pb.CurriculaInfo
	for i := 0; i < 5; i++ {
		info := &pb.CurriculaInfo{
			DataId:        curriculas[n].Id,
			CurriculaId:   uint32(curriculas[n].Cid),
			CurriculaName: curriculas[n].CurriculaName,
			Teacher:       curriculas[n].Teacher,
			Type:          uint32(curriculas[n].Type),
			Rate:          float32(curriculas[n].Rate.Float64),
		}

		res = append(res, info)
		n++
	}
	return &pb.RandomResponse{
		Info: res,
	}, nil
}
