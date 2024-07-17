package domain

type UserID uint64

type IdentifiebleBankAccount interface {
	BankAccount
	Id() UserID
}

type BankAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
}
