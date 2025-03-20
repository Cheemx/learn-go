package main

import "fmt"

func main() {
	var inp string

	fmt.Println("Enter string to reverse:")
	fmt.Scan(&inp)
	s := []rune(inp)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println("Reversed string is:", string(s))

}
