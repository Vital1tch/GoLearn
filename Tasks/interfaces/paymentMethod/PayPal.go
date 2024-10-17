package paymentMethod

import "fmt"

type PayPal struct {
	Balance float64
}

func (p *PayPal) CheckBalance() float64 {
	return p.Balance
}

func (p *PayPal) Pay(amount float64) string {
	if p.Balance > amount {
		p.Balance -= amount
		fmt.Printf("Balance PayPal:%2.f\n", p.Balance)
		return "Оплата прошла успешно с помощью PayPal"
	} else {
		return "Недостаточно средств в PayPal"
	}
}
