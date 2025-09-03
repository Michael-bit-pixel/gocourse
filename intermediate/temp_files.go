package basics

import (
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	// tempfile, err := os.CreateTemp("", "temporaryFile")
	// checkError(err)

	// fmt.Println("Temporary file created:", tempfile.Name())

	// defer os.Remove(tempfile.Name())
	// defer tempfile.Close()

	tempDir, err := os.MkdirTemp("", "GocourseTempDir")
	checkError(err)

	defer os.RemoveAll(tempDir)

	fmt.Println("Temporary Directory created:", tempDir)
}
