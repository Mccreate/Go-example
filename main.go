package main

import (
	"fmt"
	"github.com/mccreate/Go-exercise/accounts"
)

func main() {
	account := accounts.NewAccount("mccreate")
	account.Deposit(100)
	fmt.Println(account.GetBalance())
	err := account.Withdraw(1000)

	// if error is occurred,
	if err != nil {
		// sys.exit
		// log.Fatalln(err)
		fmt.Println(err)
	}
	fmt.Println(account.String())

}
