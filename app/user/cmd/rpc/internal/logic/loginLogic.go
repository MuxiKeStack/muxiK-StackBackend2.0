package logic

import (
	"context"
	"fmt"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/rpc/internal/svc"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/cmd/rpc/pb"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/app/user/model"
	"github.com/MuxiKeStack/muxiK-StackBackend2.0/common/ccnu_one_login"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginRequest) (*pb.LoginResponse, error) {
	if err := ccnu_one_login.LoginRequest(in.StudentID, in.Password); err != nil {
		return nil, err
	}
	var user *model.User
	var err error
	user, err = l.svcCtx.UserModel.FindOneBySid(l.ctx, in.StudentID)
	if err != nil {
		fmt.Println(err, user)
	}
	//if err != nil {
	//	return &pb.LoginResponse{
	//		Token: "1111",
	//		IfNew: false,
	//	}, nil
	//} else {
	//	user = &model.User{
	//		Sid:       in.StudentID,
	//		Username:  sql.NullString{},
	//		Avatar:    sql.NullString{},
	//		IsBlocked: 0,
	//		RoleGrade: 0,
	//		Integral:  0,
	//		Licence:   0,
	//	}
	//}
	//if err := l.svcCtx.UserModel.Insert(l.ctx, nil, user); err != nil {
	//
	//}
	return nil, err

}
