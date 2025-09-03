package basics

import (
	"encoding/base64"
	"fmt"
)

func main() {

	data := []byte("He~lo, Base64 Encoding")

	//Encode to Base64
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println(encoded)

	// Decode from Base64
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error in decoding:", err)
	}
	fmt.Println("Decoded", string(decoded))

	// URL safe, avoid '/' and '+'

	urlSafeencoded := base64.URLEncoding.EncodeToString(data)
	fmt.Println("URL Safe encoded:", urlSafeencoded)
}
