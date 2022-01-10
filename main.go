package main

import "fmt"

// GO has only "for"!
// for loop Example

/*
	Example 1
	for loopVar := range loopThing {
		loop code...
	}
*/
func loopExample1(numbers ...int) int {
	sum := 0
	// number == loopVar
	// numbers == loopThing
	for index, number := range numbers {
		fmt.Println(index, number)
		sum += number
	}
	return sum
}

/*
	Example 2
	It's similar C for loop style.
	for(int i = 0; i<loopEnd; i++){
		loop code...
	}

	In Go,
	for i := 0; i < loopEnd; i++ {
		loop code...
	}
*/
func loopExample2(numbers ...int) int {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
	}
	return 1
}

// "_" symbol is ignore symbol.
func loopExample3(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// Naked Return
func loopExample4(numbers ...int) (total int) {
	total = 0
	for _, number := range numbers {
		total += number
	}
	return
}

// defer
func loopExample5(numbers ...int) (total int) {

	total = 0
	for _, number := range numbers {
		total += number
	}
	defer fmt.Println("Defer... I'm done.")
	defer fmt.Println(total)
	return

}

func main() {
	fmt.Println(loopExample1(1))
	fmt.Println(loopExample2(1, 2))
	fmt.Println(loopExample3(1, 2, 3))
	fmt.Println(loopExample4(1, 2, 3, 4))
	fmt.Println(loopExample5(1, 2, 3, 4, 5))
}
