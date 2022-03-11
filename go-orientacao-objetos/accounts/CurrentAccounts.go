package accounts

import "github.com/formacao-go/go-orientacao-objetos/clients"

type CurrentAccount struct {
	Holder                      clients.Holder
	AgencyNumber, AccountNumber int
	balance                     float64
}

func (c *CurrentAccount) Withraw(withdrawValue float64) string {
	canWithdraw := withdrawValue > 0 && withdrawValue <= c.balance

	if canWithdraw {
		c.balance -= withdrawValue

		return "Saque realizado com sucesso"
	}

	return "Saldo insulficiente"
}

func (c *CurrentAccount) Deposit(amount float64) (string, float64) {
	if amount > 0 {
		c.balance += amount
		return "Depósito realizado com sucesso", c.balance
	}

	return "Valor do Depósito menor que Zero.", c.balance
}

func (c *CurrentAccount) Transfer(amountForTransfer float64, accountDestiny *CurrentAccount) bool {
	if amountForTransfer < c.balance && amountForTransfer > 0 {
		c.balance -= amountForTransfer

		accountDestiny.Deposit(amountForTransfer)

		return true
	}

	return false
}

func (c *CurrentAccount) GetBalance() float64 {
	return c.balance
}
