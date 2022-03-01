package main

import (
	"fmt"

	"github.com/plusbeauxjours/learnGo/accounts"
)


func main() {
	account := accounts.NewAccount("mj")
	fmt.Println((account))
}