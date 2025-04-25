package main

import "fmt"

func main() {
	var num1 float64
	var num2 float64
	var option uint8
	var result float64
	symbol := ""
	menu := [4]string{"Addition", "Subtraction", "Multiplication", "Division"}

	fmt.Printf("Enter Number 1: ")
	fmt.Scan(&num1)

	fmt.Printf("Enter Number 2: ")
	fmt.Scan(&num2)

	for index, option := range menu {
		fmt.Printf("%v. %v\n", index + 1, option)
	}
	fmt.Printf("Select Option: ")
	fmt.Scan(&option)

	switch option {
		case 1:
			symbol = "+"
			result = num1 + num2
		case 2:
			symbol = "-"
			result = num1 - num2
		case 3:
			symbol = "*"
			result = num1 * num2
		case 4:
			symbol = "÷"
			result = num1 / num2
		default:
			fmt.Println("Invalid Selection")
	}

	if symbol != "" {
		fmt.Printf("%v %v %v = %v\n", num1, symbol, num2, result)
	}
}
