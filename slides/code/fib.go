package main

import "fmt"

func fibbgen() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a + b
		return a
	}
}

func main() {
	fib := fibbgen()
	fmt.Println(fib(), fib(), fib(), fib(), fib(), fib(), fib())
}
