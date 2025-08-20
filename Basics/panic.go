package basics

import "fmt"

func main() {

	// panuc(interface{})

	//Example of a valid input
	process(10)

	//Example of a invalid input
	process(-10)

}

func process(input int) {

	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	if input < 0 {
		fmt.Println("Before Panic")
		panic("input must be a non-negative number")
		// fmt.PrintIn("After Panic")
		// defer fmt.Println("Deferred 3")
	}
	fmt.Println("Processing input:", input)
}
