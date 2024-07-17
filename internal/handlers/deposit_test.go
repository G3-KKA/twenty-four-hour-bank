package handlers

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type depositTestSuite struct {
	suite.Suite
}

// NOT WORKING, logger is not initialized
func TestRun(t *testing.T) {
	suite.Run(t, new(depositTestSuite))
}
func (t *depositTestSuite) SetupSuite() {
}

// ! TODO  not actulaly testing !
func (t *depositTestSuite) Test_Deposit() {
	// TODO nil body
	is := &Interstate{}
	req := httptest.NewRequest("POST", "/accounts/2/deposit", nil)
	w := httptest.NewRecorder()
	is.Deposit(w, req)

}
