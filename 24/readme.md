# No.24: スライスのコピーを正しく行わない

```
package main

import "fmt"

func main() {
	src := []int{0, 1, 2}
	var dst []int
	copy(dst, src)
	fmt.Println(dst) // 出力は []
}

```

上記のコードでは `[0, 1, 2]` が出力されるのを期待するが、実際には `[]` である。`copy()` ではコピー先のスライスにコピーされる要素の数が、コピー元のスライスの長さ・コピー先のスライスの長さの小さい方である。

上記の例では
* コピー元のスライス（`src`）の長さ：3
* コピー先のスライス（`dst`）の長さ：0
より、0個の要素がコピーされる。完全なコピーを行うにはコピー先のスライス（`dst`）の長さがコピー元のスライス（`src`）のより長い必要がある。したがって下記のようになる。

```
package main

import "fmt"

func main() {
	src := []int{0, 1, 2}
	var dst []int
	copy(dst, src)
	fmt.Println(dst) // 出力は []
}

```