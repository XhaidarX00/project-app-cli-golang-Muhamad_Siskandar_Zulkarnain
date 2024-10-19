package main

import (
	"errors"
	"fmt"
	"os"

	"strconv"

	"main/module"
	"main/utils"
)

// inisialisasi struct Bankaccount dan Tempaccount
var accounts = module.NewBankAccountManager()

// menu customer

// Menu untuk membuat akun baru
func createAccountMenu() {
	var name string
	var pin string
	fmt.Print("Masukkan nama akun baru : ")
	fmt.Scan(&name)
	for {
		fmt.Print("Buat pin rahasia (harus 4 angka) : ")
		fmt.Scan(&pin)
		utils.ClearScreen()
		_, err := strconv.Atoi(pin)
		if err != nil {
			err := errors.New("pin harus berupa angka")
			utils.ErrorMessage(err)
			continue
		}

		if accounts.IsLenVar(pin, 4) {
			break
		} else {
			err := errors.New("PIN harus terdiri dari 4 angka. Silakan coba lagi")
			utils.ErrorMessage(err)
		}
	}

	CapName := utils.Capitalize(name)
	accounts.CreateAccount(CapName, pin)
}

// Menu untuk mengecek saldo
func checkBalanceMenu() {
	account := accounts.PrintDataAccount("customer")
	if account != nil {
		module.CheckBalance(account)
	}
}

// Menu untuk deposit saldo
func depositMenu() {
	var amount float64
	account := accounts.PrintDataAccount("customer")
	if account != nil {
		fmt.Print("Masukkan jumlah deposit : ")
		fmt.Scan(&amount)
		accounts.Deposit(account.GetID(), amount)
	}
}

// Menu untuk withdraw saldo
func withdrawMenu() {
	var amount float64
	account := accounts.PrintDataAccount("customer")
	if account != nil {
		fmt.Print("Masukkan jumlah yang akan ditarik : ")
		fmt.Scan(&amount)
		accounts.Withdraw(account.GetID(), amount)
	}
}

func exitMainmenu() {
	defer os.Exit(0)
	utils.ClearScreen()
	utils.SuccesMessage("Keluar dari Program\n")
}

// menu mode dev

// untuk delete akun
func deleteAccountMenu() {
	account := accounts.PrintDataAccount("dev")
	if account != nil {
		accounts.DeleteAccount(account.GetID())
	}
}

// untuk blokir akun
func blokirAccountMenu() {
	account := accounts.PrintDataAccount("dev")
	if account != nil {
		accounts.BlockAccount(account.GetID())
	}
}

// untuk check akun blokir
func checkBlockMenu() {
	accounts.PrintBlockAccount()
}

// untuk memulihkan akun
func restoreMenu() {
	accounts.RestoreBlockAccount()
}
