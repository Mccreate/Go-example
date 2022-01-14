package main

import (
	"fmt"
	"github.com/mccreate/Go-exercise/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	key, value := "hello", "greeting"

	err := dictionary.Add(key, value)
	if err != nil {
		fmt.Println(err)
	}
	def1, _ := dictionary.Search(key)
	fmt.Println(def1)

	err2 := dictionary.Update(key, "first")
	if err2 != nil {
		fmt.Println(err2)
	}

	dictionary.Delete(key)

	def2, errMsg := dictionary.Search(key)
	fmt.Println(def2, errMsg)
}
