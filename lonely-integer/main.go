package main

import "fmt"

func lonelyIntegerMap(a []int32) int32 {
	occurs := make(map[int32]int32) 

	for _, n := range a {
		if _, ok := occurs[n]; ok {
			delete(occurs, n)
		} else {
			occurs[n] = 1
		}
	}

	for n := range occurs {
		return n
	}

	return -1

}


func lonelyInteger(a []int32) int32 {
	var result int32

	for _, n := range a {
		result ^= n
	}
	return result
}

func main() {
	n := []int32{1, 2, 3, 4, 3, 2, 1}
	fmt.Println("Result: ", lonelyIntegerMap(n))
}