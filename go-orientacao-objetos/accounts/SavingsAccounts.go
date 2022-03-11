package accounts

import "github.com/formacao-go/go-orientacao-objetos/clients"

type SavingsAccount struct {
	Holder                                 clients.Holder
	AgencyNumber, AccountNumber, Operation int
	balance                                float64
}

func (c *SavingsAccount) Withraw(withdrawValue float64) string {
	canWithdraw := withdrawValue > 0 && withdrawValue <= c.balance

	if canWithdraw {
		c.balance -= withdrawValue

		return "Saque realizado com sucesso"
	}

	return "Saldo insulficiente"
}

func (c *SavingsAccount) Deposit(amount float64) (string, float64) {
	if amount > 0 {
		c.balance += amount
		return "Depósito realizado com sucesso", c.balance
	}

	return "Valor do Depósito menor que Zero.", c.balance
}

func (c *SavingsAccount) GetBalance() float64 {
	return c.balance
}
