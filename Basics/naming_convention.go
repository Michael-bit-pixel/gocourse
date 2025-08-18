package basics

import "fmt"

type EmployeeGoogle struct {
	FirstName string
	LastName  string
	Age       int
}

type EmployeeApple struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	//PascalCase
	// Bg, CalculateArea , UserInfo , NewHTTPRequest
	// Strucks , Interfaces , enums

	//snake_case
	// Bg , user_id , first_name , http_request

	//UPPERCASE
	//Use case is Constants

	//mixedCase
	// Bg, javaScript , htmlDocument , isValid

	const MAXRETRESS = 5

	var employeeID = 1001
	fmt.Println("EmployeeID", employeeID)
}
