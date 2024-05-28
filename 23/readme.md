# No.23: スライスが空か否かを適切に検査しない

まずは失敗例

```golang
package main

func handleOperation(id string) {
	operations := getOperations(id)
	if operations != nil { // 実装ミスがあり getOperations() の返り値が何であれ、operations != nil が true になる
		print("do something\n")
	}
}

func getOperations(id string) []float32 {
	operations := make([]float32, 0) // operations スライスを初期化する
	if id == "" {
		return operations // 渡された id が空文字列の場合は operations スライスを返す
	}

    operations = append(operations, 123)

	return operations
}

func main() {
	handleOperation("") // do something が出力される

	handleOperation("123") // do something が出力される
}

```

スライスが空か nil かは、長さを検査すると良い。
* スライスが nil なら、`len(operations) != 0` は false
* スライスが空なら、`len(operations) != 0` は false

呼び出す関数を常に修正できるとは限らないので、長さを検査しよう。

nil スライスを返しても空スライスを返しても挙動が変わらないような実装をしよう。

また、マップが空かどうかを検査するのにも長さを検査するようにすると良いです。

