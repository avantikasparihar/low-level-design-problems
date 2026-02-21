package internal

import (
	"fmt"
	"github.com/avantikasparihar/low-level-design-problems/atm/pkg"
)

type Atm struct {
	idleState         AtmState
	cardInsertedState AtmState
	pinEnteredState   AtmState
	noCashState       AtmState

	currState AtmState
	cash      int
}

func (a *Atm) setCurrentState(state AtmState) {
	a.currState = state
}

func (a *Atm) InsertCard() {
	a.currState.InsertCard()
}

func (a *Atm) EnterPin() {
	a.currState.EnterPin()
}

func (a *Atm) WithdrawCash() {
	a.currState.WithdrawCash()
}

func (a *Atm) DepositCash() {
	a.currState.DepositCash()
}

func NewAtm() *Atm {
	atm := &Atm{
		cash: 1000,
	}
	atm.idleState = &idleState{atm: atm}
	atm.cardInsertedState = &cardInsertedState{atm: atm}
	atm.pinEnteredState = &pinEnteredState{atm: atm}

	atm.currState = atm.idleState
	return atm
}

type AtmState interface {
	InsertCard()
	EnterPin()
	WithdrawCash()
	DepositCash()
}

type idleState struct {
	atm *Atm
	cardReader
}

func (i idleState) InsertCard() {
	fmt.Println("card inserted")
	i.atm.setCurrentState(i.atm.cardInsertedState)
}

func (i idleState) EnterPin() {
	fmt.Println("card is not inserted")
}

func (i idleState) WithdrawCash() {
	fmt.Println("card is not inserted")
}

func (i idleState) DepositCash() {
	fmt.Println("card is not inserted")
}

type cardInsertedState struct {
	atm  *Atm
	bank pkg.BankService
}

func (c cardInsertedState) InsertCard() {
	fmt.Println("card is already inserted")
}

func (c cardInsertedState) EnterPin() {
	// init new bank service
	c.bank = pkg.NewBankService(0)
	// validate pin for card
	valid := c.bank.ValidatePin(0, 0)
	if valid {
		fmt.Println("valid pin entered")
		c.atm.setCurrentState(c.atm.pinEnteredState)
	} else {
		fmt.Println("invalid pin entered")
	}
}

func (c cardInsertedState) WithdrawCash() {
	fmt.Println("pin is not entered")
}

func (c cardInsertedState) DepositCash() {
	fmt.Println("pin is not entered")
}

type pinEnteredState struct {
	atm  *Atm
	bank pkg.BankService
}

func (p pinEnteredState) InsertCard() {
	fmt.Println("card is already inserted")
}

func (p pinEnteredState) EnterPin() {
	fmt.Println("pin is already entered")
}

func (p pinEnteredState) WithdrawCash() {
	// init new bank service
	p.bank = pkg.NewBankService(0)
	p.bank.Debit(0)
	fmt.Println("cash withdrawn")
}

func (p pinEnteredState) DepositCash() {
	// init new bank service
	p.bank = pkg.NewBankService(0)
	p.bank.Credit(0)
	fmt.Println("cash deposited")
}
