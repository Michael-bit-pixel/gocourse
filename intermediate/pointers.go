package basics

import "fmt"

func main() {

	var ptr *int

	var a int = 10
	ptr = &a //referencing

	fmt.Println(a)
	fmt.Println(ptr)
	// fmt.Println(*ptr) // dereferencing a pointer
	// if ptr == nil {
	// 	fmt.Println("Pointer is nill")
	// }

	modifyValue(ptr)
	fmt.Println(a)

}

func modifyValue(ptr *int) {
	*ptr++
}
