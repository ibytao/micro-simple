package orders

// New 新增订单
func (s *service) New(bookId int64, userId int64) (orderId int64, err error) {
	// 请求销存
	rsp, err := invClient.Sell(context.TODO(), &invS.Request{
		BookId: bookId, UserId: userId
	})
}
