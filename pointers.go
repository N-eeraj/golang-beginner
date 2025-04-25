package main

import "fmt"

func updateIntByPointer(ptr *int, value int) {
	*ptr = value
}

func updateStringByPointer(ptr *string, value string) {
	*ptr = value
}

func main() {
	var1 := 4
	ptr1 := &var1
	fmt.Printf("Address of var1: %v\n", &var1)
	fmt.Printf("Pointer of var1: %v\n", ptr1)
	fmt.Printf("Value of var1: %v & pointer value at ptr1: %v\n", var1, *ptr1)

	var2 := "hello"
	ptr2 := &var2
	fmt.Printf("Address of var2: %v\n", &var2)
	fmt.Printf("Pointer of var2: %v\n", ptr2)
	fmt.Printf("Value of var2: %v & pointer value at ptr2: %v\n", var2, *ptr2)

	updateIntByPointer(ptr1, 7)
	fmt.Printf("Updated value of var1: %v & pointer value at ptr1: %v\n", var1, *ptr1)

	updateStringByPointer(ptr2, "world")
	fmt.Printf("Updated value of var2: %v & pointer value at ptr2: %v\n", var2, *ptr2)
}
