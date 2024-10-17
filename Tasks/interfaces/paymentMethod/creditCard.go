package paymentMethod

import "fmt"

type CreditCard struct {
	Balance     float64
	CreditLimit float64
}

func (c *CreditCard) CheckBalance() float64 {
	return c.Balance
}

func (c *CreditCard) Pay(amount float64) string {
	if c.Balance > amount {
		c.Balance -= amount
		fmt.Printf("Balance CreditCard:%2.f\n", c.Balance)
		return "Оплата прошла успешно с помощью CreditCard"
	} else {
		return "Недостаточно средств в CreditCard"
	}
}
