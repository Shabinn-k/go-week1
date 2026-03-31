package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	number := []int{1, 2, 3, 4, 5}
	name := []string{"shabin"}

	for _, value := range name {
		fmt.Println(value)
	}

	for index, value := range number {
		fmt.Println(index, value)
	}
}
