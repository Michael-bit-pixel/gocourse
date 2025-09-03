package basics

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("Output.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNumber := 1

	// Keyword to filter lines
	keyboard := "important"

	// Read and filter lines
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyboard) {
			updatedLine := strings.ReplaceAll(line, keyboard, "necessary")
			fmt.Printf("%d Filtered line: %v\n", lineNumber, line)
			fmt.Printf("%d Updated line: %v\n", lineNumber, updatedLine)
			lineNumber++
		}
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error scanning file", err)
	}
}
