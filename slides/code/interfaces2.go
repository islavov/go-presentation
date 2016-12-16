package main

import (
	"fmt"
)

// START OMIT

func checkType(value interface{}) {
	_, ok := value.(string)
	if ok {
		fmt.Printf("%s is a string\n", value)
	} else {
		fmt.Printf("%v is not a string\n", value)
	}
}

func main() {
	var value interface{}
	value = 3
	checkType(value)
	value = "three"
	checkType(value)
}

// END OMIT
