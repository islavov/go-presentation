package main

import (
	"fmt"
)

type SantaReindeer struct{ Name string }

type Rudolph struct {
	SantaReindeer
	NoseLit bool
}

// START OMIT

func Greet(deer SantaReindeer) {
	fmt.Printf("Hello, %s!", deer.Name)
}

func main() {
	rudolph := Rudolph{SantaReindeer: SantaReindeer{Name: "Rudolph"}, NoseLit: true}
	Greet(rudolph)
}

// END OMIT
