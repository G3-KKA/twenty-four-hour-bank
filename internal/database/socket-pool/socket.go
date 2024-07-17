package socketpool

import (
	"bank24/internal/database"
	"bank24/internal/database/account"
	"bank24/internal/database/requests"
	log "bank24/internal/logger"
	"context"
)

type socket struct {
	ch chan message
	db database.BankDatabase
}

func NewSocket(ch chan message, db database.BankDatabase) socket {
	return socket{ch: ch, db: db}
}
func (s *socket) Serve(ctx context.Context) {
	for {
		msg := message{}
		ok := false
		select {
		case <-ctx.Done():
			// TODO graceful shutdown
			log.Fatal("unimplemented graceful shutdown")
		case msg, ok = <-s.ch:
		}
		if !ok {
			log.Fatal("unimplemented closed channel handling")
		}
		if msg.returnChan == nil {
			log.Info("received nil return channel")
			continue
		}
		go func(msg message) {
			defer close(msg.returnChan)
			switch req := msg.request.(type) {
			case requests.DepositRequest:
				err := s.db.UpadateBalance(req.Id, req.Deposit)
				rsp := requests.Response{Err: err}
				msg.returnChan <- rsp

			case requests.WithdrawRequest:
				err := s.db.UpadateBalance(req.Id, -req.Withdraw)
				rsp := requests.Response{Err: err}
				msg.returnChan <- rsp

			case requests.BalanceRequest:
				acc, err := s.db.SelectBalance(req.Id)
				rsp := requests.Response{
					Data: requests.BalanceResponse{Data: acc},
					Err:  err,
				}
				msg.returnChan <- rsp

			case requests.CreateRequest:

				err := s.db.Insert(account.New(req.Id, req.Balance))
				rsp := requests.Response{Err: err}
				msg.returnChan <- rsp
			default:
				log.Fatal("unimplemented request type")
			}
		}(msg)
	}

}
