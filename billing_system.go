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

// find & add item to bill items
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

func (b *bill) updateBillItem(itemIndex int, quantity uint64) {
	b.items[itemIndex].quantity = quantity
	fmt.Printf("Updated %v to %v\n", b.items[itemIndex].name, quantity)
	b.view()
}

// find & remove item from bill
func (b *bill) removeBillItem(itemIndex int) {
	item := b.items[itemIndex]
	b.items = append(b.items[:itemIndex], b.items[itemIndex + 1:]...)
	fmt.Printf("Removed %v from the bill\n", item.name)
	b.view()
}

// handle adding dish & prompting quantity
func (b *bill) addItem(dishIndex int) {
	dish := dishMenu[dishIndex]
	quantity := getQuantity()
	item := billItem{
		name: dish.name,
		price: dish.price,
		quantity: quantity,
	}
	b.addBillItem(item)
}

func (b *bill) updateItem() {
	for {
		b.view()
		fmt.Println("Enter 'x' to cancel")
		updateIndex := scanner("Select item to update: ")
		if updateIndex == "x" {
			break
		}
		selection, error := strconv.Atoi(updateIndex)
		if error != nil || selection <=0 || selection > len(b.items) {
			fmt.Println("Invalid Selection")
		} else {
			quantity := getQuantity()
			b.updateBillItem(selection - 1, quantity)
			break
		}
	}
}

// view bill & prompt for item to remove
func (b *bill) removeItem() {
	for {
		b.view()
		fmt.Println("Enter 'x' to cancel")
		deleteIndex := scanner("Select item to remove: ")
		if deleteIndex == "x" {
			break
		}
		selection, error := strconv.Atoi(deleteIndex)
		if error != nil || selection <=0 || selection > len(b.items) {
			fmt.Println("Invalid Selection")
		} else {
			b.removeBillItem(selection - 1)
			break
		}
	}
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
func createInvoice() bill {
	billNo := len(savedBills) + 1
	invoice := "GBI" + strings.Repeat("0", 5 - len(strconv.Itoa(billNo))) + strconv.Itoa(billNo)
	bill := bill{
		invoice: invoice,
		items: []billItem{},
	}
	fmt.Printf("\nCreating Bill: %v\n", bill.invoice)
	return bill
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

	for {
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

func saveBill(bill bill) {
	savedBills = append(savedBills, bill)
	fmt.Printf("Saved Bill: %v\n", bill.invoice)
}

// function that handles billing
func createBill() {
	bill := createInvoice()

	for {
		fmt.Println()
		billMenuInput := showBillMenu()
		switch billMenuInput {
			case 1:
				dishMenuInput := showDishMenu()
				if dishMenuInput == -1 {
					continue
				}
				bill.addItem(dishMenuInput)
			case 2:
				if len(bill.items) == 0 {
					fmt.Println("No Items")
					continue
				}
				bill.updateItem()
			case 3:
				if len(bill.items) == 0 {
					fmt.Println("No Items")
					continue
				}
				bill.removeItem()
			case 4:
				if len(bill.items) == 0 {
					fmt.Println("No items to save")
					continue
				}
				saveBill(bill)
				return
			case 5:
				return
		}
	}
}

// list & view saved bills
func showSavedBills() {
	fmt.Println()
	if len(savedBills) == 0 {
		fmt.Println("No Saved Bills")
	} else {
		for index, bill := range savedBills {
			fmt.Printf("%v. %v\n", index + 1, bill.invoice)
		}
		fmt.Println("Enter 'x' to go back")
		for {
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
	for {
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
