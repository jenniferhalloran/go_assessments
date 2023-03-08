package main 

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	fmt.Println("Please enter a string.")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	lowerInput := strings.ToLower(scanner.Text())

	if strings.HasPrefix(lowerInput, "i") && strings.HasSuffix(lowerInput, "n") && strings.Contains(lowerInput, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
