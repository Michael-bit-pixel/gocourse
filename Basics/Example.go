package basics

import (
	"fmt"
	"math/rand"
	"time"
)

var sco int
var QnA string

func init() {
	fmt.Println("Guess a calculation")
	fmt.Println("Choose wisely")
}

func main() {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	var ans int
	for {
		target := random.Intn(4) + 1
		switch target {
		case 1:
			Q1, Q2 := randoming(1)
			fmt.Printf("What is the answer of %d + %d:\n", Q1, Q2)
			fmt.Scanln(&ans)
			score(check(add(Q1, Q2), ans))
		case 2:
			Q1, Q2 := randoming(2)
			fmt.Printf("What is the answer of %d - %d:\n", Q1, Q2)
			fmt.Scanln(&ans)
			score(check(substract(Q1, Q2), ans))
		case 3:
			Q1, Q2 := randoming(3)
			fmt.Printf("What is the answer of %d / %d:\n", Q1, Q2)
			fmt.Scanln(&ans)
			score(check(divide(Q1, Q2), ans))
		default:
			Q1, Q2 := randoming(4)
			fmt.Printf("What is the answer of %d x %d:\n", Q1, Q2)
			fmt.Scanln(&ans)
			score(check(multiply(Q1, Q2), ans))
		}
		fmt.Printf("Your score is: %d\n", sco)
		for {
			fmt.Println("Do you want to continue (y or n):")
			fmt.Scanln(&QnA)
			if QnA == "y" || QnA == "n" {
				break
			} else {
				continue
			}
		}
		if QnA == "y" {
			continue
		} else {
			fmt.Println("Thank you for playing with us")
			break
		}
	}
}

func randoming(num int) (int, int) {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	Q1 := random.Intn(100) + 1
	Q2 := random.Intn(100) + 1
	if num != 3 {
		for {
			if Q1 > Q2 {
				break
			} else {
				Q2 = random.Intn(100) + 1
				continue
			}
		}
	} else {
		for {
			if Q1%Q2 == 0 {
				break
			} else {
				Q2 = random.Intn(100) + 1
				continue
			}
		}
	}
	return Q1, Q2
}

func add(f int, s int) int {
	return f + s
}

func substract(f int, s int) int {
	return f - s
}

func divide(f int, s int) int {
	return f / s
}

func multiply(f int, s int) int {
	return f * s
}

func check(A int, B interface{}) int {
	if A == B {
		return 0
	} else {
		return 1
	}
}

func score(R interface{}) {
	if R == 0 {
		defer fmt.Println("You Right")
		sco = 1 + sco
	} else {
		fmt.Println("You Wrong")
	}
}
