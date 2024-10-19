package module

import (
	"errors"
	"fmt"
	"main/models"
	"main/utils"
)

func (b *BankAccountManager) CreateAccount(name, pin string) *models.Account {
	newAccount := models.NewAccount(name, pin)
	b.accounts[newAccount.ID] = newAccount
	succes := fmt.Sprintf("Akun berhasil dibuat: ID: %s, Name: %s, Balance: %.2f\n", newAccount.ID, newAccount.Name, newAccount.Balance)
	utils.SuccesMessage(succes)
	return newAccount
}

func (b *BankAccountManager) DeleteAccount(id string) {
	if _, exists := b.accounts[id]; !exists {
		err := errors.New(fmt.Sprintf("akun dengan ID %s tidak ditemukan", id))
		utils.ErrorMessage(err)
	}

	delete(b.accounts, id)
	succes := fmt.Sprintln("Akun berhasil diblokir atau dihapus.\nHubungi dev untuk memulihkan akun.")
	utils.SuccesMessage(succes)
}

func (b *BankAccountManager) BlockAccount(id string) {
	acc, err := b.GetAccount(id)
	if err != nil {
		return
	}
	b.AddToTempAccounts(*acc)
	b.DeleteAccount(id)
}
