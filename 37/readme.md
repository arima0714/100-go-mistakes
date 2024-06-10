# No.37: 不正確な文字列の反復

不正確な文字列の反復および解決策とその出力を実践したコードは下記の通り

```golang
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
```

ルーンを反復処理したいとき、文字列に対して直接 range ループを使える。

しかし、インデックスがルーンのバイト列の開始インデックスに対応しているため、一般的なアルファベット以外が入った文字列を range ループすると想定外の事態になる。

ルーンそのものにアクセスしたい場合、文字列のインデックスではなく、range の値変数を使うべき。

i番目のルーンを取得したいならば、文字列をルーンのスライスに変換すべき
