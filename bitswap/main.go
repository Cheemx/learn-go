package main

import "fmt"

func main() {
	a := 5
	b := 13

	fmt.Println("a and b before swapping:", a, b)
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println("a and b after swapping:", a, b)
}
