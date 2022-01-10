package main

import "fmt"

// switch

func canIDrink1(age int) bool {
	switch age {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

// you can use "switch" right this.
func canIDrink2(age int) bool {
	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}
	return false
}

func main() {
	fmt.Println(canIDrink1(16))
	fmt.Println(canIDrink2(16))
}
