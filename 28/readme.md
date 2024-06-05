# No.28: 

```golang
package main

import (
	"fmt"
	"runtime"
)

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB\n", m.Alloc/1024/1024)
}

func randBytes() [128]byte {
	var b [128]byte
	for i := 0; i < 128; i++ {
		b[i] = byte(i)
	}
	return b
}

func main() {
	n := 1_000_000
	m := make(map[int][128]byte)
	printAlloc()

	for i := 0; i < n; i++ { // <- 100万個の要素を追加
		m[i] = randBytes()
	}
	printAlloc()

	for i := 0; i < n; i++ { // <- 100万個の要素を削除
		m[i] = randBytes()
	}

	for i := 0; i < n; i++ {
		delete(m, i)
	}

	runtime.GC()
	printAlloc()
	runtime.KeepAlive(m)
}

```

実行結果は下記の通り

```bash
$ go run ./28/main.go 
Alloc = 0 MiB
Alloc = 453 MiB
Alloc = 293 MiB
```

上記の結果から、ヒープは縮小したが、それほどではない。その理由は、マップから要素を取り除いても、既存のバケットの数には影響を与えず、縮小はしないから。

マップが消費するメモリ量を自動で減らすためには
* 現在のマップのコピーを定期的に再作成すること
* map 型が配列へのポインタを保持するように変更する、つまり、map[int]*[128]byte とすること。

後者を適用した先ほどと同様の例は下記の通り

```golang
package main

import (
	"fmt"
	"runtime"
)

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB\n", m.Alloc/1024/1024)
}

func randBytes() *[128]byte {
	var b [128]byte
	for i := 0; i < 128; i++ {
		b[i] = byte(i)
	}
	return &b
}

func main() {
	n := 1_000_000
	m := make(map[int]*[128]byte)
	printAlloc()

	for i := 0; i < n; i++ { // <- 100万個の要素を追加
		m[i] = randBytes()
	}
	printAlloc()

	for i := 0; i < n; i++ { // <- 100万個の要素を削除
		m[i] = randBytes()
	}

	for i := 0; i < n; i++ {
		delete(m, i)
	}

	runtime.GC()
	printAlloc()
	runtime.KeepAlive(m)
}
```

実行結果は下記の通り。

```bash
$ go run ./28/main.go 
Alloc = 0 MiB
Alloc = 181 MiB
Alloc = 38 MiB
```
