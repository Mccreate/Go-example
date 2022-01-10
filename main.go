package main

import "fmt"

func loop1(loops ...int) {
	for _, loop := range loops {
		fmt.Println(loop)
	}
}

func loop2(loops ...int) {
	for i := 0; i < len(loops); i++ {
		fmt.Println(loops[i])
	}
}

func main() {
	loop1(1, 2, 3, 4)
	loop2(5, 6, 7, 8)
}
