package module

import (
	"errors"
	"fmt"
	"main/models"
	"main/utils"
	"reflect"
	"strconv"
)

type AccountManager interface {
	CreateAccount(name, pin string) *models.Account
	GetAccount(id string) (*models.Account, error)
	DeleteAccount(id string) error
	CheckBalance(id string) (float64, error)
	Deposit(id string, amount float64) error
	Withdraw(id string, amount float64) error
	AddToTempAccounts(account models.Account)
	GetTempAccounts() []models.Account
	RemoveFromTempAccounts(index int) error
}

type BankAccountManager struct {
	accounts     map[string]*models.Account
	TempAccounts []models.Account
}

func NewBankAccountManager() *BankAccountManager {
	return &BankAccountManager{
		accounts:     make(map[string]*models.Account),
		TempAccounts: []models.Account{},
	}
}

func (b *BankAccountManager) GetAccount(id string) (*models.Account, error) {
	account, exists := b.accounts[id]
	if !exists {
		err := errors.New(fmt.Sprintf("akun dengan id %s tidak ditemukan", id))
		return nil, err
	}
	return account, nil
}

// check pin akun terdaftar
func (b *BankAccountManager) CheckPin(pin string, accNasabah *models.Account) bool {
	for _, acc := range b.accounts {
		if acc.GetID() == accNasabah.GetID() {
			if acc.GetPin() == pin {
				return true
			}
		}
	}
	err := errors.New("kode pin salah")
	utils.ErrorMessage(err)
	return false
}

// check len of BankAccountManager
func (b *BankAccountManager) CheckLen() int {
	return len(b.accounts)
}

func (b *BankAccountManager) CheckBalance(id string) (float64, error) {
	account, err := b.GetAccount(id)
	if err != nil {
		return 0, err
	}
	return account.GetBalance(), nil
}

func (b *BankAccountManager) PrintDataAccount(mode string) *models.Account {
	if b.CheckLen() == 0 {
		err := errors.New("database kosong")
		utils.ErrorMessage(err)
		return nil
	}

	var input string
	var tempId []string
	for {
		fmt.Println(utils.ColorMessage("yellow", "+++ ====== Daftar Akun Terdaftar ====== +++"))
		index := 1
		for _, acc := range b.accounts {
			fmt.Printf("%d. %s\n", index, acc.GetName())
			tempId = append(tempId, acc.GetID())
			index++
		}
		fmt.Println("0. Kembali")
		fmt.Print(utils.ColorMessage("yellow", "Masukan index : "))
		fmt.Scan(&input)
		utils.ClearScreen()

		if !b.IsLenVar(input, 1) {
			err := errors.New(fmt.Sprintf("input harus 1 angka atau input tidak boleh melebihi %d", len(tempId)))
			utils.ErrorMessage(err)
			return b.PrintDataAccount(mode)
		}

		int_input, err := strconv.Atoi(input)
		if err != nil {
			err := errors.New("input harus berupa angka")
			utils.ErrorMessage(err)
			return b.PrintDataAccount(mode)
		}
		if int_input == 0 {
			return nil
		}

		if int_input < 0 || int_input > len(tempId) {
			err := errors.New("input tidak valid, coba lagi")
			utils.ErrorMessage(err)
			return b.PrintDataAccount(mode)
		}

		if int_input == 0 {
			return nil
		}

		acc, err := b.GetAccount(tempId[int_input-1])
		if acc == nil && err != nil {
			err := errors.New("akun tidak valid, coba lagi")
			utils.ErrorMessage(err)
			return b.PrintDataAccount(mode)
		}

		inputPin := 3
		var pin string
		if mode != "dev" {
			for {
				if inputPin == 0 {
					b.AddToTempAccounts(*acc)
					b.DeleteAccount(acc.GetID())
					return nil
				}

				fmt.Print("Masukan pin anda : ")
				fmt.Scan(&pin)
				utils.ClearScreen()
				isValidPin := b.CheckPin(pin, acc)
				if isValidPin {
					return acc
				} else {
					inputPin--
					fmt.Printf(utils.ColorMessage("red", "pin salah kesempatan sisa ")+"%d\n\n", inputPin)
				}
			}
		} else {
			return acc
		}
	}
}

