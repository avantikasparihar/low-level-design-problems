package main

import . "github.com/avantikasparihar/low-level-design-problems/atm/internal"

/*
Entities:
- Card
	number: int
	cvv: int
	validThrough: date
	cardHolder: string
	type: string
- Transaction
	id: uid
	cardNumber: int
	type: enum("DEPOSIT", "WITHDRAW", "ENQUIRY")
	status: enum("INPROGRESS", "SUCCESSFUL", "FAILED")
	amount: int
	denominations: map[type]int
	completedAt: time
- Account
	id: int
	customerName: string
	type: string
	balance: int
- CashInventory
	denominations: map[type]int

Interfaces:
- CardReader
	DisplayOptions()
	Authenticate()
- CashDepositor
	Deposit()
- CashDispenser
	CanDispense(amount)
	Dispense(amount)
- BankService()
	ValidatePin()
	Debit()
	Credit()
- ReceiptGenerator
	GetReceipt()
*/

func main() {
	atm := NewAtm()

	atm.InsertCard()
	atm.EnterPin()

	atm.WithdrawCash()
	atm.DepositCash()
}
