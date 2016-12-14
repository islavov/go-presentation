package main

import "fmt"

func main() {
	basket := map[string]int{"apples": 1, "oranges": 2}
	_, have_peaches := basket["peaches"]
	fmt.Println(have_peaches)
}
