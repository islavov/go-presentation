package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "дж"
	fmt.Println(utf8.RuneCountInString(s)) // 2
}
