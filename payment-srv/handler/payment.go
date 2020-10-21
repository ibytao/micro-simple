package handler

import (
	"context"
	"micro-simple/payment-srv/model/payment"
	proto "micro-simple/payment-srv/proto/payment"

	log "github.com/micro/go-micro/v2/logger"
)

var (
	paymentService payment.Service
)

type Service struct{}

// Init 初始化handler
func Init() {
	paymentService, _ = payment.GetService()
}

// New 新增订单
func (e *Service) PayOrder(ctx context.Context, req *proto.Request, rsp *proto.Response) (err error) {
	log.Infof("[PayOrder] 收到支付请求")
	err = paymentService.PayOrder(req.OrderId)

	if err != nil {
		rsp.Success = false
		rsp.Error = &proto.Error{
			Detail: err.Error(),
		}
		return
	}

	rsp.Success = true
	return
}
