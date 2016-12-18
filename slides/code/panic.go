package main

import "fmt"

// START OMIT
func amapanic() {
	panic("Това не трябва да се случва") // HL
}

func causepanic() {
	defer func() {
		fmt.Println("Отложените функции се изпълняват")
	}()
	fmt.Println("Съобщение 1")
	amapanic() // HL
	fmt.Println("Съобщение 2")
}

func main() {
	causepanic()
}

// END OMIT
