package basics

import (
	"fmt"
	"math"
)

func main() {
	// Variable declaration
	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println("Sum of a and b is ", result)

	result = a * b
	fmt.Println("Multiplication", result)

	result = a - b
	fmt.Println("Substraction", result)

	result = a / b
	fmt.Println("Division", result)

	result = a % b
	fmt.Println("Remainder", result)

	const p float64 = 22.0 / 7.0
	fmt.Println(p)

	// Overflow with signed integers
	var maxInt int64 = 9223372036854775807 // max value that int64 can hold
	fmt.Println(maxInt)

	maxInt = maxInt + 1
	fmt.Println(maxInt)

	// Overflow with unsigned integers
	var uMaxInt uint64 = 18446744073709551615 // max value for uint64 type
	fmt.Println(uMaxInt)

	uMaxInt = uMaxInt + 1
	fmt.Println(uMaxInt)

	// Underflow with floating points number
	var smallFloat float64 = 1.0e-323

	fmt.Println(smallFloat)

	smallFloat = smallFloat / math.MaxFloat64

	fmt.Println(smallFloat)
}
