package basics

import "fmt"

var middlename = "Cane"

func main() {
	// var age int
	var name string = "John"
	var name1 = "Jane"

	count := 10
	lastName := "Smith"
	middlename := "Mayor"

	fmt.Println(middlename)
	fmt.Println(lastName)
	fmt.Println(count)
	fmt.Println(name1)
	fmt.Println(name)
	// default values
	// numeric Types: 6
	// Boolean Types: false
	// String Types: ""
	// Pointers, slices, maps,function,and structure:nil

	// ---- Scope
	// fmt.Println(firstname)
}

func printName() {
	firstName := "Michael"
	fmt.Println(firstName)
}
