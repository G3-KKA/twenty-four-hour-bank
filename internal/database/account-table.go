package database

import (
	"bank24/domain"
	"fmt"
	"sync"
)

type AccountTable interface {
	Insert(domain.IdentifiebleBankAccount) error
	UpadateBalance(domain.UserID, float64) error
	SelectBalance(domain.UserID) (float64, error)
}
type accountTable struct {
	table map[domain.UserID]domain.IdentifiebleBankAccount
	mx    *sync.RWMutex
}

func (t *accountTable) Insert(acc domain.IdentifiebleBankAccount) error {
	t.mx.Lock()
	id := acc.Id()
	defer t.mx.Unlock()
	if _, ok := t.table[id]; ok {
		return fmt.Errorf("account [%d] already exists", id)
	}
	t.table[id] = acc
	return nil

}

func (t *accountTable) UpadateBalance(id domain.UserID, amount float64) (err error) {
	t.mx.Lock()
	defer t.mx.Unlock()
	if _, ok := t.table[id]; !ok {
		return fmt.Errorf("account [%d] not found", id)
	}
	if amount < 0 {
		err = t.table[id].Withdraw(-amount)
		return
	}
	if amount > 0 {
		err = t.table[id].Deposit(amount)
		return
	}
	return

}
func (t *accountTable) SelectBalance(id domain.UserID) (float64, error) {

	t.mx.RLock()
	defer t.mx.RUnlock()
	if _, ok := t.table[id]; !ok {
		return -1.0, fmt.Errorf("account [%d] not found", id)
	}
	return t.table[id].GetBalance(), nil
}
