package evainfologic

import (
	"context"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/model"
	"strconv"

	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/evaluation/cmd/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEvaluationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEvaluationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEvaluationLogic {
	return &GetEvaluationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetEvaluationLogic) GetEvaluation(in *pb.GetEvaluationRequest) (*pb.EvaluationResponse, error) {
	// todo: testing
	var buf []model.EvaluationInfo
	id, _ := strconv.Atoi(in.Id)
	switch in.Method {
	case "pid":
		info, err := l.svcCtx.InfoModel.FindOneByPid(l.ctx, int64(id))
		if err != nil {
			return nil, err
		}
		buf = append(buf, *info)
	case "sid":
		info, err := l.svcCtx.InfoModel.GetEvaluationBySid(l.ctx, strconv.FormatInt(int64(id), 10))
		if err != nil {
			return nil, err
		}
		buf = info
	case "cid":
		info, err := l.svcCtx.InfoModel.GetEvaluationByCid(l.ctx, strconv.FormatInt(int64(id), 10))
		if err != nil {
			return nil, err
		}
		buf = info
	}

	var resp []*pb.Evaluation
	for _, v := range buf {
		resp = append(resp, &pb.Evaluation{
			Pid:      strconv.FormatInt(v.Pid, 10),
			Sid:      v.Sid,
			Cid:      v.Cid,
			Folded:   v.Folded > 0,
			Deleted:  v.Deleted > 0,
			Info:     v.Info,
			Liked:    int32(v.Liked.Int64),
			Disliked: int32(v.Disliked.Int64),
			CreateAt: v.CreatedAt.Time.String(),
			UpdateAt: v.UpdatedAt.Time.String(),
			DeleteAt: v.DeletedAt.Time.String(),
		},
		)
	}

	return &pb.EvaluationResponse{E: resp}, nil
}
