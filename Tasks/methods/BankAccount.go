package methods

import (
	"errors"
	"fmt"
)

type BankAccount struct { //Создаём структуру
	balance      float64
	limit        float64
	countOfCalls int //Счётчик обращений за балансом
}

func NewBankAccount(balance, limit float64) *BankAccount { //Функция для создания нового банковского аккаунта
	return &BankAccount{
		balance: balance,
		limit:   limit,
	}
}

func (a *BankAccount) Deposit(amount float64) { //Закинуть деньги  (Метод)
	a.balance = a.balance + amount
}

func (a *BankAccount) Withdraw(amount float64) (float64, error) { //Снять деньги (Метод)
	if a.balance < amount {
		return 0, errors.New("Не хватает денег для снятия")
	}
	a.balance -= amount
	return a.balance, nil
}

func (a *BankAccount) Balance() { //Метод
	a.countOfCalls += 1 //Обновляем счётчик обращений за балансом
	fmt.Printf("Баланс аккаунта: %2.f\n Номер обращения: %d\n\n", a.balance, a.countOfCalls)
}
