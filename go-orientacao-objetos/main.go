package main

import (
	"fmt"
	"go-orientacao-objetos/accounts"
)

func main() {
	daviAccount := accounts.CurrentAccount{holder: "Davi", balance: 500}
	jhoyAccount := accounts.CurrentAccount{holder: "Joelma", balance: 200}

	fmt.Println(daviAccount)
	fmt.Println(jhoyAccount)

	status := daviAccount.Transfer(150, &jhoyAccount)

	fmt.Println(status)
	fmt.Println(daviAccount)
	fmt.Println(jhoyAccount)
}
