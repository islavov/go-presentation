package main

import (
	"fmt"
)

// START OMIT

type SantaReindeer struct {
	Name   string
	energy int
}

func (sr SantaReindeer) Bawl() string { // HL
	return "Baa" // HL
} // HL

func (sr *SantaReindeer) Feed(energy int) { // HL
	sr.energy += energy // HL
} // HL

func main() {
	blitzen := SantaReindeer{Name: "Blitzen", energy: 20}
	blitzen.Feed(10)
	fmt.Printf("%s says '%+v'!\n", blitzen.Name, blitzen.Bawl())
}

// END OMIT
