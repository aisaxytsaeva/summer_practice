package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	owner   string
	balance float64
}

func NewBankAccount(owner string, initialBalance float64) BankAccount {
	return BankAccount{
		owner:   owner,
		balance: initialBalance,
	}
}

// Пополнение счета
func (acc *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("сумма должна быть положительной")
	}
	acc.balance += amount
	return nil
}

// Снятие средств
func (acc *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("сумма должна быть положительной")
	}
	if amount > acc.balance {
		return errors.New("недостаточно средств")
	}
	acc.balance -= amount
	return nil
}

// Проверка баланса
func (acc BankAccount) GetBalance() float64 {
	return acc.balance
}

// Вывод информации
func (acc BankAccount) PrintInfo() {
	fmt.Printf("Владелец: %s\n", acc.owner)
	fmt.Printf("Баланс: %.2f ₽\n", acc.GetBalance())
}

func main() {
	account := NewBankAccount("Алексей Иванов", 1000)

	account.PrintInfo()

	// Пополнение
	if err := account.Deposit(500); err != nil {
		fmt.Println("Ошибка пополнения:", err)
	} else {
		fmt.Println("Пополнено 500 ₽")
	}

	// Снятие
	if err := account.Withdraw(200); err != nil {
		fmt.Println("Ошибка снятия:", err)
	} else {
		fmt.Println("Снято 200 ₽")
	}

	// Неудачное снятие
	if err := account.Withdraw(2000); err != nil {
		fmt.Println("Ошибка снятия:", err)
	}

	account.PrintInfo()
}
