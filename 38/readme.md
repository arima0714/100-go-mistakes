# No.38: Trim 関数の誤用

TrimXXXX() のよくあるミスを含んだコードは下記の通り

```golang
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.TrimRight("123oxo", "xo")) // 123

	fmt.Println(strings.TrimSuffix("123oxo", "xo")) // 123o

	fmt.Println(strings.TrimLeft("oxo123", "ox")) // 123
	
	fmt.Println(strings.TrimPrefix("oxo123", "ox")) // o123

	fmt.Println(strings.Trim("oxo123oxo", "ox")) // 123
}
```

> * TrimRight/ TrimLeft は、集合内の末尾/先頭のルーンを削除します。
> * TrimSuffix/ TrimPrefix は、指定された接尾辞/接頭辞を削除します。
