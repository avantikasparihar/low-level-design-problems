package pkg

import "fmt"

type Account struct {
	Id           int
	CustomerName string
	Type         string
	Balance      int
	Pin          int
}

type BankService interface {
	ValidatePin(cardNumber, pin int) bool
	Debit(amount int)
	Credit(amount int)
}

func NewBankService(cardNumber int) BankService {
	return NewSbiBankService()
}

type sbiBankService struct {
}

func (s sbiBankService) ValidatePin(cardNumber, pin int) bool {
	return true
}

func (s sbiBankService) Debit(amount int) {
	fmt.Println("...Sbi bank...")
	fmt.Printf("amount %d debited\n", amount)
}

func (s sbiBankService) Credit(amount int) {
	fmt.Println("...Sbi bank...")
	fmt.Printf("amount %d credited\n", amount)
}

func NewSbiBankService() BankService {
	return &sbiBankService{}
}
