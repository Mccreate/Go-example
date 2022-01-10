package main

import "fmt"

func condition1(age int) bool {
	if age < 18 {
		return false
	}
	return true
}

func condition2(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}
func main() {
	fmt.Println(condition1(17))
	fmt.Println(condition2(17))
}
