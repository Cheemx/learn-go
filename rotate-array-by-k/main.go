package main

import "fmt"

func main() {
	arr := []int{12, 15, 88, 36, 7, 5, 11}
	fmt.Println("Following is the array:")
	fmt.Println(arr)
	var k int
	fmt.Println("Enter k(<7) to rotate array")
	fmt.Scan(&k)
	ans := make([]int, 0)
	for i := k; i < len(arr); i++ {
		ans = append(ans, arr[i])
	}
	for i := 0; i < k; i++ {
		ans = append(ans, arr[i])
	}
	fmt.Println("Following is the rotated array:")
	fmt.Println(ans)
}
