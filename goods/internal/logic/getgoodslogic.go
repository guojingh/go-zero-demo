package logic

import (
	"context"

	"go-zero-demo/goods/internal/svc"
	"go-zero-demo/goods/types/goods"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsLogic {
	return &GetGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// rpc方法
func (l *GetGoodsLogic) GetGoods(in *goods.GoodsRequest) (res *goods.GoodsResponse, err error) {
	// todo: add your logic here and delete this line
	//根据订单 id 获取商品信息
	goodsId := in.GoodsId
	res = new(goods.GoodsResponse)
	res.GoodsId = goodsId
	res.Name = "茅台" + l.svcCtx.Config.ListenOn
	return
}
