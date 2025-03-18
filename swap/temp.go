package main

import "fmt"

func main() {
	a := 5
	b := 8

	fmt.Println("Values of a and b before swap respectively:", a, b)
	temp := a
	a = b
	b = temp
	fmt.Println("Values of a and b after swap respectively:", a, b)
}
