package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Grade int
}

func main() {
	students := []Student{
		{Name: "shabin", Age: 18, Grade: 62},
		{Name: "john", Age: 20, Grade: 12},
	}
	fmt.Println(students)
}
