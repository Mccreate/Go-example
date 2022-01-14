package main

import (
	"fmt"
	"github.com/mccreate/Go-exercise/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "myFirst"}
	key, value := "hello", "greeting"

	err := dictionary.Add(key, value)
	if err != nil {
		fmt.Println(err)
	}
	def, err := dictionary.Search(key)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(def)
	}
}
