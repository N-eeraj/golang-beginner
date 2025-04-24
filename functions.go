package main

import (
	"fmt"
	"math"
	"strings"
)

func fn1() {
	fmt.Println("fn1")
}

func greet(name string) {
	fmt.Printf("Hello %v, welcome\n", name)
}

func add(num1 int, num2 int) {
	fmt.Println(num1 + num2)
}

func funcArg(fn func()) {
	fmt.Println("Going to call function passed")
	fn()
	fmt.Println("Called function")
}

func listString(list []string, fn func(string, int)) {
	for index, value := range list {
		fn(value, index)
	}
}

func slNoName(name string, index int) {
	fmt.Printf("Sl No. %v: %q\n", index, name)
}

func getSquare(num int) int {
	return num * num
}

func getCircleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

func getInitials(fullName string) (string, string) {
	var initials []string
	for _, name := range strings.Split(strings.ToUpper(fullName), " ") {
		if len(name) > 0 {
			initials = append(initials, name[:1])
		}
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	if (len(initials) == 1) {
		return initials[0], ""
	}
	return "", ""
}

func main() {
	fn1()
	greet("Neo")
	greet("Trinity")
	greet("Morpheus")
	add(50, 25)

	// passing function as an argument
	funcArg(func() {fmt.Println("Hey")})
	funcArg(fn1)

	androids := []string{"Ash", "Bishop", "Call", "David", "Walter", "Andy"}
	listString(androids, slNoName)

	fmt.Println(getSquare(4))
	fmt.Printf("%0.2f\n", getCircleArea(4.5))

	names := []string{"John Connor", "Sarah Jeanette Connor", "Reese", ""}
	for _, name := range names {
		firstInitial, lastInitial := getInitials(name)
		fmt.Printf("%q: First Initial: %q & Last Initial %q\n", name, firstInitial, lastInitial)
	}
}
