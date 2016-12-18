package main

import "fmt"

func main() {
	var x string = "a" // HL

	func() {
		x, y := "b", "c" // HL
		fmt.Println(x, y)
	}()

	fmt.Println(x) // "a", not "b"
}
