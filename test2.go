package main

import (
	"fmt"

	"github.com/plusbeauxjours/learnGo/accounts"
)


func test2() {
	account := accounts.NewAccount("mj")
	fmt.Println(account)
	account.Deposit(100)
	fmt.Println(account)
}