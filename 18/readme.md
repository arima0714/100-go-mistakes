# No.18: 整数のオーバーフローを無視する

```golang
package main

import "math"

func main() {
	var counter int32 = math.MaxInt32 - 1
	counter += 2

	print(counter)
}


```

上記の例のような場合は簡単にオーバーフローの存在に気付けるが普段はそうではない。そして、オーバーフローが発生してもパニックになるということもない。これに気をつけよう

1加算時のオーバーフロー対策は加算前に加算対象の値が `math.MaxInt`, `math.MinInt`, `math.MaxUint` であるかを事前に確認することで防ぐことができる。これは uint 型でも同様。

加算時のオーバーフローの検出は下記の例を参考に行う。

```golang
package main

import (
	"fmt"
	"math"
)

func addInt(a, b int) int {
	// 整数でオーバーフローが発生するかを検査
	if (b > 0 && a > math.MaxInt-b) || (b < 0 && a < math.MinInt-b) {
		panic("int overflow")
	}
	return a + b
}

func main() {
	var base int = math.MaxInt - 100

	// オーバーフロー発生しない
	num := addInt(base, 99)
	fmt.Printf("num = %d\n", num)

	// オーバーフロー発生しない２
	num = addInt(base, 100)
	fmt.Printf("num = %d\n", num)

	// オーバーフロー発生する
	num = addInt(base, 200)
	fmt.Printf("num = %d\n", num)
}

```

上記プログラムの出力は下記のようになる。

```
% go run ./main.go 
num = 9223372036854775806
num = 9223372036854775807
panic: int overflow

goroutine 1 [running]:
main.addInt(...)
        /home/user/docs/100-go-mistakes/18/main.go:11
main.main()
        /home/user/docs/100-go-mistakes/18/main.go:25 +0xce
exit status 2

```
