package main 

import (
	"fmt"
	"os"
	"bufio"
)

type Name struct {
		fname string 
		lname string
}

func getFileName() (inputPath string, err error) {
	for {
		fmt.Println("Please enter the path for the input file to read:")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			inputPath = scanner.Text()
			break
		}
		if err := scanner.Err(); err != nil {
			return "", err
		}

		if inputPath == "" {
			fmt.Println("Nothing entered. Please try again.")
			continue
		} else {
			return inputPath, nil
		}
	}
}

func main() {
	// input, _ := getFileName()
	// fmt.Println(input)

}