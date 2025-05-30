package main

import "fmt"

// Custom Type/Structure
type bill struct {
	amount float32
	tax float32
	isDeliveryOrder bool
}

// Receiver Functions
func (b bill) getTotal() float32 {
	taxedAmount := b.amount * (1 + b.tax / 100)
	if (b.isDeliveryOrder) {
		return taxedAmount + 10
	}
	return taxedAmount
}

func (b bill) getTax() float32 {
	return b.amount * (1 + b.tax / 100) - b.amount
}

// use pointer in receiver functions to handle update
func (b *bill) updateTotal(newAmt float32) {
	b.amount = newAmt
}

func main() {
	// Initialization
	bill1 := bill{
		amount: 320,
		tax: 5,
		isDeliveryOrder: false,
	}

	fmt.Println(bill1)
	fmt.Printf("%0.2f\n", bill1.getTotal())
	fmt.Printf("Tax: %0.2f\n", bill1.getTax())

	bill2 := bill{
		amount: 760,
		tax: 4.5,
		isDeliveryOrder: true,
	}

	fmt.Println(bill2)
	fmt.Printf("%0.2f\n", bill2.getTotal())
	fmt.Printf("Tax: %0.2f\n", bill2.getTax())

	bill2.updateTotal(800)
	fmt.Printf("%0.2f\n", bill2.getTotal())
	fmt.Printf("Tax: %0.2f\n", bill2.getTax())
}
