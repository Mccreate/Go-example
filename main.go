package main

import (
	"fmt"
	"time"
)

func main() {
	// how to make channel
	// chan dType
	c := make(chan bool)
	people := [2]string{"mc", "create"}
	for _, person := range people {
		go isSexy(person, c)
	}
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	// You can assign value with channel "val <- channel_val"
	c <- true
}
