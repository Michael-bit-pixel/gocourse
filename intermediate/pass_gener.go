package intermidiate

import (
	"bytes"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	text     bytes.Buffer
	command  string
	strop    []getpass
	lenstrop []string
	finalpas string
)

type getpass struct {
	Var1 string `json:"word"`
}

type encdec struct {
	eco, edc string
}

func init() {
	text.WriteString(fmt.Sprintf("Hi %s,Welcome to password and crack generator", os.Getenv("COMPUTERNAME")))
	fmt.Println(text.String())
	text.Reset()
}

func (e encdec) encoded() {
	text.WriteString("How many number do you want to encoded (Please note if you type a word we will make sure that word become number just in case!!!)")
	fmt.Println(text.String())
	text.Reset()
	fmt.Scan(&command)
	v, _ := strconv.Atoi(command)
	for i := 0; i < v; i++ {
		e.eco = base64.URLEncoding.EncodeToString([]byte(e.eco))
	}
	text.WriteString("FYI do you want to add hash to ensure safety of encoded message choose y or n")
	fmt.Println(text.String())
	text.Reset()
	fmt.Scan(&command)
	switch command {
	case "y":
		text.WriteString("Please input an answer if you want to hash it:")
		fmt.Println(text.String())
		fmt.Scan(&command)
		has := sha512.Sum512([]byte(command))
		fmt.Println("Your encode message is:", base64.URLEncoding.EncodeToString(has[:])+e.eco)
		fmt.Println("If you want to decoded this message please remove this has:", base64.URLEncoding.EncodeToString(has[:]))
	case "n":
		fmt.Println("Your encode message is:", e.eco)
	default:
		fmt.Println("You can say n duhhh")
	}

}

func (e encdec) decoded() {
	text.WriteString("How many number do you want to decoded(Please note if you type a word we will make sure that word become number just in case!!!)")
	fmt.Println(text.String())
	fmt.Scan(&command)
	v, _ := strconv.Atoi(command)
	for i := 0; i < v; i++ {
		b, err := base64.URLEncoding.DecodeString(e.edc)
		error_handling(err)
		e.edc = string(b)
	}
	fmt.Println("Your decode message is:", e.edc)
}

func ask() {
	text.WriteString("Please choose\n 1: Generate strong password\n 2: Check the password power\n 3: Make Hash\n 4: Encode/Decode your message\n 5: For the end ")
	fmt.Println(text.String())
	fmt.Scan(&command)
	text.Reset()
}

func main() {
	for {
		ask()
		switch command {
		case "1":
			generate_rand()
		case "2":
			cheker()
		case "3":
			hashing()
		case "4":
			base46()
		case "5":
			fmt.Println("Good luck with your password journey")
			return
		default:
			dec()
		}

	}
}

func hashing() {
	text.WriteString("Please input a line with no space between to hashing ")
	fmt.Println(text.String())
	defer text.Reset()
	text.Reset()
	fmt.Scan(&command)
	hash := sha512.Sum512([]byte(command))
	text.WriteString("Do you want to add more salt length and encoded it ?\n If you want to encoding your hash please type e and if you want to salt type s and if you wan't to see your hash please type n")
	fmt.Println(text.String())
	text.Reset()
	fmt.Scan(&command)
	switch command {
	case "e":
		fmt.Println("This is your hash result with encoding:", base64.URLEncoding.EncodeToString(hash[:]))
	case "s":
		text.WriteString("Please choose salt length.Type with only integer value if not we will valued your typing")
		fmt.Println(text.String())
		fmt.Scan(&command)
		v, _ := strconv.Atoi(command)
		salt_length := make([]byte, v)
		hash_salt := append(hash[:], salt_length...)
		hash_final := sha512.Sum512(hash_salt)
		fmt.Println("This is your hash with salt:", base64.URLEncoding.EncodeToString(hash_final[:]))
	case "n":
		fmt.Printf("This is your hash: %x\n", hash)
	default:
		fmt.Printf("This is your hash asshole.YOU can just following order without ignore it: %x\n", string(hash[:]))
	}

}

func base46() {
	text.WriteString("Do you want to encode please type e and if you want to decode please type d")
	defer text.Reset()
	fmt.Println(text.String())
	fmt.Scan(&command)
	text.Reset()
	switch command {
	case "e":
		text.WriteString("Please type the text do you want to encode:")
		fmt.Println(text.String())
		text.Reset()
		ans := ""
		// Can't add space betwee ans it will cause some trouble
		fmt.Scan(&ans)
		encode := encdec{eco: strings.ReplaceAll(ans, " ", "_")}
		fmt.Println(ans)
		encode.encoded()
	case "d":
		text.WriteString("Please type the text do you want to decode:")
		fmt.Println(text.String())
		text.Reset()
		ans := ""
		// Can't add space betwee ans it will cause some trouble
		fmt.Scan(&ans)
		decode := encdec{edc: strings.ReplaceAll(ans, " ", "_")}
		decode.decoded()
	default:
		fmt.Println("Sorry but this is absolutely mistaking by not following the order above please don't be sad and repeat again")
		return
	}

}

