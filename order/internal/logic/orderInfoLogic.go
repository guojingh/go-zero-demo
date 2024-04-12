package logic

import (
	"context"

	"go-zero-demo/goods/types/goods"
	"go-zero-demo/order/internal/svc"
	"go-zero-demo/order/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取订单信息
func NewOrderInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderInfoLogic {
	return &OrderInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderInfoLogic) OrderInfo(req *types.OrderInfoReq) (resp *types.OrderInfoResp, err error) {
	// todo: add your logic here and delete this line
	orderId := req.OrderId
	goodRequest := new(goods.GoodsRequest)
	goodRequest.GoodsId = 25
	goodsInfo, err := l.svcCtx.Goods.GetGoods(l.ctx, goodRequest)

	if err != nil {
		return nil, err
	}

	resp = new(types.OrderInfoResp)
	resp.GoodsName = goodsInfo.Name
	resp.OrderId = orderId

	return
}
