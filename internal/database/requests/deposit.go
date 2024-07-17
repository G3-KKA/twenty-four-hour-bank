package requests

import "bank24/domain"

func (t DepositRequest) ValidRequest() {}

type DepositRequest struct {
	Id      domain.UserID
	Deposit float64
}
type DepositResponse struct{}

func Deposit(id domain.UserID, amount float64) DepositRequest {
	return DepositRequest{
		Id:      id,
		Deposit: amount,
	}
}
