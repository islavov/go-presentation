package main

import "fmt"

func amapanic() {
	panic("Това не трябва да се случва")
}

func causepanic() {
	defer func() {
		fmt.Println("Отложените функции се изпълняват")
	}()
	fmt.Println("Съобщение 1")
	amapanic()
	fmt.Println("Съобщение 2")
}

func main() {
	causepanic()
}