func (b *BankAccountManager) PrintBlockAccount() []string {
	var tempIdBlockAcc []string
	if len(b.TempAccounts) == 0 {
		err := errors.New("tidak ada akun yang diblokir")
		utils.ErrorMessage(err)
		return nil
	}

	fmt.Println(utils.ColorMessage("yellow", "+++ ====== Daftar Akun yang Diblokir ====== +++"))
	for i, acc := range b.TempAccounts {
		tempIdBlockAcc = append(tempIdBlockAcc, acc.GetID())
		fmt.Printf("%d. ID : %s, Name : %s, Pin : %s, Balance : %.2f\n", i+1, acc.GetID(), acc.GetName(), acc.GetPin(), acc.GetBalance())
	}

	return tempIdBlockAcc
}

func (b *BankAccountManager) RestoreBlockAccount() bool {
	var inputBlockAcc string
	var intInput int
	tempIdBlockAcc := b.PrintBlockAccount()
	for {
		if tempIdBlockAcc == nil {
			return false
		}
		fmt.Println("0. Kembali")
		fmt.Print(utils.ColorMessage("yellow", "Masukan index : "))
		fmt.Scan(&inputBlockAcc)
		utils.ClearScreen()

		if !b.IsLenVar(inputBlockAcc, 1) {
			err := errors.New(fmt.Sprintf("input harus 1 angka atau input tidak boleh melebihi %d", len(tempIdBlockAcc)))
			utils.ErrorMessage(err)
			return b.RestoreBlockAccount()
		}

		int_input, err := strconv.Atoi(inputBlockAcc)
		intInput = int_input
		if err != nil {
			err := errors.New("input harus berupa angka")
			utils.ErrorMessage(err)
			return b.RestoreBlockAccount()
		}

		if int_input == 0 {
			return false
		}

		if int_input < 0 || int_input > len(tempIdBlockAcc) {
			err := errors.New("input tidak valid, coba lagi")
			utils.ErrorMessage(err)
			return b.RestoreBlockAccount()
		} else {
			break
		}

	}
	for i, acc := range b.TempAccounts {
		if acc.GetID() == tempIdBlockAcc[intInput-1] {
			b.accounts[acc.GetID()] = &acc
			err := b.RemoveFromTempAccounts(i)
			if err != nil {
				utils.ErrorMessage(err)
				return false
			} else {
				messageSucces := fmt.Sprintf("Akun %s berhasil dipulihkan\n", tempIdBlockAcc[intInput-1])
				utils.SuccesMessage(messageSucces)
				return true
			}
		}
	}

	err := errors.New("akun tidak ditemukan di daftar blokir")
	utils.ErrorMessage(err)
	return false
}

func (b *BankAccountManager) AddToTempAccounts(account models.Account) {
	b.TempAccounts = append(b.TempAccounts, account)
}

func (b *BankAccountManager) RemoveFromTempAccounts(index int) error {
	if index < 0 || index >= len(b.TempAccounts) {
		return errors.New("index out of range")
	}

	b.TempAccounts = append(b.TempAccounts[:index], b.TempAccounts[index+1:]...)

	return nil
}

func (b *BankAccountManager) GetTempAccounts() []models.Account {
	return b.TempAccounts
}

func (b *BankAccountManager) IsLenVar(input any, len_ int) bool {
	lenVal := reflect.ValueOf(input)

	switch lenVal.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
		if lenVal.Len() == len_ {
			return true
		}
	default:
		defaultMsg := errors.New(fmt.Sprintf("tipe data %s tidak mendukung operasi Len()\n", lenVal.Kind()))
		utils.ErrorMessage(defaultMsg)
	}

	return false
}
