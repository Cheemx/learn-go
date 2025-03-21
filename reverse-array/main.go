package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the size of array you want to reverse")
	fmt.Scan(&n)
	arr := make([]int, 0)
	fmt.Println("Enter numbers you want to add one by one")
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		arr = append(arr, num)
	}
	fmt.Println("Array in normal flow", arr)
	for i, j := 0, n-1; i <= j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	fmt.Println("Reversed array is:", arr)
}
