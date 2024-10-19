package module

import (
	"errors"
	"fmt"

	"main/utils"
)

func (b *BankAccountManager) Deposit(id string, amount float64) {
	account, err := b.GetAccount(id)
	if err != nil {
		return
	}

	if amount < 1000 {
		err := errors.New("gagal deposit, minimum deposit 1000")
		utils.ErrorMessage(err)
		return
	}

	account.Deposit(amount)
	utils.SuccesMessage(fmt.Sprintf("Deposit %.2f berhasil! Saldo sekarang : %.2f\n", amount, account.Balance))
}
