package module

import (
	"fmt"
	"main/models"
	"main/utils"
)

func CheckBalance(account *models.Account) {
	CBMsg := fmt.Sprintf("Saldo akun %s (ID: %s) : %.2f\n", account.Name, account.ID, account.Balance)
	utils.SuccesMessage(CBMsg)
}
