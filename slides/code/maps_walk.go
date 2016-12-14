package main

import "fmt"

func main() {
	basket := map[string]int{"apples": 1, "oranges": 2}
	for k := range basket {
		fmt.Println(k, basket[k])
	}
}
