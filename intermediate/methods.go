package basics

import "fmt"

type Shape struct {
	Rectangle
}

type Rectangle struct {
	length float64
	width  float64
}

// Method with value reciever
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Method with pointer reciever
func (r *Rectangle) Scale(factor float64) {
	r.length *= factor // r.length = r.length * factor
	r.width *= factor
}

func main() {

	rec := Rectangle{length: 10, width: 10}
	area := rec.Area()
	fmt.Println("Area of rectangle with width 9 and length 10 is ", area)
	rec.Scale(2)
	area = rec.Area()
	fmt.Println("Area of rectangle with a factor of 2 ", area)

	num := MyInt(-5)
	numl := MyInt(9)
	fmt.Println(num.IsPositive())
	fmt.Println(numl.IsPositive())
	fmt.Println(num.welcomeMessage())

	s := Shape{Rectangle: Rectangle{length: 10, width: 9}}
	fmt.Println(s.Area())
	fmt.Println(s.Rectangle.Area())
}

type MyInt int

//Method on a user-defined type
func (m MyInt) IsPositive() bool {
	return m > 0
}

func (MyInt) welcomeMessage() string {
	return "Welcome to MyInt Type"
}
