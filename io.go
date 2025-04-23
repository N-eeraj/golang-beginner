package main

import "fmt"

func main() {
	fmt.Printf("Inline Print\n") // with formatting
	fmt.Println("Multi-line Print")

	// Read by type
	var str1 string
	var int1 int
	fmt.Printf("Enter String: ")
	fmt.Scanf("%s", &str1)
	fmt.Printf("Enter Num: ")
	fmt.Scanf("%d", &int1)
	fmt.Printf("str1 is %s & int1 is %d\n", str1, int1)

	// Just Read
	var str2 string
	var int2 int
	fmt.Printf("Enter String: ")
	fmt.Scanln(&str2)
	fmt.Printf("Enter Num: ")
	fmt.Scanln(&int2)
	fmt.Printf("str2 is %s & int2 is %d\n", str2, int2)
}
