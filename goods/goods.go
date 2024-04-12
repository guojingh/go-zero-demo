package main

import (
	"flag"
	"fmt"
	"net"

	"go-zero-demo/goods/internal/config"
	"go-zero-demo/goods/internal/server"
	"go-zero-demo/goods/internal/svc"
	"go-zero-demo/goods/types/goods"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/goods.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	//获取动态接口
	//port, _ := GetFreePort()
	//替换 yaml 里面的 host 和 端口
	//fmt.Printf("port=%d\n", port)
	//c.ListenOn = fmt.Sprintf("0.0.0.0:%d", port)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		goods.RegisterGoodsServer(grpcServer, server.NewGoodsServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	//把服务信息注册到consul
	_ = consul.RegisterService(c.ListenOn, c.Consul)
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

// GetFreePort 动态获取可用端口
func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "192.168.222.136:0")
	if err != nil {
		return 0, err
	}

	fmt.Println(addr.Port)

	l, err := net.Listen("tcp", addr.String())
	if err != nil {
		return 0, err
	}

	return l.Addr().(*net.TCPAddr).Port, nil
}
