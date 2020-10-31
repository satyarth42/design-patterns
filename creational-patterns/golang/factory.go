package golang

import (
	"errors"
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float32) string
}

const (
	Cash = 1
	Card = 2
)

type CashPM struct {
}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash", amount)
}

type CardPM struct {
}

func (c *CardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using card", amount)
}

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case Card:
		return new(CardPM), nil
	default:
		return nil, errors.New(fmt.Sprintf("No payment method for %d", m))
	}
}
