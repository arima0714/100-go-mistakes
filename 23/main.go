package main

func handleOperation(id string) {
	operations := getOperations(id)
	if len(operations) != 0 {
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
