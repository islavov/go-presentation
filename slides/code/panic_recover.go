package main

import "fmt"

// START OMIT
func openPresent() {
	fmt.Println("В кутията има огромен паяк")
	panic("Настава паника")
}

func main() {
	defer func() {
		if r := recover(); r != nil { // HL
			fmt.Println(r)
			fmt.Println("Използваш чехъл.")
		}
	}()
	openPresent()
}

// END OMIT
