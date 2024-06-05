package main

import (
	"fmt"
	"runtime"
)

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB\n", m.Alloc/1024/1024)
}

func randBytes() *[128]byte {
	var b [128]byte
	for i := 0; i < 128; i++ {
		b[i] = byte(i)
	}
	return &b
}

func main() {
	n := 1_000_000
	m := make(map[int]*[128]byte)
	printAlloc()

	for i := 0; i < n; i++ { // <- 100万個の要素を追加
		m[i] = randBytes()
	}
	printAlloc()

	for i := 0; i < n; i++ { // <- 100万個の要素を削除
		m[i] = randBytes()
	}

	for i := 0; i < n; i++ {
		delete(m, i)
	}

	runtime.GC()
	printAlloc()
	runtime.KeepAlive(m)
}
