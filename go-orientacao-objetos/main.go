package main

import (
	"fmt"

	"github.com/formacao-go/go-orientacao-objetos/accounts"
	"github.com/formacao-go/go-orientacao-objetos/clients"
)

func PayTicket(account verifyAccount, ticketAmount float64) {
	account.Withraw(ticketAmount)
}

type verifyAccount interface {
	Withraw(amunt float64) string
}

func main() {
	clientDavi := clients.Holder{
		Name:       "Davi",
		CPF:        "001.002.003-44",
		Occupation: "Software Developer",
	}

	daviAccount := accounts.CurrentAccount{Holder: clientDavi}
	daviAccount.Deposit(200)

	fmt.Println(daviAccount.GetBalance())

	PayTicket(&daviAccount, 50)

	fmt.Println(daviAccount.GetBalance())
}
