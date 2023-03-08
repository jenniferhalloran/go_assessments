package main 

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"sort"
)

func main() {
	slice := make([]int, 3)
	for {

		fmt.Println("Please enter an integer to add to the slice or X to exit.")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		val, err:= strconv.Atoi(input)
			if err != nil && input == "X" {
				fmt.Println("Nice knowing ya, bye!")
				break
			} else if err != nil {
				fmt.Println("That wasn't a valid response!")
			} else {
				slice = append(slice, val)
				sort.Ints(slice)
				fmt.Println(slice)
			}
	}
}