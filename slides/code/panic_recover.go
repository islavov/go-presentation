package main

import "fmt"

// START OMIT
func openbox() {
	fmt.Println("А wild spider appears")
	panic("Start panicking") // HL
}

func main() {
	defer func() {
		if r := recover(); r != nil { // HL
			fmt.Println(r)
			fmt.Println("Used a slipper. It was very effective.")
		}
	}()
	openbox() // HL
}

// END OMIT
