package utils

import (
	"bufio"
	"fmt"
	"os"
)

/*
Reads a text file line by line.
*/
func LineByLine(input string) []string {
	var s []string
	// Open the file
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		s = append(s, line)
	}
	return s
}
