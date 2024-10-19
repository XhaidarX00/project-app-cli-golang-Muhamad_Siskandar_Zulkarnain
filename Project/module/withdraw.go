package module

import (
	"errors"
	"fmt"

	"main/utils"
)

func (b *BankAccountManager) Withdraw(id string, amount float64) error {
	account, err := b.GetAccount(id)
	if err != nil {
		return err
	}
	if amount <= 0 {
		err := errors.New("gagal Credit, tidak bisa menarik 0 atau nilai negatif")
		utils.ErrorMessage(err)
		return err
	}

	success := account.Withdraw(amount)
	if success {
		utils.SuccesMessage(fmt.Sprintf("Penarikan %.2f berhasil! Saldo sekarang : %.2f\n", amount, account.Balance))
	} else {
		err := errors.New("penarikan gagal, saldo tidak mencukupi")
		utils.ErrorMessage(err)
	}
	return nil
}
