package main

import (
	"fmt"
	"github.com/mccreate/Go-exercise/accounts"
	"log"
)

func main() {
	account := accounts.NewAccount("mccreate")
	account.Deposit(100)
	fmt.Println(account.GetBalance())
	err := account.Withdraw(1000)

	// if error is occurred,
	if err != nil {
		// sys.exit
		log.Fatalln(err)
	}
	fmt.Println(account.GetBalance())

}
