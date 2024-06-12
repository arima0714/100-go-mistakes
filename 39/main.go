package main

import "strings"

func concat(values []string) string {

	// 合計バイト数を計算
	total := 0
	for i := 0; i < len(values); i++ {
		total += len(values[i])
	}

	sb := strings.Builder{}
	sb.Grow(total) // バイト数を指定してスライスの大きさを決定
	for _, value := range values {
		_, _ = sb.WriteString(value)
	}
	return sb.String()
}

func main() {
	values := []string{"Hello", "World", "!!!"}
	println(concat(values))
}
