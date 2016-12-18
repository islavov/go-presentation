package main

import "fmt"

func setFirst(a [3]string, v string) { // HL
	a[0] = v // HL
}

func main() {
	a := [3]string{"one", "two", "three"}
	setFirst(a, "xxx") // HL
	fmt.Println(a[0])  // "one"
}
