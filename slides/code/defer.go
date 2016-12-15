package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Мечката си отива")
	}()

	fmt.Println("Мечката идва")
}
