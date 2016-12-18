package main

import "fmt"

func main() {
	defer func() { // HL
		fmt.Println("Мечката си отива")
	}() // HL

	fmt.Println("Мечката идва")
}
