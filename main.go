package main

import "fmt"

// list

func main() {
	// how to make list?
	// list_name := [list_size]Data_Type{"your", "list", "element"}
	names := [3]string{"mc", "cre", "ate"}
	names2 := [3]string{"mc"}
	numbers := [3]int{1, 2}

	// OverFlow!!
	// names[3] = "Overflow String"
	fmt.Println(names)
	fmt.Println(names2)
	fmt.Println(numbers)

	// You can assign unlimited length list
	// list_name := []Data_type{list_elements}
	non_limited_list := []int{1, 2, 3, 4, 5}
	fmt.Println(non_limited_list)

	// append Example
	non_limited_list = append(non_limited_list, 4)
	fmt.Println(non_limited_list)
}
