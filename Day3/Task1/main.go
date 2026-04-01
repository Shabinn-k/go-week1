package main

import "fmt"

func maxMin(arr [5]int) (int, int) {
	max := arr[0]
	min := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}
	return max, min

}
func main() {
	arr := [5]int{23, 13, 1, 532, 65}
	max, min := maxMin(arr)
	fmt.Println(max)
	fmt.Println(min)
}
