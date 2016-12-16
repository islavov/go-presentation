package main

import (
	"fmt"
)

// START OMIT

type SantaReindeer struct{ Name string }

func (sr SantaReindeer) Bawl() string { return "Baa" }

type Rudolph struct {
	SantaReindeer
	NoseLit bool
}

func main() {
	rudolph := Rudolph{SantaReindeer: SantaReindeer{Name: "Rudolph"}, NoseLit: true}
	rudolphPtr := &rudolph
	rudolphPtr.NoseLit = false
	fmt.Printf("%s says '%s' (nose lit: %t)\n",
		rudolphPtr.Name,
		rudolphPtr.Bawl(), // HL
		rudolphPtr.NoseLit)
}

// END OMIT
