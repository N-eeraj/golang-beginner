package main

import (
	"fmt"
	"strings"
)

func main() {
	message := "Hello World!"

	// Compare Method
	fmt.Println(strings.Compare("a", "a")) // returns 0 for equal
	fmt.Println(strings.Compare("a", "b")) // a < b => -1
	fmt.Println(strings.Compare("b", "a")) // b > a => 1

	// Contains Method
	fmt.Println(strings.Contains(message, "Hello")) // true
	fmt.Println(strings.Contains(message, "hey")) // false

	// ContainsAny Method
	fmt.Println(strings.ContainsAny(message, "H")) // true
	fmt.Println(strings.ContainsAny(message, "h")) // false
	fmt.Println(strings.ContainsAny(message, "Hh")) // true
	fmt.Println(strings.ContainsAny(message, "lo")) // true
	fmt.Println(strings.ContainsAny(message, "")) // false

	// Count Method
	fmt.Println(strings.Count(message, "h")) // 0
	fmt.Println(strings.Count(message, "H")) // 1
	fmt.Println(strings.Count(message, "l")) // 3
	fmt.Println(strings.Count(message, "")) // 13

	// Cut Method
	cutBefore, cutAfter, cutFound := strings.Cut(message, "o ")
	fmt.Printf("%q, %q, %v\n", cutBefore, cutAfter, cutFound) // Hell, World!, true

	// CutPrefix Method
	cutPrefixAfter, cutPrefixFound := strings.CutPrefix(message, "Hel")
	fmt.Printf("%q, %v\n", cutPrefixAfter, cutPrefixFound) // lo World!, true

	// CutSuffix Method
	cutSuffixBefore, cutSuffixFound := strings.CutSuffix(message, "World")
	fmt.Printf("%q, %v\n", cutSuffixBefore, cutSuffixFound) // Hello World!, false

	// Replace Method
	fmt.Println(strings.Replace(message, "l", "L", 2)) // Replaces first 2 "l" with "L"
	fmt.Println(strings.Replace(message, "l", "L", -1)) // Replaces all "l" with "L"

	// ReplaceAll Method
	fmt.Println(strings.ReplaceAll(message, "l", "L")) // Replaces all "l" with "L"

	// ToUpper & ToLower Methods
	fmt.Println(strings.ToUpper(message)) // HELLO WORLD!
	fmt.Println(strings.ToLower(message)) // hello world!

	// Index Method
	fmt.Println(strings.Index(message, "l")) // 2
	fmt.Println(strings.Index(message, "x")) // -1

	// Split Method
	fmt.Println(strings.Split(message, ""))
	fmt.Println(strings.Split(message, " "))
	fmt.Println(strings.Split(message, "l"))

	fmt.Println(message)
}
