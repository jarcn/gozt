package logic

import (
	"context"

	"shorturl/api/internal/svc"
	"shorturl/api/internal/types"
	"shorturl/rpc/abtest/abtest"

	"github.com/zeromicro/go-zero/core/logx"
)

type AbLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAbLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AbLogic {
	return &AbLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AbLogic) Ab(req *types.AbReq) (types.AbRsp, error) {
	// 手动处理rpc接口调用
	resp, err := l.svcCtx.AbTest.Usecase(l.ctx, &abtest.AbReq{})
	if err != nil {
		return types.AbRsp{}, err
	}
	return types.AbRsp{
		Plan:     resp.Plan,
		UserCase: resp.UserCase}, nil
}
