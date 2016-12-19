package main

import "fmt"

// START OMIT
func checkChristmasTree() {
	panic("Елхата гори. Това не трябва да се случва") // HL
}

func findPresents() {
	defer func() {
		fmt.Println("Излизаш от стаята")
	}()
	fmt.Println("Влизаш в стаята")
	checkChristmasTree() // HL
	fmt.Println("Отваряш подаръците")
}

func main() {
	findPresents()
}

// END OMIT
