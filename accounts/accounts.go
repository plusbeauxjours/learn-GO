package accounts

import "fmt"

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	fmt.Println("owner", owner)
	account := Account{owner: owner, balance: 0}
	fmt.Println("account", account)
	fmt.Println("&account", &account)
	return &account	// struct의 address는 이렇게 생겼다. &{mj 0}
}