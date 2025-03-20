package main

import "fmt"

func main() {
	var inp string

	fmt.Println("Enter the string to get length")
	fmt.Scan(&inp)
	fmt.Println("Length of the entered string:", len(inp))
}
