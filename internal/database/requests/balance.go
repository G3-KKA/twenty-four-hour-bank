package requests

import "bank24/domain"

func (t BalanceRequest) ValidRequest() {}

type BalanceRequest struct {
	Id domain.UserID
}
type BalanceResponse struct {
	Data float64 `json:"data"`
}

func Balance(id domain.UserID) BalanceRequest {
	return BalanceRequest{
		Id: id,
	}
}
