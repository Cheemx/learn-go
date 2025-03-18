package main

import "fmt"

func main() {
	var i, j int

	fmt.Println("Enter 2 numbers for addition:")
	fmt.Scan(&i, &j)
	fmt.Println("Sum of", i, "and", j, "is:", i+j)
}
