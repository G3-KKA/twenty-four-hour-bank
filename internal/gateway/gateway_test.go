package gateway

import (
	"bank24/internal/config"
	"bank24/internal/database/requests"
	log "bank24/internal/logger"
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
	"golang.org/x/sync/errgroup"
)

type gatewayTestSuite struct {
	suite.Suite
}

func TestRun(t *testing.T) {
	suite.Run(t, new(gatewayTestSuite))
}
func (t *gatewayTestSuite) SetupSuite() {
	viper.Set("WORKSPACE", "../..")
	viper.Set("CONFIG_FILE", "../../config.yaml")
	config.MustInitConfig()
	log.InitGlobalLogger(config.Get())

	// Run server
	gateway, err := New(context.TODO(), config.Get())
	if err != nil {
		log.Fatal("gateway initialisation failed. error: ", err)
	}
	go func() {

		errgroup, _ := errgroup.WithContext(context.TODO())
		errgroup.Go(
			func() error {
				log.Info("server started")
				return gateway.ListenAndServe(context.TODO(), config.Get())
			},
		)
		err = errgroup.Wait()
		if err != nil {
			log.Fatal("server shutdown with error: ", err)
		}
		log.Info("server shutdown successfully")
	}()

}

// ! TODO !
func (t *gatewayTestSuite) Test_Deposit() {

	time.Sleep(1 * time.Second)
	// CREATE ----------------------------------------------
	request := struct {
		Id      int     `json:"id"`
		Balance float64 `json:"balance"`
	}{
		Id:      0,
		Balance: 1000.0,
	}
	rbytes, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	buf := bytes.NewBuffer(rbytes)
	rsp, err := http.Post("http://localhost:8080/accounts", "application/json", buf)
	if err != nil {
		log.Fatal(err)
	}
	if rsp.StatusCode != 200 {
		log.Fatal("unexpected status code: ", rsp.StatusCode)
	}
	// WITHDRAW ------------------------------------------
	wrequest := struct {
		Id       int     `json:"id"`
		Withdraw float64 `json:"withdraw"`
	}{
		Id:       0,
		Withdraw: 100.0,
	}
	rbytes, err = json.Marshal(wrequest)
	if err != nil {
		log.Fatal(err)
	}
	buf = bytes.NewBuffer(rbytes)
	rsp, err = http.Post("http://localhost:8080/accounts/0/withdraw", "application/json", buf)
	if err != nil {
		log.Fatal(err)
	}
	if rsp.StatusCode != 200 {
		log.Fatal("unexpected status code: ", rsp.StatusCode)
	}
	// DEPOSIT --------------------------------------------
	drequest := struct {
		Id      int     `json:"id"`
		Deposit float64 `json:"deposit"`
	}{
		Id:      0,
		Deposit: 50.0,
	}
	rbytes, err = json.Marshal(drequest)
	if err != nil {
		log.Fatal(err)
	}
	buf = bytes.NewBuffer(rbytes)
	rsp, err = http.Post("http://localhost:8080/accounts/0/deposit", "application/json", buf)
	if err != nil {
		log.Fatal(err)
	}
	if rsp.StatusCode != 200 {
		log.Fatal("unexpected status code: ", rsp.StatusCode)
	}

	// BALANCE --------------------------------------------
	rsp, err = http.Get("http://localhost:8080/accounts/0/balance")
	if err != nil {
		log.Fatal(err)
	}
	if rsp.StatusCode != 200 {
		log.Fatal("unexpected status code: ", rsp.StatusCode)
	}
	balanceResp, err := io.ReadAll(rsp.Body)
	if err != nil {
		log.Fatal(err)
	}
	brsp := requests.BalanceResponse{}

	err = json.Unmarshal(balanceResp, &brsp)
	if err != nil {
		log.Fatal(err)
	}
	// create 1000
	// withdraw 100
	// deposit 50
	// balance should be 950
	if brsp.Data != 950.0 {
		log.Fatal("unexpected balance: ", brsp.Data)
	}

}
