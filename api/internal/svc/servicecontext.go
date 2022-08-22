package svc

import (
	"shorturl/api/internal/config"
	"shorturl/rpc/abtest/usecase"
	"shorturl/rpc/transform/transformer"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	Transformer transformer.Transformer //手动代码
	AbTest      usecase.Usecase         //手动代码
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		Transformer: transformer.NewTransformer(zrpc.MustNewClient(c.Transform)), // 手动代码
		AbTest:      usecase.NewUsecase(zrpc.MustNewClient(c.AbTest)),            // 手动代码
	}
}
