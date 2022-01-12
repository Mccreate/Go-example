package main

import "fmt"

// Struct

// If you want right this with Maps,
/*
	{
		"name" : "Min Woo Kim",
		"age" : 26,
		"school" : "UOU"
	}
*/

// struct is needed.
type person struct {
	name   string
	age    int
	school string
	hobby  []string
}

func main() {
	mccreate := person{"Min Woo Kim", 26, "UOU", []string{"Playing a guitar", "Playing a Game with XBOX"}}
	fmt.Println(mccreate)

	// You also call a element.
	fmt.Println(mccreate.age)
	fmt.Println(mccreate.name)
	fmt.Println(mccreate.hobby)
	fmt.Println(mccreate.hobby[0])
}
