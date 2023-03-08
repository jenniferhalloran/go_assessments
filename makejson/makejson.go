/*
1. prompt user to enter a name
2. prompt user to enter an address.
3. create a map and add he name and address to the map using the keys "name" and "address".
4. Use Marshal() to create json and print the JSON
*/

package main

import ( 
	"fmt"
	"os"
	"bufio"
	"encoding/json"
)


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("Please enter your name.")
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("Please enter your address.")
	scanner.Scan()
	address := scanner.Text()

	dataMap := map[string]string {"name": name, "address": address}

	dataJson, _ := json.Marshal(dataMap)

	fmt.Println("Generated JSON: \n" + string(dataJson))
}
