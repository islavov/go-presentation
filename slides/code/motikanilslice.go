package main

import "fmt"

func main() {
	var m map[string]string
	fmt.Println(m["foo"]) // Empty string
	m["foo"] = "bar"      // Panic
}
