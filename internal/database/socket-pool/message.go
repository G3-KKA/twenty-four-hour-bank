package socketpool

import "bank24/internal/database/requests"

func NewMessage(req any) message {
	return message{request: req, returnChan: make(chan requests.Response, 1)}
}

type message struct {
	returnChan chan requests.Response
	request    any
}
