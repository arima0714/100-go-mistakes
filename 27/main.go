package main

import (
	"fmt"
	"time"
)

const length = 10000000

// 適切にマップを初期化する関数
func goodPractice() map[string]int {
	m := make(map[string]int, length)
	for i := 0; i < length; i++ {
		m[fmt.Sprint(i)] = i
	}
	return m
}

// 不適切にマップを初期化する関数
func badPractice() map[string]int {
	m := make(map[string]int)
	for i := 0; i < length; i++ {
		m[fmt.Sprint(i)] = i
	}
	return m
}

func main() {
	// 1000000 の要素を持ったマップを２つ作成する。

	// goodPractice(), badPractice() でマップを作成し、その時間を計測する
	time1 := time.Now()
	goodPractice()
	time2 := time.Now()
	// time2 から time1 を引いて、その時間を出力する
	fmt.Println(time2.Sub(time1))

	time3 := time.Now()
	badPractice()
	time4 := time.Now()
	// time4 から time3 を引いて、その時間を出力する
	fmt.Println(time4.Sub(time3))

}
