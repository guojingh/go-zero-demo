package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	//定义 RPC 服务
	GoodsRpc zrpc.RpcClientConf
}
