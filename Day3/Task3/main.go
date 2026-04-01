package main

import "fmt"

func main() {
	students := make(map[string]int)
	students["shabin"] = 81
	students["john"] = 51
	students["rahul"] = 11

	fmt.Println("Shabin Grade", students["shabin"])

	for name, grade := range students {
		fmt.Println(name, ": ", grade)
	}
}
