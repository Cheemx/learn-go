package main

import "fmt"

func main() {
	var n int

	fmt.Println("Enter the size of array you want:")
	fmt.Scan(&n)
	// arr := [n]int{}
	arr := make([]int, 0)
	fmt.Println("Enter the elements one by one")
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		arr = append(arr, num)
	}
	fmt.Print(arr)
}
