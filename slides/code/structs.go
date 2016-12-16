package main

import (
	"fmt"
)

// START OMIT

type SantaReindeer struct {
	Name string
}

func main() {
	blitzen := SantaReindeer{Name: "Blitzen"}
	vixen := SantaReindeer{"Vixen"}
	anonyxen := SantaReindeer{}
	fmt.Printf("%+v\n", vixen)
	fmt.Printf("%+v\n", blitzen)
	fmt.Printf("%+v\n", anonyxen)
}

// END OMIT
