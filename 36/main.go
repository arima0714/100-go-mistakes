package main

import "fmt"

func main() {
	s := "Hello"
	fmt.Println(len(s))

	s = "漢"
	fmt.Println(len(s))

	s = string([]byte{0xE6, 0xBC, 0xA2})
	fmt.Println(s)
}
