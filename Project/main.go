package main

import (
	"errors"
	"fmt"
	"main/utils"
)

// Aplikasi CLI Pengelolaan Akun dan Saldo Customer

// Customer mode
func main() {
	for {
		var input int
		fmt.Println("+++ ====== Menu ====== +++")
		fmt.Println("1. " + utils.ColorMessage("green", "Buat Akun"))
		fmt.Println("2. " + utils.ColorMessage("green", "Cek Saldo"))
		fmt.Println("3. " + utils.ColorMessage("green", "Deposit"))
		fmt.Println("4. " + utils.ColorMessage("green", "Penarikan"))
		fmt.Println("5. " + utils.ColorMessage("red", "Keluar"))
		fmt.Print("Masukan nomor menu : ")
		fmt.Scan(&input)
		utils.ClearScreen()

		switch input {
		case 1:
			utils.ClearScreen()
			createAccountMenu()
		case 2:
			utils.ClearScreen()
			checkBalanceMenu()
		case 3:
			utils.ClearScreen()
			depositMenu()
		case 4:
			utils.ClearScreen()
			withdrawMenu()
		case 5:
			exitMainmenu()
		case 777:
			utils.ClearScreen()
			DevModeMenu()
		default:
			err := errors.New("pilihan tidak valid")
			utils.ErrorMessage(err)
		}
	}
}

// Dev mode

func DevModeMenu() {
	for {
		var input int
		fmt.Println(utils.ColorMessage("blue", "+++ ====== Dev Mode Menu ====== +++"))
		fmt.Println("1. " + utils.ColorMessage("green", "Hapus Akun"))
		fmt.Println("2. " + utils.ColorMessage("green", "Blokir Akun"))
		fmt.Println("3. " + utils.ColorMessage("green", "Cek Akun Blokir"))
		fmt.Println("4. " + utils.ColorMessage("green", "Pulihkan Akun"))
		fmt.Println("0. " + utils.ColorMessage("red", "Untuk Keluar"))
		fmt.Print(utils.ColorMessage("blue", "Masukan nomor menu : "))
		fmt.Scan(&input)

		switch input {
		case 1:
			utils.ClearScreen()
			deleteAccountMenu()
		case 2:
			utils.ClearScreen()
			blokirAccountMenu()
		case 3:
			utils.ClearScreen()
			checkBlockMenu()
		case 4:
			utils.ClearScreen()
			restoreMenu()
		case 0:
			utils.ClearScreen()
			utils.SuccesMessage("Keluar dari program dev mode... \n")
			return
		default:
			utils.ClearScreen()
			err := errors.New("pilhan tidak valid")
			utils.ErrorMessage(err)
		}
	}
}
