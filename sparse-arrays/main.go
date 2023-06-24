package main

import "fmt"

func matchingStrings(sStrings []string, sQuerys []string) []int32 {
	results := make([]int32, len(sQuerys))

	for i, query := range sQuerys {
		count := int32(0)
		for _, str := range sStrings {
			if str == query {
				count += 1
			}
		}
		results[i] = count
	}

	return results
	
}

func main() {
	s := []string{"ab", "ab", "abc"}
	q := []string{"ab", "abc", "bc"}

	fmt.Println(matchingStrings(s, q))
}