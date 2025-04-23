package main

import "fmt"

func main() {
	num1 := 10
	num2 := 7

	// Arithmetic
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	difference := num1 - num2
	fmt.Println("Difference:", difference)

	prod := num1 * num2
	fmt.Println("Product:", prod)

	quotient := num1 / num2
	fmt.Println("Quotient:", quotient)

	remainder := num1 % num2
	fmt.Println("Remainder:", remainder)

	// Comparison
	fmt.Println(num1 > num2)
	fmt.Println(num1 < num2)
	fmt.Println(num1 == num2)
	fmt.Println(num1 != num2)
	fmt.Println(num1 >= num2)
	fmt.Println(num1 <= num2)

	// Logical
	fmt.Println(true && true)
	fmt.Println(false || true)
	fmt.Println(!false)
}
