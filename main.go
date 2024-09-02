package main

import (
	"fmt"
)

func main() {
	var length int
	fmt.Scanln(&length)

	arr := make([]int, 0, length)
	sum := 0
	for i := 1; i <= length; i++ {
		var num int
		fmt.Scanln(&num)
		sum += num
		arr = append(arr, sum)
	}
	for {
		var (
			begin int
			end   int
		)
		fmt.Scanf("%d %d", &begin, &end)
		if begin == 0 {
			fmt.Println(arr[end])
		} else {
			fmt.Println(arr[end] - arr[begin-1])
		}
	}
}
