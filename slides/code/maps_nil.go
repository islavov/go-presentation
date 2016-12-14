package main

import "fmt"

func main() {
	var basket map[string]int
	fmt.Println(len(basket))       // 0
	fmt.Println(basket == nil)     // true
	fmt.Println(basket["oranges"]) // 0
	basket["oranges"] = 3          // panic
}
