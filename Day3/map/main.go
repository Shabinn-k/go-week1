package main

import "fmt"

func main() {
	age := make(map[string]int)

	name := map[string]string{
		"head": "shabin",
		"tail": "john",
	}
	name["initial"] = "k"

	age["shabin"] = 18
	age["johncena"] = 30
	age["shabin"] = 55

	delete(age, "johncena")

	for key, value := range name {
		fmt.Println(key, value)
	}
	fmt.Println(name)
	fmt.Println(age)
}
