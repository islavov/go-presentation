package main

import "fmt"

type entry struct {
	value int
}

func main() {
	a := []entry{entry{value: 1}, entry{value: 2}}
	for _, e := range a {
		e.value *= 10
	}
	for _, e := range a {
		fmt.Println(e.value) // 1 and 2
	}
}
