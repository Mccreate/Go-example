package main

import "fmt"

//Go needs Data Type in below.
/*
   func funcName(argv1 argv1's Datatype, argv2 argv2's Datatype, ...) return's DataType {
	   return blah
   }
*/
func multiply(a int, b int) int {
	return a * b
}

func main() {
	fmt.Println(multiply(2, 2))
}
