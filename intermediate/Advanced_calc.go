package intermediate

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

var (
	mode     string
	account  string
	opening  strings.Builder
	command  string
	tagtrial string
	pass     string
)

func init() {
	tm, err := os.ReadFile("brute.txt")
	if err != nil {
		return
	}
	mat, _ := strconv.Atoi(string(tm))
	if mat >= time.Now().Minute() {
		log.Fatalln("Sorry you have been excused for brute attempt try again for 10 minutes")
	} else {
		os.Remove("brute.txt")
		return
	}
}

func init() {
	acc := flag.NewFlagSet("ID", flag.ExitOnError)
	rmv := flag.NewFlagSet("RMV", flag.ExitOnError)
	act := acc.String("account", "Midok@example", "For input a account")
	past := acc.String("password", "12345", "For input a password")
	armv := rmv.String("account", "Midok@example", "For input a account")
	prmv := rmv.String("password", "12345", "For input a password")
	if len(os.Args) < 2 {
		fmt.Println("We assume you are having an account if you want go to setup account please end and in the beggining (-ID -account=Midok@example -password=12345)")
		return
	}
	switch os.Args[1] {
	case "-ID":
		acc.Parse(os.Args[2:])
		err := godotenv.Load(fmt.Sprintf(".env.%s", *act))
		if err != nil {
			fmt.Println("That account is not used")
			break
		}
		log.Fatalln("Sorry but that account is already exist use another")
	case "-RMV":
		rmv.Parse(os.Args[2:])
		err := godotenv.Load(fmt.Sprintf(".env.%s", *armv))
		if err != nil {
			log.Fatalln("Upsss something went wrong:", err)
		}
		val, content := os.LookupEnv(strings.ReplaceAll(base64.StdEncoding.EncodeToString([]byte(strings.ToUpper(*armv))), "=", ""))
		if !content {
			fmt.Println("Sorry but that account does not exist")
			return
		}
		pases := sha256.Sum256([]byte(*prmv))
		if val == strings.ReplaceAll(strings.ToUpper(base64.StdEncoding.EncodeToString(pases[:])), "=", "") {
			os.Remove(".env." + *armv)
			os.Remove(*armv + "_Activity.log")
			os.Remove(*armv + "hist.log")
			os.Remove(*armv + "history.txt")
			fmt.Println("You has been sucessfully delete all of it's content")
			return
		} else {
			fmt.Println("Nice try asshole.Good luck trying to exterminated other account")
			filepathA, _ := os.OpenFile("brute.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			filepathA.WriteString(fmt.Sprintf("%v", time.Now().Minute()+10))
			closef(filepathA)
			log.Fatalln("Sorry no offense,after all this is just security issue")
		}

	default:
		log.Fatalln("Upsss something went wrong")
	}
	passe := sha256.Sum256([]byte(*past))
	envMap := map[string]string{
		strings.ReplaceAll(base64.StdEncoding.EncodeToString([]byte(strings.ToUpper(*act))), "=", ""): strings.ReplaceAll(strings.ToUpper(base64.StdEncoding.EncodeToString(passe[:])), "=", ""),
	}
	godotenv.Write(envMap, fmt.Sprintf(".env.%s", *act))
	err := godotenv.Write(envMap, fmt.Sprintf(".env.%s", *act))
	if err != nil {
		log.Fatalln("Upsss something went wrong")
	}
	fmt.Println("Your account is sucessfully setup(Note:If you make account without following those format > -ID -account=Midok@example -password=12345 <,it will be set to default which we display)")
	filepath, _ := os.OpenFile(*act+"_Activity.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	log.New(filepath, "INFO:", log.Ldate|log.Ltime|log.Llongfile).Println("Account is sucesfully created")
	closef(filepath)
}

func init() {
	for i := 1; i < 10; i++ {
		opening.WriteString("Please enter your Account name(If you don't have one please in the beggining type (-ID -account=\"Migue@exam\" -password=\"12345\"))")
		fmt.Println(opening.String())
		fmt.Scan(&account)
		godotenv.Load(fmt.Sprintf(".env.%s", account))
		val, content := os.LookupEnv(strings.ReplaceAll(base64.StdEncoding.EncodeToString([]byte(strings.ToUpper(account))), "=", ""))
		fmt.Println(os.Getenv("QURWQU5DRURNQVNURVJAR01BSUw"))
		if !content {
			fmt.Println("Sorry we couldn't found your account")
			fmt.Println("You can type y for continuing and n for surrend or you can eat shit")
			fmt.Scan(&tagtrial)
			switch tagtrial {
			case "y":
				if i == 9 {
					command = "BRUTE"
					filepathB, _ := os.OpenFile("brute.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
					filepathB.WriteString(fmt.Sprintf("%v", time.Now().Minute()+10))
					closef(filepathB)
					log.Fatalln("Sorry But Stop bruteforcing this calculator")
				}
				opening.Reset()
				continue
			case "n":
				command = "endgame"
				log.Fatalln("Yuppp try the best next time")
			default:
				command = "Notype"
				log.Fatalln("Nice Try asshole next time type with your mind opened wide")
			}
		} else {
			tagtrial = val
			fmt.Println("That account is exist")
			filepath, _ := os.OpenFile(account+"_Activity.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			log.New(filepath, "INFO:", log.Ldate|log.Ltime|log.Llongfile).Println("Account is being scanned")
			closef(filepath)
			break
		}

	}
	for i := 1; i < 5; i++ {
		opening.Reset()
		opening.WriteString("Please type your password:")
		fmt.Println(opening.String())
		fmt.Scan(&pass)
		passe := sha256.Sum256([]byte(pass))
		if strings.ReplaceAll(strings.ToUpper(base64.StdEncoding.EncodeToString(passe[:])), "=", "") == tagtrial {
			fmt.Println("Your Password is absolutely cinema")
			filepath, _ := os.OpenFile(account+"_Activity.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			log.New(filepath, "INFO:", log.Ldate|log.Ltime|log.Llongfile).Println("Account is being entered(If this is not your being entered around that time please change Your Password!!!!)")
			closef(filepath)
			opening.Reset()
			break
		} else {
			fmt.Println("Sorry Your password is not match.If you want end this please type ctrl-c")
			filepath, _ := os.OpenFile(account+"_Activity.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			log.New(filepath, "INFO:", log.Ldate|log.Ltime|log.Llongfile).Println("Wrong Password is being Types(If this is not your being typing wromg password around that time please change Your Password!!!!)")
			closef(filepath)
		}
		if i == 4 {
			filepath, _ := os.OpenFile(account+"_Activity.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			filepathB, _ := os.OpenFile("brute.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			log.New(filepath, "INFO:", log.Ldate|log.Ltime|log.Llongfile).Println("Brute Force Password(If this is not your being typing wrong password around that time please change Your Password!!!!)")
			filepathB.WriteString(fmt.Sprintf("%v", time.Now().Minute()+10))
			closef(filepath)
			closef(filepathB)
			command = "BRUTE"
			log.Fatalln("You're type wrong password for 5 times and we can't tolerance brute force")
		}
	}
}

func init() {
	opening.WriteString(fmt.Sprintf("Hi %s\n", account))
	opening.WriteString("Welcome to Advanced calculator project\n Choose mode advanced or basic:")
	fmt.Println(opening.String())
	fmt.Scan(&command)
	if command != "basic" && command != "advanced" {
		fmt.Println("We assume you are not chossing so we choose basic mode if you want go to advanced please end and in the beggining (-mode advanced)")
		mode = "basic"
	} else {
		switch command {
		case "advanced":
			mode = "advanced"
		case "basic":
			mode = "basic"
		default:
			mode = "basic"
		}
	}

}

type two struct {
	first, second int
	third, fourth float64
	order         string
}

type one struct {
	variable float64
}

func (a two) adding() {
	var (
		file, _ = os.OpenFile(account+"hist.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		ftxt, _ = os.OpenFile(account+"history.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		times   = time.Now().Format("Jan 02, 2006 03:04 PM")
	)
	defer closef(ftxt)
	defer closef(file)
	switch a.order {
	case "i1i2":
		log.New(file, "INFO:", log.Ldate|log.Ltime|log.Llongfile).Printf("%v + %v = %v", a.first, a.second, a.first+a.second)
		ftxt.WriteString(fmt.Sprintf("%s: %v + %v = %v\n", times, a.first, a.second, a.first+a.second))
		fmt.Printf("The result is %v\n", a.first+a.second)

	case "f1f2":
		log.New(file, "INFO:", log.Ldate|log.Ltime|log.Llongfile).Printf("%v + %v = %v", a.third, a.fourth, a.third+a.fourth)
		ftxt.WriteString(fmt.Sprintf("%s: %v + %v = %v\n", times, a.third, a.fourth, a.third+a.fourth))
		fmt.Printf("The result is %v\n", a.third+a.fourth)
	default:
		fmt.Println("Sorry but we cannot process float with integer because it can cause lack of calculating.")
		log.New(file, "ERROR:", log.Ldate|log.Ltime|log.Llongfile).Println("Error")
	}

}

func (a two) substract() {
	var (
		file, _ = os.OpenFile(account+"hist.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		ftxt, _ = os.OpenFile(account+"history.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		times   = time.Now().Format("Jan 02, 2006 03:04 PM")
	)
	defer closef(ftxt)
	defer closef(file)
	switch a.order {
	case "i1i2":
		log.New(file, "INFO:", log.Ldate|log.Ltime|log.Llongfile).Printf("%v - %v = %v", a.first, a.second, a.first-a.second)
		ftxt.WriteString(fmt.Sprintf("%s: %v - %v = %v\n", times, a.first, a.second, a.first-a.second))
		fmt.Printf("The result is %v\n", a.first-a.second)

	case "f1f2":
		log.New(file, "INFO:", log.Ldate|log.Ltime|log.Llongfile).Printf("%v - %v = %v", a.third, a.fourth, a.third-a.fourth)
		ftxt.WriteString(fmt.Sprintf("%s: %v - %v = %v\n", times, a.third, a.fourth, a.third-a.fourth))
		fmt.Printf("The result is %v\n", a.third-a.fourth)
	default:
		fmt.Println("Sorry but we cannot process float with integer because it can cause lack of calculating.")
		log.New(file, "ERROR:", log.Ldate|log.Ltime|log.Llongfile).Println("Error")
	}
}

func (a two) multiple() {
	var (
		file, _ = os.OpenFile(account+"hist.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		ftxt, _ = os.OpenFile(account+"history.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		times   = time.Now().Format("Jan 02, 2006 03:04 PM")
	)
	defer closef(ftxt)
	defer closef(file)
	switch a.order {
	case "i1i2":
		log.New(file, "INFO:", log.Ldate|log.Ltime|log.Llongfile).Printf("%v * %v = %v", a.first, a.second, a.first*a.second)
		ftxt.WriteString(fmt.Sprintf("%s: %v * %v = %v\n", times, a.second, a.first, a.first*a.second))
		fmt.Printf("The result is %v\n", a.first*a.second)

	case "f1f2":
		log.New(file, "INFO:", log.Ldate|log.Ltime|log.Llongfile).Printf("%v * %v = %v", a.third, a.fourth, a.third*a.fourth)
		ftxt.WriteString(fmt.Sprintf("%s: %v * %v = %v\n", times, a.third, a.fourth, a.third*a.fourth))
		fmt.Printf("The result is %v\n", a.third*a.fourth)
	default:
		fmt.Println("Sorry but we cannot process float with integer because it can cause lack of calculating.")
		log.New(file, "ERROR", log.Ldate|log.Ltime|log.Llongfile).Println("Error")
	}
}

func (a two) divide() {
	var (
		file, _ = os.OpenFile(account+"hist.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		ftxt, _ = os.OpenFile(account+"history.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		times   = time.Now().Format("Jan 02, 2006 03:04 PM")
	)
	defer closef(ftxt)
	defer closef(file)
	switch a.order {
	case "i1i2":
		log.New(file, "INFO:", log.Ldate|log.Ltime|log.Llongfile).Printf("%v / %v = %v", a.first, a.second, a.first/a.second)
		ftxt.WriteString(fmt.Sprintf("%s: %v / %v = %v\n", times, a.second, a.first, a.first/a.second))
		fmt.Printf("The result is %v\n", a.first/a.second)

	case "f1f2":
		log.New(file, "INFO:", log.Ldate|log.Ltime|log.Llongfile).Printf("%v / %v = %v", a.third, a.fourth, a.third/a.fourth)
		ftxt.WriteString(fmt.Sprintf("%s: %v / %v = %v\n", times, a.third, a.fourth, a.third/a.fourth))
		fmt.Printf("The result is %v\n", a.third/a.fourth)
	default:
		fmt.Println("Sorry but we cannot process float with integer because it can cause lack of calculating.")
		log.New(file, "ERROR:", log.Ldate|log.Ltime|log.Llongfile).Println("Error")
	}
}

func (v one) Abs() {
	var (
		file, _ = os.OpenFile(account+"hist.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		ftxt, _ = os.OpenFile(account+"history.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		times   = time.Now().Format("Jan 02, 2006 03:04 PM")
	)
	defer closef(ftxt)
	defer closef(file)
	log.New(file, "INFO:", log.Ldate|log.Ltime|log.Llongfile).Printf("Before Absolute: %v ,After Absolute: %.2f", v.variable, math.Abs(v.variable))
	ftxt.WriteString(fmt.Sprintf("%s: Abs(%v) = %.2f\n", times, v.variable, math.Abs(v.variable)))
	fmt.Printf("Your Absolute result is: %.2f\n", math.Abs(v.variable))

}

func (v one) Square() {
	var (
		file, _ = os.OpenFile(account+"hist.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		ftxt, _ = os.OpenFile(account+"history.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		times   = time.Now().Format("Jan 02, 2006 03:04 PM")
	)
	defer closef(ftxt)
	defer closef(file)
	log.New(file, "INFO:", log.Ldate|log.Ltime|log.Llongfile).Printf("Before Square: %v ,After Square: %.2f", v.variable, math.Sqrt(v.variable))
	ftxt.WriteString(fmt.Sprintf("%s: Sqrt(%v) = %.2f\n", times, v.variable, math.Sqrt(v.variable)))
	fmt.Printf("Your Square result is: %.2f\n", math.Sqrt(v.variable))

}

func (v one) Exp() {
	var (
		file, _ = os.OpenFile(account+"hist.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		ftxt, _ = os.OpenFile(account+"history.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		times   = time.Now().Format("Jan 02, 2006 03:04 PM")
	)
	defer closef(ftxt)
	defer closef(file)
	log.New(file, "INFO:", log.Ldate|log.Ltime|log.Llongfile).Printf("Before Exponing: %v ,After Exponing: %.2f", v.variable, math.Exp(v.variable))
	ftxt.WriteString(fmt.Sprintf("%s: Exp(%v) = %.2f\n", times, v.variable, math.Exp(v.variable)))
	fmt.Printf("Your Exponen result is: %.2f\n", math.Exp(v.variable))

}

func (v one) Log() {
	var (
		file, _ = os.OpenFile(account+"hist.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		ftxt, _ = os.OpenFile(account+"history.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		times   = time.Now().Format("Jan 02, 2006 03:04 PM")
	)
	defer closef(ftxt)
	defer closef(file)
	log.New(file, "INFO:", log.Ldate|log.Ltime|log.Llongfile).Printf("Before Log: %v ,After Log: %.2f", v.variable, math.Log(v.variable))
	ftxt.WriteString(fmt.Sprintf("%s: Log(%v) = %.2f\n", times, v.variable, math.Log(v.variable)))
	fmt.Printf("Your Log result is: %.2f\n", math.Log(v.variable))

}

func (v one) Sin() {
	var (
		file, _ = os.OpenFile(account+"hist.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		ftxt, _ = os.OpenFile(account+"history.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		times   = time.Now().Format("Jan 02, 2006 03:04 PM")
	)
	defer closef(ftxt)
	defer closef(file)
	log.New(file, "INFO:", log.Ldate|log.Ltime|log.Llongfile).Printf("Before Sin: %v ,After Sin: %.2f", v.variable, math.Sin(math.Pi/v.variable))
	ftxt.WriteString(fmt.Sprintf("%s: Sin(PI/%v) = %.2f\n", times, v.variable, math.Sin(math.Pi/v.variable)))
	fmt.Printf("Your Sin result is: %.2f\n", math.Sin(math.Pi/v.variable))

}

func (v one) Cos() {
	var (
		file, _ = os.OpenFile(account+"hist.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		ftxt, _ = os.OpenFile(account+"history.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		times   = time.Now().Format("Jan 02, 2006 03:04 PM")
	)
	defer closef(ftxt)
	defer closef(file)
	log.New(file, "INFO:", log.Ldate|log.Ltime|log.Llongfile).Printf("Before Cos: %v ,After Cos: %.2f", v.variable, math.Cos(v.variable))
	ftxt.WriteString(fmt.Sprintf("%s: Cos(%v) = %.2f\n", times, v.variable, math.Cos(v.variable)))
	fmt.Printf("Your Cos result is: %.2f\n", math.Cos(v.variable))

}

func (v one) Tan() {
	var (
		file, _ = os.OpenFile(account+"hist.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		ftxt, _ = os.OpenFile(account+"history.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		times   = time.Now().Format("Jan 02, 2006 03:04 PM")
	)
	defer closef(ftxt)
	defer closef(file)
	log.New(file, "INFO:", log.Ldate|log.Ltime|log.Llongfile).Printf("Before Tan: %v ,After Tan: %.2f", v.variable, math.Tan(math.Pi/v.variable))
	ftxt.WriteString(fmt.Sprintf("%s: Tan(%v) = %.2f\n", times, v.variable, math.Tan(math.Pi/v.variable)))
	fmt.Printf("Your Tan result is: %.2f\n", math.Tan(math.Pi/v.variable))

}

func (v one) Asin() {
	var (
		file, _ = os.OpenFile(account+"hist.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		ftxt, _ = os.OpenFile(account+"history.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		times   = time.Now().Format("Jan 02, 2006 03:04 PM")
	)
	defer closef(ftxt)
	defer closef(file)
	log.New(file, "INFO:", log.Ldate|log.Ltime|log.Llongfile).Printf("Before Asin: %v ,After Asin: %.2f", v.variable, math.Asin(v.variable))
	ftxt.WriteString(fmt.Sprintf("%s: Asin(%v) = %.2f\n", times, v.variable, math.Asin(v.variable)))
	fmt.Printf("Your Asin result is: %.2f\n", math.Asin(v.variable))

}

func (v one) Floor() {
	var (
		file, _ = os.OpenFile(account+"hist.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		ftxt, _ = os.OpenFile(account+"history.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		times   = time.Now().Format("Jan 02, 2006 03:04 PM")
	)
	defer closef(ftxt)
	defer closef(file)
	log.New(file, "INFO:", log.Ldate|log.Ltime|log.Llongfile).Printf("Before Floor: %v ,After Floor: %v", v.variable, math.Floor(v.variable))
	ftxt.WriteString(fmt.Sprintf("%s: Floor(%v) = %v\n", times, v.variable, math.Floor(v.variable)))
	fmt.Println("Your Floor result is:", math.Floor(v.variable))

}

func (v one) Round() {
	var (
		file, _ = os.OpenFile(account+"hist.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		ftxt, _ = os.OpenFile(account+"history.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		times   = time.Now().Format("Jan 02, 2006 03:04 PM")
	)
	defer closef(ftxt)
	defer closef(file)
	log.New(file, "INFO:", log.Ldate|log.Ltime|log.Llongfile).Printf("Before Round: %v ,After Round: %v", v.variable, math.Round(v.variable))
	ftxt.WriteString(fmt.Sprintf("%s: Round(%v) = %v\n", times, v.variable, math.Round(v.variable)))
	fmt.Println("Your Round result is:", math.Round(v.variable))

}

func main() {
	var (
		order bytes.Buffer
		type1 string
		type2 string
		type3 string
	)
	for {
		switch mode {
		case "basic":
			for {
				order.WriteString("Please type a number do you want to work with (integer,float)(10,0.8)\n \"Warning Do not type string and other symbol or there'll be consequence\"")
				fmt.Println(order.String())
				fmt.Scan(&type1)
				if checkError(crossBorder(type1)) == "y" {
					fmt.Println("You type with right ")
					order.Reset()
					break
				}
				dec()
				order.Reset()
			}

			for {
				order.WriteString("Now for the next step please type the next number you want to count")
				fmt.Println(order.String())
				fmt.Scan(&type2)
				if checkError(crossBorder(type2)) == "y" {
					fmt.Println("You type with right ")
					order.Reset()
					break
				}
				dec()
				order.Reset()
			}

			for {
				var cho string
				order.WriteString("Please choose for the menu.We serve:\n 1: adding your number with right propotion\n 2: substracting your propotion down to healty way\n 3: Our best multiplying to bigger size\n 4: And last not least we divided the number so you are not greedy")
				fmt.Println(order.String())
				procedure := sorted(type1, type2)
				fmt.Scan(&cho)
				switch cho {
				case "1":
					procedure.adding()
				case "2":
					procedure.substract()
				case "3":
					procedure.multiple()
				case "4":
					procedure.divide()
				default:
					dec()
				}
				order.Reset()
				break
			}
		case "advanced":
			for {
				order.WriteString("Please type a number do you want to work with (integer,float)(10,0.8)\n \"Warning Do not type string and other symbol or there'll be consequence\"")
				fmt.Println(order.String())
				fmt.Scan(&type3)
				if checkError(crossBorder(type3)) == "y" {
					fmt.Println("You type with right ")
					order.Reset()
					break
				}
				dec()
				order.Reset()
			}

			for {
				var cho string
				data := one{variable: 0}
				switch strings.Contains(type3, ".") {
				case true:
					n1, _ := strconv.ParseFloat(type3, 64)
					data.variable = n1
				default:
					n2, _ := strconv.Atoi(type3)
					datafirst := int32(n2)
					data.variable = float64(datafirst)
				}
				order.WriteString("Please choose for the menu.We serve:\n 1: Absoluting Your value\n 2: Squaring your value\n 3: Exponing your value\n 4: Log your value\n 5: Sin your value\n 6: Cost your value\n 7: Tan your value\n 8: Asin your value\n 9: Floor your value\n 10: Last not least rounding your value")
				fmt.Println(order.String())
				fmt.Scan(&cho)
				switch cho {
				case "1":
					data.Abs()
				case "2":
					data.Square()
				case "3":
					data.Exp()
				case "4":
					data.Log()
				case "5":
					data.Sin()
				case "6":
					data.Cos()
				case "7":
					data.Tan()
				case "8":
					data.Asin()
				case "9":
					data.Floor()
				case "10":
					data.Round()
				default:
					dec()
				}
				order.Reset()
				break
			}
		default:
			log.Fatalln("Something strange happen????")
		}
		order.WriteString("Type c if you want to end this calculating and type f if you want continue calculating and type h for history\n\"Don't type other than f,c, and h or you will face our consequences")
		fmt.Println(order.String())
		fmt.Scan(&type3)
		switch type3 {
		case "c":
			fmt.Println("Thank you for using our calculator and don't forget to practice manual counting to practice your lazy brain")
			return
		case "f":
			order.Reset()
			fmt.Println("Don't be naugty using our calculator")
		case "h":
			sehist()
			order.Reset()
			fmt.Println("Don't be naugty using our calculator")
		default:
			log.Fatalln("This is absolutly disgrace")
		}
	}

}

func sorted(t1 string, t2 string) two {
	data := two{first: 0, second: 0, third: 0, fourth: 0, order: ""}
	switch strings.Contains(t1, ".") {
	case true:
		n1, _ := strconv.ParseFloat(t1, 64)
		data.third = n1
		data.order = "f1"
	default:
		n1, _ := strconv.Atoi(t1)
		data.first = n1
		data.order = "i1"
	}
	switch strings.Contains(t2, ".") {
	case true:
		n2, _ := strconv.ParseFloat(t2, 64)
		data.fourth = n2
		data.order += "f2"
	default:
		n2, _ := strconv.Atoi(t2)
		data.second = n2
		data.order += "i2"
	}
	return data
}

func crossBorder(s string) error {
	file, _ := os.OpenFile(account+"hist.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer closef(file)
	re := regexp.MustCompile(`[0-9]`)
	if re.MatchString(s) == true && strings.ContainsAny(s, ", / ' ] [] {} [ { } ? > < ! @ # $ % ^ & * () ( ) _ - + = ` ~ ? q w e r t y u i o p l k j h g f d s a z x c v b n m") == false {
		log.New(file, "INFO", log.Ldate|log.Ltime|log.Llongfile).Println("Successed making number")
		return nil
	} else {
		log.New(file, "WARNING", log.Ldate|log.Ltime|log.Llongfile).Println("Making mistake")
		return fmt.Errorf("please type a number or float without other symbol")
	}

}

func checkError(err error) any {
	defer func() {
		if i := recover(); i != nil {
			fmt.Println(err)
			fmt.Println("If you want our to stop this you might want to type C \n(Note: Please do not ctrl C it will be causing some file not closing properly and unrecorded log)")
		}

	}()
	if err != nil {
		panic(err)
	} else {
		return "y"
	}

}

func dec() {
	file, _ := os.OpenFile(account+"hist.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer closef(file)
	decid := ""
	for {
		fmt.Println("Do you want to continue despite your mistake: (y or n)")
		fmt.Scan(&decid)
		if decid == "y" {
			log.New(file, "INFO", log.Ldate|log.Ltime|log.Llongfile).Println("Being stubborn")
			break
		} else if decid == "n" {
			log.New(file, "INFO", log.Ldate|log.Ltime|log.Llongfile).Println("Giveup")
			log.Fatalln("Thank you for your journey with us and we hope for better improvement")
		} else {
			continue
		}
	}

}

func sehist() {
	file, err := os.Open(account + "history.txt")
	if err != nil {
		fmt.Println("Sorry buddy but history is not been made and we can't find it:", err)
	}
	defer closef(file)
	read := bufio.NewScanner(file)
	for read.Scan() {
		fmt.Println(read.Text())
	}
}

func closef(f io.Closer) {
	err := f.Close()
	if err != nil {
		log.Fatalln("Error closing file:", err)
	}
}
