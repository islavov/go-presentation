package main

import "fmt"

// BEGIN
func main() {
	defer func() { // HL
		fmt.Println("Край на 2016 година")
	}() // HL

	fmt.Println("Реч на президента")
	fmt.Println("Обратно броене от 10")
}
// END
