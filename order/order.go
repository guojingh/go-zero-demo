package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"go-zero-demo/order/internal/config"
	"go-zero-demo/order/internal/handler"
	"go-zero-demo/order/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

var configFile = flag.String("f", "etc/order-api.yaml", "the config file")

// 注意要导入 consul 依赖进行支持 _ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"   切记....
func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	ctxStr, _ := json.Marshal(ctx)
	fmt.Printf("ctxStr=%s", ctxStr)

	//全局中间件
	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			httpx.OkJson(w, "不能随便进入啊")
			//return
			next(w, r)
		}
	})

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
