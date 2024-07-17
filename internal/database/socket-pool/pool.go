package socketpool

import (
	"bank24/internal/config"
	"bank24/internal/database"
	"bank24/internal/database/requests"
	log "bank24/internal/logger"
	"context"
	"math/rand/v2"
)

// TODO move to config!!!
const CONFIGSocketBufferSize = 5
const CONFIGSockerCount = 3
const CONFIGDatabaseSize = 2048

// Multi-Connection Pool
type SocketPool struct {
	sockets []socket
}
type DatabaseRequestHandler interface {
	Send(req requests.ValidRequest) requests.Response
}

func (s *SocketPool) Send(req requests.ValidRequest) requests.Response {
	i := rand.Uint32() % CONFIGSockerCount
	returnChan := make(chan requests.Response, 1)
	s.sockets[i].ch <- message{request: req, returnChan: returnChan}
	return <-returnChan
}
func NewSocketPool(ctx context.Context, config config.Config) SocketPool {
	// its insane that socket pool creates database by himself
	db := database.NewBankDatabase(ctx, config)
	err := db.Init(config)
	if err != nil {
		log.Fatal("database initialisation failed. error: ", err)

	}
	sockets := make([]socket, 0, CONFIGSockerCount)

	for range CONFIGSockerCount {
		sockets = append(sockets, NewSocket(make(chan message, CONFIGSocketBufferSize), db))
	}
	for _, s := range sockets {
		go s.Serve(context.TODO())
	}
	pool := SocketPool{sockets: sockets}
	return pool
}
