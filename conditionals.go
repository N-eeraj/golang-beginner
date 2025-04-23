package main

import "fmt"

func main() {
	var age uint8
	fmt.Printf("Enter Age: ")
	fmt.Scanln(&age)

	if age > 18 {
		fmt.Println("Is Adult")
	} else {
		fmt.Println("Isn't Adult")
	}
}
