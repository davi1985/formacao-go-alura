```go

package main

import "fmt"

type CurrentAccount struct {
	holder        string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func main() {
	daviAccount := CurrentAccount{
		holder:        "Davi",
		agencyNumber:  3232,
		accountNumber: 1212,
		balance:       250.25,
	}

	fmt.Println(daviAccount)

	jhoyAccount := CurrentAccount{"Joelma", 3232, 1212, 250}

	fmt.Println(jhoyAccount)

	var sarahAccount *CurrentAccount
	sarahAccount = new(CurrentAccount)
	sarahAccount.holder = "Sarah"
	sarahAccount.balance = 1000.50

	fmt.Println(sarahAccount)
}

// retorno mútiplos
package main

import "fmt"

type CurrentAccount struct {
	holder        string
	agencyNumber  int
	accountNumber int
	balance       float64
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

func main() {
	daviAccount := CurrentAccount{}
	daviAccount.holder = "Davi"
	daviAccount.balance = 500

	// fmt.Println(daviAccount.balance)

	// fmt.Println(daviAccount.Withraw(-100))
	// fmt.Println(daviAccount.balance)
	status, amount := daviAccount.Deposit(300)

	fmt.Println(status, amount)
	// fmt.Println(daviAccount.Deposit(200))
	// fmt.Println(daviAccount.balance)

}

// composition
package main

import (
	"fmt"

	"github.com/formacao-go/go-orientacao-objetos/accounts"
	"github.com/formacao-go/go-orientacao-objetos/clients"
)

func main() {
	clientDavi := clients.Holder{
		Name:       "Davi",
		CPF:        "001.002.003-44",
		Occupation: "Software Developer",
	}

	daviAccount := accounts.CurrentAccount{
		Holder:        clientDavi,
		AgencyNumber:  3879,
		AccountNumber: 1212,
	}

	clientJhoy := clients.Holder{
		Name:       "Joelma",
		CPF:        "001.002.003-44",
		Occupation: "Manager",
	}

	jhoyAccount := accounts.CurrentAccount{
		Holder:        clientJhoy,
		AgencyNumber:  3879,
		AccountNumber: 1414,
	}

	fmt.Println(daviAccount)
	fmt.Println(jhoyAccount)

	// status := daviAccount.Transfer(150, &jhoyAccount)

	fmt.Println(daviAccount.GetBalance())
	fmt.Println(jhoyAccount.GetBalance())
}



```
