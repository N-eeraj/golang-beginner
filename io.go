package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func main() {
	fmt.Print("Inline Print\n") // without formatting
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

	// Formatted print
	fmt.Printf("Var %v\n", 123) // variable
	fmt.Printf("Quoted var %q\n", "John") // quoted
	fmt.Printf("Float %f\n", 5.5) // float
	fmt.Printf("Rounded Float %0.2f\n", 5.5) // float rounded to 2 decimal points

	input1 := readLine("Enter a line: ")
	fmt.Println(input1)
	input2 := readLine("Enter another line: ")
	fmt.Println(input2)
}

// custom function to read line input
func readLine(prompt string) string {
	fmt.Print(prompt) // message to print prior to input
	reader := bufio.NewReader(os.Stdin) // creating a reader
	input, _ := reader.ReadString('\n') // read entered input up to a new-line
	return strings.TrimSpace(input) // return trimmed input
}
