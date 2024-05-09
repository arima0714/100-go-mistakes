package main

func main() {
	var i any
	i = 43 // int
	i = "foo" // string
	i = struct { // 構造体
		a string
	}{
		a: "bar",
	}
	i = f // 関数

	_ = i // コンパイルエラー対策
}

func f() {}
