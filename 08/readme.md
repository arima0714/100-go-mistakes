# No.8: any は何も伝えない

```golang

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


```

上記のように any を使うとなんでも入れられるが、変数 i から有用な値を得るには型アサーションが必要になる。