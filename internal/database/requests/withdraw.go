package requests

import "bank24/domain"

func (t WithdrawRequest) ValidRequest() {}

type WithdrawRequest struct {
	Id       domain.UserID
	Withdraw float64
}
type WithdrawResponse struct{}

func Withdraw(id domain.UserID, amount float64) WithdrawRequest {
	return WithdrawRequest{
		Id:       id,
		Withdraw: amount,
	}
}
