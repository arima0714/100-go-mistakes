# No.43: 名前付き結果パラメータを使わない

名前付き結果パラメータの使用例と使った方が良い例は下記の通り

```golang
package main

// 名前付きパラメータの使用例
func f(a int) (b int) { // int の結果変数に b と名づける
	b = a
	return // b の現在値を返す
}

type locator interface {
	// getCoordinates(address string) (float32, float32, error)
	// 上記は緯度と軽度のどちらがどちらかわからない
	getCoordinates(address string) (latitude float32, longitude float32, err error)
	// 緯度と軽度の順番がわかる
}
```

エラーを返す時に err という名前では意味はない。いつでも名前付き結果パラメータを使えば良いというわけではないのだ。

* インターフェース定義で名前付き結果パラメータを使うと可読性を高められる
* 2つの結果パラメータが同じ型を持っている場合や利便性を高めるために、メソッドでも名前付き結果パラメータを使うと良い
