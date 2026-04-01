package main

import "fmt"

func RemoveDuplicate(nums []int) []int {
	unique := []int{}
	seen := make(map[int]bool)

	for _, value := range nums {
		if !seen[value] {
			seen[value] = true
			unique = append(unique, value)
		}
	}
	return unique
}

func main() {
	nums := []int{1, 1, 1, 2, 33, 44, 22, 33}
	result := RemoveDuplicate(nums)
	fmt.Println(result)
}
