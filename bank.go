package main

import (
	"errors"
	"fmt"
)

// BankAccount представляет банковский счет
type BankAccount struct {
	owner   string  // Владелец счета
	balance float64 // Текущий баланс
}

// NewBankAccount создает новый банковский счет
func NewBankAccount(owner string, initialBalance float64) *BankAccount {
	return &BankAccount{
		owner:   owner,
		balance: initialBalance,
	}
}

// Deposit пополняет счет на указанную сумму
func (acc *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("сумма пополнения должна быть положительной")
	}
	acc.balance += amount
	return nil
}

// Withdraw снимает деньги со счета
func (acc *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("сумма снятия должна быть положительной")
	}
	if amount > acc.balance {
		return errors.New("недостаточно средств на счете")
	}
	acc.balance -= amount
	return nil
}

// Balance возвращает текущий баланс
func (acc *BankAccount) Balance() float64 {
	return acc.balance
}

// Owner возвращает владельца счета
func (acc *BankAccount) Owner() string {
	return acc.owner
}

func main() {
	// Создаем два счета
	account1 := NewBankAccount("Иван Иванов", 1000)
	account2 := NewBankAccount("Петр Петров", 500)

	// Выводим информацию о счетах
	fmt.Printf("Счет %s: %.2f руб.\n", account1.Owner(), account1.Balance())
	fmt.Printf("Счет %s: %.2f руб.\n", account2.Owner(), account2.Balance())

	// Пополняем счет
	if err := account1.Deposit(300); err != nil {
		fmt.Println("Ошибка пополнения:", err)
	} else {
		fmt.Println("Счет успешно пополнен")
	}

	// Снимаем деньги
	if err := account2.Withdraw(200); err != nil {
		fmt.Println("Ошибка снятия:", err)
	} else {
		fmt.Println("Снятие успешно выполнено")
	}

	// Пытаемся снять слишком много
	if err := account1.Withdraw(2000); err != nil {
		fmt.Println("Ошибка снятия:", err)
	}

	// Выводим итоговые балансы
	fmt.Printf("\nИтоговый баланс:\n")
	fmt.Printf("Счет %s: %.2f руб.\n", account1.Owner(), account1.Balance())
	fmt.Printf("Счет %s: %.2f руб.\n", account2.Owner(), account2.Balance())
}
