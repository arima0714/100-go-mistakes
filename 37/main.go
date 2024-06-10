package main

import "fmt"

func main() {
	s := "hêllo"
	for i := range s {
		fmt.Printf("position %d: %c\n", i, s[i])
	}
	fmt.Printf("len=%d\n", len(s))
	// 出力は下記の通り
	// position 0: h
	// position 1: Ã
	// position 3: l
	// position 4: l
	// position 5: o
	// len=6

	s = "hêllo"
	for i, r := range s {
		fmt.Printf("position %d: %c\n", i, r)
	}
	// 出力は下記の通り
	// position 0: h
	// position 1: ê
	// position 3: l
	// position 4: l
	// position 5: o

	s = "hêllo"
	runes := []rune(s)
	for i, r := range runes {
		fmt.Printf("position %d: %c\n", i, r)
	}
	// 	出力は下記の通り
	// position 0: h
	// position 1: ê
	// position 2: l
	// position 3: l
	// position 4: o

	s = "hêllo"
	r := []rune(s)[4]
	fmt.Printf("%c\n", r) // o
}
