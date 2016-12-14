package main

import "fmt"

func main() {
	var s string = "a"
	var x int = 1
	s = "b"
	fmt.Println(s)
	x, s = 1, "c"
	fmt.Println(x, s)
	x += 1
	x++
	fmt.Println(x)
}
