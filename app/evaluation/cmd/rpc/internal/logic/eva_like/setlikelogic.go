package evalikelogic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/model"
	"strconv"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetLikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetLikeLogic {
	return &SetLikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SetLike : set like status and automatically update the liked/disliked number
func (l *SetLikeLogic) SetLike(in *pb.SetLikeRequest) (*pb.StatusResponse, error) {
	// todo: testing
	pid, _ := strconv.Atoi(in.Pid)
	info, err := l.svcCtx.LikeModel.FindLike(l.ctx, in.Pid, in.Sid)
	if err != nil {
		l.svcCtx.LikeModel.Insert(l.ctx, &model.EvaluationLike{
			Pid:    int64(pid),
			Sid:    in.Sid,
			Status: 0,
		})
	}
	if info.Status == int64(in.Status) {
		return &pb.StatusResponse{Status: false}, nil
	}

	eva, err := l.svcCtx.InfoModel.FindOneByPid(l.ctx, int64(pid))
	if err != nil {
		return nil, err
	}

	if info.Status == 1 {
		eva.Liked.Int64--
	} else if info.Status == -1 {
		eva.Disliked.Int64--
	}

	info.Status = int64(in.Status)

	if info.Status == 1 {
		eva.Liked.Int64++
	} else if info.Status == -1 {
		eva.Disliked.Int64++
	}

	err = l.svcCtx.InfoModel.UpdateEvaluationInfo(l.ctx, eva)
	if err != nil {
		return nil, err
	}

	return &pb.StatusResponse{Status: true}, nil
}
