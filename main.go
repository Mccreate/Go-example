package main

import (
	"fmt"
	"time"
)

func main() {
	// how to make channel
	// chan dType
	c := make(chan string)
	people := [5]string{"mc", "create", "Kim", "Min", "Woo"}
	for _, person := range people {
		go isSexy(person, c)
	}
	// channel msg receiver
	for i := 0; i < len(people); i++ {
		fmt.Print("Waiting for", i)
		fmt.Println(<-c)
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 2)
	// You can assign value with channel "val <- channel_val"
	c <- person + " is sexy"
}
