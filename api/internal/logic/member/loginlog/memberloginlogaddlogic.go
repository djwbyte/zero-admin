package logic

import (
	"context"
	"go-zero-admin/api/internal/common/errorx"
	"go-zero-admin/rpc/ums/umsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberLoginLogAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberLoginLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberLoginLogAddLogic {
	return MemberLoginLogAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberLoginLogAddLogic) MemberLoginLogAdd(req types.AddMemberLoginLogReq) (*types.AddMemberLoginLogResp, error) {
	_, err := l.svcCtx.Ums.MemberLoginLogAdd(l.ctx, &umsclient.MemberLoginLogAddReq{
		MemberId:   req.MemberId,
		CreateTime: req.CreateTime,
		Ip:         req.Ip,
		City:       req.City,
		LoginType:  req.LoginType,
		Province:   req.Province,
	})

	if err != nil {
		return nil, errorx.NewDefaultError("添加会员登录记录失败")
	}

	return &types.AddMemberLoginLogResp{
		Code:    "000000",
		Message: "添加员登录记录成功",
	}, nil
}
