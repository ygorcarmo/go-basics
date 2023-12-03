package main

import "fmt"

func main() {
	var choice string
	var number1, number2 int
	fmt.Println("Welcome to The Calculator")
	fmt.Println("==================================")
	fmt.Println("Choose One of the next operations: (add, substract, divide, multiply)")
	fmt.Scan("%s", &choice)
	fmt.Println("Insert number 1:")
	fmt.Scanf("%d\n", number1)
	fmt.Println("Insert number 2:")
	fmt.Scanf("%d\n", number2)

	switch choice {
	case "add":
		fmt.Println(number1 + number2)
	case "substract":
		fmt.Println(number1 - number2)
	default:
		fmt.Println("Wrong choice")

	}

}
