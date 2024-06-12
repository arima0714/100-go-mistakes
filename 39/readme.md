# No.39: 最適化されていない文字列の連結

下記は性能の悪い文字列の結合の実例。

```golang
package main

func concat(values []string) string {
	s := ""
	for _, value := range values {
		// 一見良さそうだが、反復ごとに s が再生成されるため、性能は悪い
		s += value
	}
	return s
}

func main() {
	values := []string{"Hello", "World", "!!!"}
	println(concat(values))
}
```

strings.Builder を利用して性能向上させたものが下記

```golang
package main

import "strings"

func concat(values []string) string {
	sb := strings.Builder{} // strings.Builder の利用
	for _, v := range values {
		_, _ = sb.WriteString(v)
	}
	return sb.String()
}

func main() {
	values := []string{"Hello", "World", "!!!"}
	println(concat(values))
}

```

strings.Builder では内部でスライスを使っている。スライスの長さが固定長であれば指定する方が効率的。従って、スライスの長さ指定を行い、更なる効率化を果たしたコードが下記。

```golang
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

```

* 5個以上の文字列を連結するときから、strings.Builderを使った方が性能が良くなる
* 最終的な文字列のバイト数が既知の場合、Grow() を使って内部のバイトスライスを事前に割り当てよう
