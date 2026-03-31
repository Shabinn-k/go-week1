package main

import "fmt"

func main() {
	for {
		fmt.Println("\n CRUDE APP")
		fmt.Println("1. CREATE")
		fmt.Println("2. READ")
		fmt.Println("3. UPDATE")
		fmt.Println("4. DELETE")
		fmt.Println("5. EXIT")

		fmt.Println("Enter your choice")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("CREATE SELECTED")
		case 2:
			fmt.Println("READ SELECTED")
		case 3:
			fmt.Println("UPDATE SELECTED")
		case 4:
			fmt.Println("DELETE SELECTED")
		case 5:
			fmt.Println("EXIT SELECTED")
			return
		default:
			fmt.Println("OVERSMART !")
		}
	}
}
