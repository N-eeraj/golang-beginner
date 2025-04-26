package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

// types
type dish struct {
	name string
	price float64
}
type billItem struct {
	name string
	price float64
	quantity uint64
}
type bill struct {
	invoice string
	items []billItem
}

// show items in current bill
func (b *bill) view () {
	header := fmt.Sprintf("| Sl.| %-25v| %7v | %8v | %9v |", "Dish", "Price", "Quantity", "Total")
	printLine := func() {
		for i:= 0; i < len(header); i++ {
			fmt.Printf("-")
		}
		fmt.Printf("\n")
	}
	printLine()
	fmt.Printf("| Bill No: %-53v |\n", b.invoice)
	printLine()
	fmt.Println(header)
	printLine()
	grandTotal := 0.0
	for index, dish := range b.items {
		total := dish.price * float64(dish.quantity)
		grandTotal += total
		fmt.Printf("| %2v | %-25v| $%6v | %8v | $%8.2f |\n", index + 1, dish.name, dish.price, dish.quantity, total)
	}
	printLine()
	fmt.Printf("| %-50v | $%8.2f |\n", "Grand Total", grandTotal)
	printLine()
}

// show items in current bill
func (b *bill) addBillItem (item billItem) {
	foundIndex := -1
	for index, billItem := range b.items {
		if item.name == billItem.name {
			foundIndex = index
		}
	}
	if foundIndex == -1 {
		b.items = append(b.items, item)
	} else {
		b.items[foundIndex].quantity += item.quantity
	}
	fmt.Printf("Added %v %v to the bill\n", item.quantity, item.name)
	b.view()
}

// common variables
var reader = bufio.NewReader(os.Stdin)
var mainMenu = [...]string{
	"New Bill",
	"Show Bills",
	"Exit",
}
var billMenu = [...]string{
	"Add Item",
	"Update Item Count",
	"Remove Item",
	"Save Bill",
	"Cancel",
}
var dishMenu = [...]dish{
	{
		name: "Cheeseburger",
		price: 7.99,
	},
	{
		name: "Fried Chicken",
		price: 9.49,
	},
	{
		name: "BBQ Ribs",
		price: 13.99,
	},
	{
		name: "Mac and Cheese",
		price: 6.75,
	},
	{
		name: "Pulled Pork Sandwich",
		price: 8.50,
	},
	{
		name: "Chicken Pot Pie",
		price: 10.25,
	},
	{
		name: "Hot Dog",
		price: 4.25,
	},
}
var savedBills = []bill{}

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

// shows bill menu & returns selection
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

// returns an invoice number based on bills generated
func createInvoice() string {
	billNo := len(savedBills) + 1
	return "GBI" + strings.Repeat("0", 5 - len(strconv.Itoa(billNo))) + strconv.Itoa(billNo)
}

// show all items available
func displayDishMenu() {
	header := fmt.Sprintf("| Sl.| %-25v| %7v |", "Dish", "Price")
	printLine := func() {
		for i:= 0; i < len(header); i++ {
			fmt.Printf("-")
		}
		fmt.Printf("\n")
	}
	printLine()
	fmt.Println(header)
	printLine()
	for index, dish := range dishMenu {
		fmt.Printf("| %2v | %-25v| $%6v |\n", index + 1, dish.name, dish.price)
	}
	printLine()
	fmt.Println("Enter 'x' to exit")
}

// handle dish menu selection
func showDishMenu () int {
	displayDishMenu()

	for true {
		dishMenuInput := scanner("Select your option: ")
		if dishMenuInput == "x" {
			return -1
		}
		selection, error := strconv.Atoi(dishMenuInput)
		if error != nil || selection <=0 || selection > len(dishMenu) {
			fmt.Println("Invalid Selection")
			showDishMenu()
		}
		return selection - 1
	}
	return -1
}

// get an whole number input from user
func getQuantity() uint64 {
	input := scanner("Enter Quantity: ")
	quantity, error := strconv.Atoi(input)
	if error != nil || quantity < 1 {
		fmt.Println("Invalid Quantity")
		getQuantity()
	}
	return uint64(quantity)
}

func addItem(bill *bill, dishIndex int) {
	dish := dishMenu[dishIndex]
	quantity := getQuantity()
	item := billItem{
		name: dish.name,
		price: dish.price,
		quantity: quantity,
	}
	bill.addBillItem(item)
}

// function that handles billing
func createBill() {
	invoice := createInvoice()
	bill := bill{
		invoice: invoice,
		items: []billItem{},
	}
	fmt.Printf("\nCreating Bill: %v\n", bill.invoice)
	for true {
		fmt.Println()
		billMenuInput := showBillMenu()
		switch billMenuInput {
			case 1:
				dishMenuInput := showDishMenu()
				if dishMenuInput == -1 {
					continue
				}
				addItem(&bill, dishMenuInput)
			case 2:
				bill.view()
			case 3:
				bill.view()
			case 4:
				if len(bill.items) == 0 {
					fmt.Println("No items to save")
					continue
				}
				savedBills = append(savedBills, bill)
				fmt.Printf("Saved Bill: %v\n", invoice)
				return
			case 5:
				return
		}
	}
}

func showSavedBills() {
	fmt.Println()
	if len(savedBills) == 0 {
		fmt.Println("No Saved Bills")
	} else {
		for index, bill := range savedBills {
			fmt.Printf("%v. %v\n", index + 1, bill.invoice)
		}
		fmt.Println("Enter 'x' to go back")
		for true {
			invoiceMenuInput := scanner("Select bill to view: ")
			if invoiceMenuInput == "x" {
				break
			}
			selection, error := strconv.Atoi(invoiceMenuInput)
			if error != nil || selection <=0 || selection > len(savedBills) {
				fmt.Println("Invalid Selection")
			} else {
				savedBills[selection - 1].view()
				break
			}
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
				showSavedBills()
			case 3:
				fmt.Println("Thank you for using Go Bill, have a nice day")
				return
		}
		fmt.Println()
	}
}
