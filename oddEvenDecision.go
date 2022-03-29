package main

import "fmt"

func main() {
	var input int

	fmt.Println("Enter number:")
	fmt.Scanln(&input)

	if input%2 == 0 {
		fmt.Println(input, "is even.")
	} else if input%2 == 1 {
		fmt.Println(input, "is odd.")
	}
}
