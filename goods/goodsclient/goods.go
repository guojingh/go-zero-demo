// Code generated by goctl. DO NOT EDIT.
// Source: goods.proto

package goodsclient

import (
	"context"

	"go-zero-demo/goods/types/goods"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GoodsRequest  = goods.GoodsRequest
	GoodsResponse = goods.GoodsResponse

	Goods interface {
		// rpc方法
		GetGoods(ctx context.Context, in *GoodsRequest, opts ...grpc.CallOption) (*GoodsResponse, error)
	}

	defaultGoods struct {
		cli zrpc.Client
	}
)

func NewGoods(cli zrpc.Client) Goods {
	return &defaultGoods{
		cli: cli,
	}
}

// rpc方法
func (m *defaultGoods) GetGoods(ctx context.Context, in *GoodsRequest, opts ...grpc.CallOption) (*GoodsResponse, error) {
	client := goods.NewGoodsClient(m.cli.Conn())
	return client.GetGoods(ctx, in, opts...)
}
