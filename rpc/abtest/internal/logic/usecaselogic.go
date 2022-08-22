package logic

import (
	"context"

	"shorturl/rpc/abtest/abtest"
	"shorturl/rpc/abtest/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UsecaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUsecaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UsecaseLogic {
	return &UsecaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//go run transform.go -f etc/transform.yaml
// go run shorturl.go -f etc/shorturl-api.yaml
func (l *UsecaseLogic) Usecase(in *abtest.AbReq) (*abtest.AbRsp, error) {
	//业务逻辑代码开始
	return &abtest.AbRsp{
		UserCase: "用户注册流程",
		Plan:     "A",
	}, nil
	//业务逻辑代码结束
}
