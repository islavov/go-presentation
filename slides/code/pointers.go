package main

import "fmt"

func main() {
	var (
		x int  = 10
		p *int = &x
	)
	fmt.Println(p)
	*p = 2
	fmt.Println(x) // 2
}
