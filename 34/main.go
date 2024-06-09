package main

import "fmt"

func main() {
loop: // loop ラベルの定義
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
		switch i {
		default:
		case 2:
			break loop
		}
	}
}
