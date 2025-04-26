package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

// types
type billItem struct {
	name string
	price float64
	quantity uint64
}
type bill struct {
	invoice string
	items []billItem
	deliveryCharge float64
}

// common variables
var reader = bufio.NewReader(os.Stdin)
var mainMenu = [...]string{
	"New Bill",
	"Exit",
}
var billMenu = [...]string{
	"Add Item",
	"Update Item Count",
	"Remove Item",
	"Save Bill",
	"Cancel",
}
var billsGenerated = 0

// common input function
func scanner(prompt string) string {
	fmt.Printf(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// shows main menu & returns selection
func showMainMenu() int {
	for i, option := range mainMenu {
		fmt.Printf("%v. %v\n", i + 1, option)
	}
	mainMenuInput := scanner("Select your option: ")
	selection, error := strconv.Atoi(mainMenuInput)
	if error != nil || selection <=0 || selection > len(mainMenu) {
		fmt.Println("Invalid Selection")
		showMainMenu()
	}
	return selection
}

func showBillMenu() int {
	for i, option := range billMenu {
		fmt.Printf("%v. %v\n", i + 1, option)
	}
	billMenuInput := scanner("Select your option: ")
	selection, error := strconv.Atoi(billMenuInput)
	if error != nil || selection <=0 || selection > len(billMenu) {
		fmt.Println("Invalid Selection")
		showBillMenu()
	}
	return selection
}

func createInvoice() string {
	billNo := billsGenerated + 1
	return "GBI" + strings.Repeat("0", 5 - len(strconv.Itoa(billNo))) + strconv.Itoa(billNo)
}

func createBill() {
	invoice := createInvoice()
	fmt.Printf("Creating Bill: %v\n", invoice)
	for true {
		billMenuInput := showBillMenu()
		switch billMenuInput {
			case 4:
				billsGenerated++
				fmt.Printf("Saved Bill: %v\n", invoice)
				return
			case 5:
				return
		}
	}
}

func main() {
	fmt.Println("Welcome to Go Bill")
	fmt.Println("------------------")
	for true {
		mainMenuInput := showMainMenu()
		switch mainMenuInput {
			case 1:
				createBill()
			case 2:
				fmt.Println("Thank you for using Go Bill, have a nice day")
				return
		}
	}
}
