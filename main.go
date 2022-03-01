package main

import (
	"fmt"

	"github.com/plusbeauxjours/learnGo/accounts"
)


func main() {
	account := accounts.NewAccount("mj")
	fmt.Println(account)
	account.Deposit(100)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
}