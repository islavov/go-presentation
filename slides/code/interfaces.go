package main

import (
	"fmt"
)

// START OMIT
type SantaReindeer struct{ Name string }

func (sr SantaReindeer) String() string { // HL
	return "A deer named " + sr.Name // HL
} // HL

func main() {
	dancer := SantaReindeer{Name: "Dancer"}
	fmt.Printf("%s", dancer) // HL
}

// END OMIT
