package main

import "fmt"

func minMaxSum(arr []int32) {
	var min, max, sum int32 = arr[0], arr[0], arr[0]

	for i := 1; i < len(arr); i++ {
		sum += arr[i]
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}

	fmt.Printf("%d %d\n", sum - max, sum - min)
}

func main() {
	arr := []int32{1, 3, 5, 7, 9}
	minMaxSum(arr)
}