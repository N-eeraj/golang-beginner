package main

import "fmt"

func main() {
	// arrays
	var avengers [6]string = [6]string{"Tony", "Steve", "Thor", "Bruce", "Clint", "Natasha"}
	fmt.Println(avengers)

	var justiceLeague = [6]string{"Bruce", "Clark", "Diana", "Barry", "Arthur", "Victor"}
	fmt.Println(justiceLeague)

	saiyans := [7]string{"Kakarot", "Vegeta", "Gohan", "Broly", "Trunks", "Goten", "Bardock"}
	fmt.Println(saiyans)

	boys := [...]string{"Butcher", "Hughie", "MM", "Frenchie", "Kimiko", "Annie"} // n number of elements in array
	fmt.Println(boys)

	fmt.Println(avengers[0])
	fmt.Println(justiceLeague[0])
	fmt.Println(saiyans[1])

	saiyans[0] = "Goku"

	fmt.Printf("There are %v Avengers\n", len(avengers))
	for index, name := range avengers {
		fmt.Printf("%v. %v\n", index + 1, name)
	}

	fmt.Printf("There are %v Justice League Members\n", len(justiceLeague))
	for index, name := range justiceLeague {
		fmt.Printf("%v. %v\n", index + 1, name)
	}

	fmt.Printf("There are %v Sayians\n", len(saiyans))
	for index, name := range saiyans {
		fmt.Printf("%v. %v\n", index + 1, name)
	}

	// slices
	titans := []string{"Robin", "Raven", "Starfire", "Beast Boy", "Cyborg"}
	titans[0] = "Nightwing"
	titans = append(titans, "Kid Flash")
	titans = append(titans, "Super Boy")
	titans = append(titans, "Robin")
	for index, name := range titans {
		fmt.Printf("%v. %v\n", index + 1, name)
	}

	// ranges
	fmt.Println(avengers[2:4])
	fmt.Println(titans[5:])
	fmt.Println(saiyans[:2])
}
