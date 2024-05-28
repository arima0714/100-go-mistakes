package main

import "fmt"

func main() {
	// 長さと容量が5のスライスを作成
	s := make([]int, 5)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	// スライスに新たな要素を追加
	s = append(s, 1)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
