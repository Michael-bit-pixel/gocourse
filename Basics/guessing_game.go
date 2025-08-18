package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// Generate a random number between 1 and 100
	target := random.Intn(100) + 1

	// Welcome message
	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("I have chosen a number between 1 and 100")
	fmt.Println("Try to guess the number I have chosen")

	var guess int
	for {
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&guess)

		//Check if the guess if correct
		if guess == target {
			fmt.Println("Congratulations, you have guessed the number correctly!")
			break
		} else if guess < target {
			fmt.Println("Too low! Try guessing high number.")
		} else {
			fmt.Println("Too high! Try guessing a lower number.")
		}
	}

}
