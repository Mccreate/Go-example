package main

import (
	"fmt"
	"github.com/mccreate/Go-exercise/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "myFirst"}
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
