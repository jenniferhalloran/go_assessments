package main

import "fmt"

func main() {
	fmt.Println("Hello, please enter a floating point number.")

	var input float32

	fmt.Scanln(&input)

	truncated_input := int(input)

	fmt.Println("You entered", input,", your truncated input is", truncated_input)
}