package main

import "fmt"

func main() {
	var x string = "a"

	func() {
		x, y := "b", "c"
		fmt.Println(x, y)
	}()

	fmt.Println(x)
	x, y := "x", "y"
	fmt.Println(x, y)
}
