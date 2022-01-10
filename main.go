package main

import (
	"fmt"
	"strings"
)

// Go needs Data Type in below.
/*
   func funcName(argv1 argv1's Datatype, argv2 argv2's Datatype, ...) return's DataType {
	   return blah
   }
*/
func multiply(a int, b int) int {
	return a * b
}

// Go can return many argvs.
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// Nake Return Example.
func lenAndUpperNaked(name string) (length int, uppercase string) {
	fmt.Println("Naked Return Example.")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

// defer Example.
// After Func returns, "defer" keyowrd executes.
func lenAndUpperDefer(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done.")
	fmt.Println("Defer Example.")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

// When there are many func arguments, You can shorten using "..." keyword.
func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	fmt.Println(multiply(2, 2))
	totalLength, upperName := lenAndUpper("mccreate")
	// Go must be use all variable.
	fmt.Println(totalLength, upperName)

	// You can ignore "Variable declared but not used" in below.
	exLength, _ := lenAndUpper("mccreate")
	fmt.Println(exLength)

	repeatMe("mc", "cre", "ate")

	fmt.Println(lenAndUpperNaked("mccreate"))

	fmt.Println(lenAndUpperDefer("mccreate"))
}
