package internal

import (
	"fmt"
	"github.com/avantikasparihar/low-level-design-problems/atm/pkg"
)

type Card struct {
	Number int
	Cvv    int
	Pin    int
	Type   string
}

type CardReader interface {
	DisplayOptions()
	Authenticate(pin int) bool
}

type cardReader struct {
	bank pkg.BankService
	card Card
}

func (c cardReader) DisplayOptions() {
	fmt.Printf("Card Details: %+v\n", c.card)
}

func (c cardReader) Authenticate(pin int) bool {
	return c.bank.ValidatePin(c.card.Number, pin)
}

func NewCardReader(card Card) CardReader {
	var bank pkg.BankService
	switch getBankNameFromCardNo(card.Number) {
	case "sbi":
		bank = pkg.NewSbiBankService()
	}
	return &cardReader{
		bank: bank,
		card: card,
	}
}

func getBankNameFromCardNo(cardNo int) string {
	return "sbi"
}
