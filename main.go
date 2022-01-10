package main

import (
	"fmt"
	"strings"
)

//Go needs Data Type in below.
/*
   func funcName(argv1 argv1's Datatype, argv2 argv2's Datatype, ...) return's DataType {
	   return blah
   }
*/
func multiply(a int, b int) int {
	return a * b
}

//Go can return many argvs.
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	fmt.Println(multiply(2, 2))
	totalLength, upperName := lenAndUpper("mccreate")
	//Go must be use all variable.
	fmt.Println(totalLength, upperName)

	//You can ignore Variable declared but not used in below.
	exLength, _ := lenAndUpper("mccreate")
	fmt.Println(exLength)
}
