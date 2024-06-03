package main

import "fmt"

func main() {
	src := []int{0, 1, 2}
	dst := make([]int, len(src))
	copy(dst, src)
	fmt.Println(dst) // 出力は [0, 1, 2]
}
