package main

import (
	"fmt"
	"math"
)

func addInt(a, b int) int {
	// 整数でオーバーフローが発生するかを検査
	if (b > 0 && a > math.MaxInt-b) || (b < 0 && a < math.MinInt-b) {
		panic("int overflow")
	}
	return a + b
}

func main() {
	var base int = math.MaxInt - 100

	// オーバーフロー発生しない
	num := addInt(base, 99)
	fmt.Printf("num = %d\n", num)

	// オーバーフロー発生しない２
	num = addInt(base, 100)
	fmt.Printf("num = %d\n", num)

	// オーバーフロー発生する
	num = addInt(base, 200)
	fmt.Printf("num = %d\n", num)
}
