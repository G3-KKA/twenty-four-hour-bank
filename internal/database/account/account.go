package account

import (
	"bank24/domain"
	log "bank24/internal/logger"
	"fmt"
	"sync"
)

type Account struct {
	id      domain.UserID
	balance float64
	mx      *sync.RWMutex
}

func New(id domain.UserID, balance float64) *Account {
	return &Account{
		id:      id,
		balance: balance,
		mx:      &sync.RWMutex{},
	}
}

// Deposits amount into the bank, that this account is part of.
func (acc *Account) Deposit(amount float64) error {

	acc.mx.Lock()
	defer acc.mx.Unlock()
	if amount < 0 {
		log.Info(fmt.Sprintf("account [%d] unsuccessfully deposited: %f", acc.id, amount))
		return fmt.Errorf("depositing negative amount is not allowed")
	}
	acc.balance += amount
	log.Info(acc.id, "  account successfully deposited\t", amount)
	return nil
}

// Returns the account's Balance
func (acc *Account) GetBalance() float64 {

	acc.mx.RLock()
	// defer's are performance costly
	// but we need to ensure that balance not changed during the execution
	defer acc.mx.RUnlock()
	log.Info(fmt.Sprintf("account [%d] accesed it's balance", acc.id))
	return acc.balance
}

// Withdraws amount from the bank, that this account is part of.
func (acc *Account) Withdraw(amount float64) error {
	acc.mx.Lock()
	defer acc.mx.Unlock()

	if acc.balance < amount {
		log.Info(fmt.Sprintf("account [%d] unsuccessfully withdrawn: %f", acc.id, amount))
		return fmt.Errorf("not enough money to withdraw")
	}
	acc.balance -= amount
	log.Info(fmt.Sprintf("account [%d] successfully withdrawn: %f", acc.id, amount))
	return nil
}
func (acc *Account) Id() domain.UserID {
	return acc.id
}
