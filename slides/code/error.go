package main

import (
	"io/ioutil"
	"fmt"
)


func checkerrors(filename string) error {
	_, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to open %s", err)
	}
	return nil
}

func main() {
	checkerrors("f1")
}
