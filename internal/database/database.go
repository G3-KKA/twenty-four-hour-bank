package database

import (
	"bank24/domain"
	"bank24/internal/config"
	"context"
	"sync"
)

type BankDatabase interface {
	AccountTable
	Database
}

type Database interface {
	Init(config.Config) error
}
type database struct {
	accountTable
}

const CONFIGSocketBufferSize = 5
const CONFIGSockerCount = 3
const CONFIGDatabaseSize = 2048

func (db *database) Init(config config.Config) error {
	db.accountTable = accountTable{table: make(map[domain.UserID]domain.IdentifiebleBankAccount, CONFIGDatabaseSize), mx: &sync.RWMutex{}}
	return nil
}
func NewBankDatabase(ctx context.Context, config config.Config) BankDatabase {
	db := database{}
	db.Init(config)
	return &db
}
