package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fc, err := ioutil.ReadFile("a/b/c/d.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("File contents: %s", fc)
}
