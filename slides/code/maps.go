package main

import "fmt"

func main() {
	basket := map[string]int{"apples": 1, "oranges": 2}
	fmt.Println(basket["apples"])
	basket["apples"]++
	fmt.Println(basket["apples"])
}
