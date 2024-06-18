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