func generate_rand() {
	for {
		file, err := http.Get("https://random-words-api.kushcreates.com/api?category=brainrot&type=capitalized&words=10")
		error_handling(err)
		defer file.Body.Close()
		readAPI, err := io.ReadAll(file.Body)
		error_handling(err)
		json.Unmarshal(readAPI, &strop)
		for _, b := range strop {
			lenstrop = append(lenstrop, b.Var1)
		}
		rand.Shuffle(len(lenstrop), func(i, j int) {
			lenstrop[i], lenstrop[j] = lenstrop[j], lenstrop[i]
		})
		for i := 0; i < 10; i++ {
			finalpas = finalpas + lenstrop[i][:rand.Intn(len(lenstrop[i]))]
		}
		rand_tag := []string{"!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "_", "+", "=", "{", "}", "[", "]", ":", ";", "<", ">", "?", "/"}
		finalpas = rand_tag[rand.Intn(len(rand_tag))] + rand_tag[rand.Intn(len(rand_tag))] + "-" + finalpas[:rand.Intn(len(finalpas))] + fmt.Sprintf("%d", rand.Intn(len(finalpas))) + fmt.Sprintf("%d", rand.Intn(len(finalpas))) + finalpas[rand.Intn(len(finalpas)):] + rand_tag[rand.Intn(len(rand_tag))]
		re := regexp.MustCompile(`[A-Za-z0-9-]{8,}`)
		if !re.MatchString(finalpas) {
			lenstrop = nil
			strop = nil
			finalpas = ""
			continue
		} else {
			fmt.Println("The greatest password have been made,now you can use it:", strings.ReplaceAll(finalpas, " ", ""))
			lenstrop = nil
			strop = nil
			finalpas = ""
			return
		}

	}
}

func cheker() {
	text.WriteString("Please input your password")
	defer text.Reset()
	fmt.Println(text.String())
	fmt.Scan(&command)
	first := regexp.MustCompile(`[.!#$%&’*+/=?^_{|}~-]`)
	second := regexp.MustCompile(`[A-Z]`)
	third := regexp.MustCompile(`[a-z]{5,}`)
	fourth := regexp.MustCompile(`[0-9]{2,}`)
	if !first.MatchString(command) {
		fmt.Println("This password doesn't contain any symbol at the beggining")
	}
	if !second.MatchString(command) {
		fmt.Println("This password doesn't contain any Capital letter")
	}
	if !third.MatchString(command) {
		fmt.Println("This password doesn't contain 5 letter")
	}
	if !fourth.MatchString(command) {
		fmt.Println("This password doesn't contain 2 number")
	}
	if !first.MatchString(command) && !second.MatchString(command) && !third.MatchString(command) && !fourth.MatchString(command) {
		fmt.Println("Your password is weak please change if you wan't your account secure")
	} else if !first.MatchString(command) && !second.MatchString(command) {
		fmt.Println("Your password is medium please add some Capital letter and symbol for security")
	} else if !third.MatchString(command) && !fourth.MatchString(command) {
		fmt.Println("Your password is medium please add some  letter and number for security")
	} else if !first.MatchString(command) && !fourth.MatchString(command) {
		fmt.Println("Your password is medium please add some Symbol and number")
	} else if !second.MatchString(command) && !third.MatchString(command) {
		fmt.Println("Your password is medium please add some Capital letter and more letter")
	} else if !first.MatchString(command) && !third.MatchString(command) {
		fmt.Println("Your password is medium please add some Symbol and some letter")
	} else if !fourth.MatchString(command) && !second.MatchString(command) {
		fmt.Println("Your password is medium please add some number and some  Capital letter")
	} else if !first.MatchString(command) || !second.MatchString(command) || !third.MatchString(command) || !fourth.MatchString(command) {
		fmt.Println("Your password is hard")
	} else {
		fmt.Println("Your password is uncrackble")
	}
	for _, b := range command {
		if strings.Count(command, string(b)) >= 2 {
			fmt.Printf("Please do not use %c repeatedly for %v times\n", b, strings.Count(command, string(b)))
			command = strings.ReplaceAll(command, string(b), "")
		}
	}
	text.Reset()
	text.WriteString("If you want to generate uncrackable paswword please type y and type n if you don't and your password is sucked then you can fix that")
	fmt.Println(text.String())
	fmt.Scan(&command)
	switch command {
	case "y":
		generate_rand()
	case "n":
		fmt.Println("Good luck with your password")
	default:
		fmt.Println("if you don't like that,you can just type n")
	}

}

func error_handling(r error) {
	if r != nil {
		log.Fatalln("Sorry something went wrong:", r)
	}
}

func dec() {
	var des string
	text.WriteString("Sorry but that is absolutely mistake by not chossing 1,2,3,4.For continuing press y and n for stop this shit")
	fmt.Println(text.String())
	text.Reset()
	fmt.Scan(&des)
	switch des {
	case "y":
		return
	case "n":
		text.WriteString("Don't be sad")
		fmt.Println(text.String())
		os.Exit(1)
	default:
		text.WriteString("Yuppp you did it again but this time you won't be dirty again")
		fmt.Println(text.String())
		os.Exit(1)
	}

}
