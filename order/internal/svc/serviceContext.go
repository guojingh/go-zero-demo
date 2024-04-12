package svc

import (
	"go-zero-demo/goods/goodsclient"
	"go-zero-demo/order/internal/config"
	"go-zero-demo/order/internal/middleware"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	//定义 RPC 服务
	Goods goodsclient.Goods
	Login rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		//引入 grpc 服务
		Goods: goodsclient.NewGoods(zrpc.MustNewClient(c.GoodsRpc)),
		Login: middleware.NewLoginMiddleware().Handle,
	}
}
