package main

import "fmt"

func main() {
	invincible := map[string]string {
		"Invincible": "Mark Grayson",
		"Omni-Man": "Nolan Grayson",
		"Atom Eve": "Samantha Eve Wilkins",
	}

	fmt.Println(invincible)
	fmt.Println(invincible["Omni-Man"])

	invincible["Kid Omni-Man"] = "Oliver Grayson"

	for superHeroName, realName := range invincible {
		fmt.Printf("%v goes by the hero name %q\n", realName, superHeroName)
	}
}
