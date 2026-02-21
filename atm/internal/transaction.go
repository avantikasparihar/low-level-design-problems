package internal

import "time"

type Transactions struct {
	Id            int
	cardNumber    int
	Type          string
	Status        string
	Amount        int
	Denominations map[string]int
	CompletedAt   time.Time
}

type CashDepositor interface {
	Deposit()
}

type CashDispenser interface {
	Dispense()
}
