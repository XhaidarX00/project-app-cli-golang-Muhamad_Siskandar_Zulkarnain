package models

import "github.com/google/uuid"

type Account struct {
	ID      string
	Name    string
	Balance float64
	Pin     string
}

func (a *Account) Deposit(amount float64) {
	a.Balance += amount
}

func (a *Account) Withdraw(amount float64) bool {
	if a.Balance >= amount {
		a.Balance -= amount
		return true
	}
	return false
}

func NewAccount(name, pin string) *Account {
	return &Account{
		ID:      uuid.New().String(),
		Name:    name,
		Balance: 0,
		Pin:     pin,
	}
}

func (a *Account) GetID() string {
	return a.ID
}

func (a *Account) GetName() string {
	return a.Name
}

func (a *Account) GetBalance() float64 {
	return a.Balance
}

func (a *Account) GetPin() string {
	return a.Pin
}

func (a *Account) CheckPin(pin string) bool {
	return a.Pin == pin
}
