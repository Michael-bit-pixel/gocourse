package basics

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
	"strings"
	// "golang.org/x/text/cases"
)

var (
	sco      int
	data_all Response
	dif      string
)

type Response struct {
	Results []Database
}

type Database struct {
	Category       string
	Difficulty     string
	Question       string
	Correct_Answer string
}

func init() {
	fmt.Println("Welcome To Our Quiz Game")
	fmt.Println("Please answer with right and get better score")
}

func main() {
	for {
		for {
			fmt.Println("Chooose your difficulty:(easy,medium,hard)")
			fmt.Scanln(&dif)
			if dif == "easy" || dif == "medium" || dif == "hard" {
				break
			} else {
				fmt.Println("Please typing easy, medium,hard with no capslock or whatsover")
				continue
			}
		}

		test := "https://opentdb.com/api.php?amount=10&type=multiple&difficulty=" + dif
		data, err := http.Get(test)
		if err != nil {
			fmt.Println("Sorry We run into API problem we will shutdown this temporarily.Once again sorry for your incovenience")
			os.Exit(1)
		}
		read, erread := io.ReadAll(data.Body)
		if erread != nil {
			fmt.Println("Sorry we got data problem.It's our mandatory to shutdown this to prevent crash.Once again sorry for your incovenience")
			os.Exit(1)
		}
		json.Unmarshal(read, &data_all)
		for i := 0; i < len(data_all.Results); i++ {
			fmt.Printf("Category %s\n", html.UnescapeString(data_all.Results[i].Category))
			fmt.Printf("Difficulty: %s\n", data_all.Results[i].Difficulty)
			fmt.Printf("The Question is: %s\n", html.UnescapeString(data_all.Results[i].Question))
			checking(html.UnescapeString(data_all.Results[i].Correct_Answer))

		}
		decision()
	}

}

func checking(A string) {
	// var ans, sca string
	// fmt.Scan(&ans, &sca)
	// fmt.Println(ans, sca)
	var ans string
	fmt.Scanln(&ans)
	if strings.EqualFold(strings.ReplaceAll(strings.ToLower(A), " ", ""), strings.ToLower(ans)) == true {
		defer fmt.Println("You are right score adds")
		sco++
	} else {
		fmt.Println("You are wrong please learn again next time")
		fmt.Printf("The correct answer is: %s\n", A)

	}
}

func decision() {
	var dec string
	for {
		fmt.Println("Do you want to continue (y or n):")
		fmt.Scanln(&dec)
		if dec == "y" || dec == "n" {
			break
		} else {
			fmt.Println("Please type y or n if you want to play")
			continue

		}
	}
	if dec == "y" {
		fmt.Printf("You score is %d\n", sco)
		clear(data_all.Results)

	} else {
		fmt.Printf("You score is %d\n", sco)
		fmt.Println("Thank You for playing with us enjoy your day and don't forget to practice again")
		os.Exit(1)
	}
}
