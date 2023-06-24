package main

import "fmt"

func plusMinus(arr []int32) {
	var pos, neg, zer float32
	size := float32(len(arr))
	for _, n := range arr {
		if n == 0 {
			zer += 1
		} else if n < 0 {
			neg += 1
		} else {
			pos += 1
		}
	}
	fmt.Printf("%.6f\n%.6f\n%.6f\n",
		ratio(pos, size),
		ratio(neg, size),
		ratio(zer, size))
}

func ratio(count, size float32) float32 {
	return count / size
}

func main() {
	arr := []int32{1, 1, 0, -1, -1}
	plusMinus(arr)
}
