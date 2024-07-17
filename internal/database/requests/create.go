package requests

import "bank24/domain"

func (t CreateRequest) ValidRequest() {}

type CreateRequest struct {
	Id      domain.UserID
	Balance float64
}
type CreateResponse struct{}

func NewAccount(id domain.UserID, balance float64) CreateRequest {
	return CreateRequest{
		Id:      id,
		Balance: balance,
	}

}
