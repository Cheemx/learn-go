package main

import "fmt"

func main() {
	mep := make(map[int]int, 0)

	arr := []int{12, 10, 10, 10, 10, 5, 8, 8, 12, 5, 8, 10, 5, 12}

	for i := 0; i < len(arr); i++ {
		mep[arr[i]]++
	}
	fmt.Println(mep)
}
