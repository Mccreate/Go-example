package main

import "fmt"

// Map.

// You assign map, in below.
// var_name := map[key_dType]val_dType{key1: val1, key2: val2, ...}
func main() {
	mccreate := map[string]string{"name": "Min Woo Kim", "school": "UOU"}
	fmt.Println(mccreate)

	// You can call map in below.
	// var_name[key]
	fmt.Println(mccreate["name"])

	for key, _ := range mccreate {
		fmt.Println(key)
	}
}
