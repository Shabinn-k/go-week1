package main

import "fmt"

func main() {
	// day:="monday"
	// switch day{
	// 	case "monday":
	// 		fmt.Println("Start")

	// 	default :
	// 		fmt.Println("end")
	// }

	num := 10
	switch {
	case num%2 != 0:
		fmt.Println("odd")
	case num%2 == 0:
		fmt.Println("even")
	default:
		fmt.Println("other")
	}

}
