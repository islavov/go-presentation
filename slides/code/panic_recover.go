package main

import "fmt"

func openbox() {
	fmt.Println("А wild spider appears")
	panic("Start panicking")
}

func main() {
	defer func() {
		if r := recover();r != nil {
			fmt.Println(r)
			fmt.Println("Used a slipper. It was very effective.")
		}
	}()
	openbox()
}
