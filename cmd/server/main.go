package main

import (
	"bank24/internal/config"
	"bank24/internal/gateway"
	log "bank24/internal/logger"
	"context"

	stdlog "log"

	"golang.org/x/sync/errgroup"
)

const HARDCODED_SOCKET_COUNT = 3

// SOMEWHERE write method to make shure that in Message.returnChan
// - - value was written exactly once, channel is initialised and with capacity 1

func main() {
	// I am not a big fun of init() func

	config.MustInitConfig()
	initialConfig := config.Get()
	err := log.InitGlobalLogger(initialConfig)
	if err != nil {
		stdlog.Fatalf("logger initialisation falied. error: %e", err)
	}
	//
	// context with cancel or so from config
	//
	// should also accept context
	//
	gateway, err := gateway.New(context.TODO(), initialConfig)
	if err != nil {
		log.Fatal("gateway initialisation failed. error: ", err)
	}
	errgroup, _ := errgroup.WithContext(context.TODO())
	errgroup.Go(
		func() error {
			return gateway.ListenAndServe(context.TODO(), initialConfig)
		},
	)
	err = errgroup.Wait()
	if err != nil {
		log.Fatal("server shutdown with error: ", err)
	}
	log.Info("server shutdown successfully")
}

const CONFIGSocketBufferSize = 5
const CONFIGSockerCount = 3
const CONFIGDatabaseSize = 2048

//
