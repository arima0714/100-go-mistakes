package main

import "fmt"

func copyMap(m map[int]bool) map[int]bool {
	m2 := make(map[int]bool)
	for k, v := range m {
		m2[k] = v
	}
	return m2
}

func main() {
	for i := 0; i < 3; i++ {
		m := map[int]bool{1: true, 2: false, 3: true}
		m2 := copyMap(m) // 初期マップのコピーを作成
		for k, v := range m {
			if v {
				m2[10+k] = v
			}
		}
		fmt.Println(m2)
	}
}
