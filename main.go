package main

import "fmt"

// Pointer

func main() {
	a := 2
	// linking b addr to a addr
	b := &a
	a = 10
	// you can see memory addr with & keyword
	fmt.Println(&a, &b)

	// you can see value in memory addr * keyword
	fmt.Println(a, *b)

	// you can change a val with *b (if you link a with b)
	*b = 20
	fmt.Println(a, *b)
}
