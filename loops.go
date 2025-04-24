package main

import "fmt"

func main() {
	// classic for loop
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Skipping", i)
			continue
		}
		fmt.Println("For loop:", i)
	}

	// while loop "kinda"
	i := 0
	for i < 10 {
		fmt.Println("While loop:", i)
		i++
	}

	// infinite loop
	i = 0
	for {
		if i > 10 {
			break
		}
		fmt.Println("Infinite Loop:", i)
		i++
	}

	words := []string {"go", "is", "fun"}
	for index, word := range words {
		fmt.Println(index, word)
	}

}
