package basics

import "fmt"

// func swap[T any](a, b T) (T, T) {
// 	return b, a
// }

type Stack[T any] struct {
	element []T
}

func (s *Stack[T]) push(element T) {
	s.element = append(s.element, element)
}

func (s *Stack[T]) pop() (T, bool) {
	if len(s.element) == 0 {
		var zero T
		return zero, false
	}
	element := s.element[len(s.element)-1]
	s.element = s.element[:len(s.element)-1]
	return element, true
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.element) == 0
}

func (s *Stack[T]) printAll() {
	if len(s.element) == 0 {
		fmt.Println("The stack is empty.")
		return
	}
	fmt.Print("Stack elements: ")
	for _, element := range s.element {
		fmt.Print(element)
	}
	fmt.Println()
}

func main() {

	// x, y := 1, 2
	// x, y = swap(x, y)
	// fmt.Println(x, y)

	// x1, y1 := "John", "Jane"
	// x1, y1 = swap(x1, y1)
	// fmt.Println(x1, y1)

	intStack := Stack[int]{}
	intStack.push(1)
	intStack.push(2)
	intStack.push(3)
	intStack.printAll()
	fmt.Println(intStack.pop())
	intStack.printAll()
	fmt.Println(intStack.pop())
	fmt.Println("Is stack empty:", intStack.isEmpty())
	fmt.Println(intStack.pop())
	fmt.Println("Is stack empty:", intStack.isEmpty())

	stringStack := Stack[string]{}
	stringStack.push("Hello")
	stringStack.push("World")
	stringStack.push("John")
	stringStack.printAll()
	fmt.Println(stringStack.pop())
	fmt.Println("Is stringStack is empty:", stringStack.isEmpty())
	stringStack.printAll()
	fmt.Println(stringStack.pop())
	fmt.Println(stringStack.pop())
	fmt.Println("Is stringStack is empty:", stringStack.isEmpty())
}
