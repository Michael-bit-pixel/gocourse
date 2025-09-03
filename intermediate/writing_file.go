package basics

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("Output.txt")
	if err != nil {
		fmt.Println("Error creating file.", file)
	}
	defer file.Close()

	// write data to file
	data := []byte("Hello World!\n\n\n")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}

	fmt.Println("Data has been written to file succesfully.")

	file, err = os.Create("writeString.txt")
	if err != nil {
		fmt.Println("Error creating file", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Hello Go!\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Writing to writeString.txt complete .")
}
