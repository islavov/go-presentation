package main

import "fmt"

func main() {
	var a [5]int
	a[0] = 1
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	b := [3]int{1, 2, 3}
	fmt.Println(b[0])
}
