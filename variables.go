package main

import "fmt"

// cannot use shorthand in package level
var outerVar = 10

func main() {
	fmt.Println(outerVar)

	var uninitialized string
	fmt.Println(uninitialized)

	// inferred types
	var stringVar1 = "String"
	var intVar1 = 7
	var floatVar1 = 3.14
	fmt.Println(stringVar1)
	fmt.Println(intVar1)
	fmt.Println(floatVar1)

	// explicitly defined types
	var stringVar2 string = "String"
	var intVar2 int = 7
	var floatVar2 float64 = 3.14
	fmt.Println(stringVar2)
	fmt.Println(intVar2)
	fmt.Println(floatVar2)

	// short-hand
	name := "John Smith"
	age := 25
	grade := 8.7
	isEmployed := true

	fmt.Println(name, age, grade, isEmployed)
}
