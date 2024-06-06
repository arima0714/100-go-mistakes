# No.33: マップの反復処理で誤った仮定をする

マップの重要基本事項
* データをキーでソートして保持していない
* データが追加されたときの順序を保持していない

Go 開発者は、マップを反復処理するときに順序に依存しない実装にする必要がある。

ある反復処理中にマップに要素が追加される場合、次の反復でその要素が取り出されるかもしれないし、取り出されないかもしれない。

上記を表すためのプログラムと出力は下記の通り

```golang
package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		m := map[int]bool{1: true, 2: false, 3: true}
		for k, v := range m {
			if v {
				m[10+k] = v
			}
		}
		fmt.Println(m)
	}
}

```

```bash
$ go run ./33/main.go 
map[1:true 2:false 3:true 11:true 13:true 23:true 33:true 43:true 53:true 63:true]
map[1:true 2:false 3:true 11:true 13:true 21:true]
map[1:true 2:false 3:true 11:true 13:true]
```

コードを予測可能な出力にするには、マップのコピーを作成して、そのコピーを更新する

上記を実現したコードは下記の通り

```golang
package main

import "fmt"

func copyMap(m map[int]bool) map[int]bool {
	m2 := make(map[int]bool)
	for k, v := range m {
		m2[k] = v
	}
	return m2
}

func main() {
	for i := 0; i < 3; i++ {
		m := map[int]bool{1: true, 2: false, 3: true}
		m2 := copyMap(m) // 初期マップのコピーを作成
		for k, v := range m {
			if v {
				m2[10+k] = v
			}
		}
		fmt.Println(m2)
	}
}

```

出力は下記

```bash
map[1:true 2:false 3:true 11:true 13:true]
map[1:true 2:false 3:true 11:true 13:true]
map[1:true 2:false 3:true 11:true 13:true]

```

最後に、マップを使った実装は下記に依存するな
* キーで順序付けされたデータ
* 挿入順序の保持
* 決定的な反復順序
* 反復処理中に追加された要素の生成
