package main

import "fmt"

func main() {
	s := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	s2 := s[2:4]
	s2[0] = 'x'

	fmt.Println(s[2] == 'x')
}
