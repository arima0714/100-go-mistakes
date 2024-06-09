# No.34: break 文の仕組みを無視する

* break は一番内側のfor文、switch文、select文の実行を終了させる

下記はよくある過ちの例

```golang
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)

		switch i {
		default:
		case 2:
			break // i が 2 の時にループを抜けるつもりで break を使っているが、実際には switch 文を抜けてしまう。したがって、for 文は終了しない
            // 出力は 0 1 2 3 4 である
		}
	}
}
```

上記を期待通りに動かすために修正したコードは下記の通り

```
package main

import "fmt"

func main() {
loop: // loop ラベルの定義
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
		switch i {
		default:
		case 2:
			break loop
		}
	}
}
```

ラベルを使ったfor文からのブレイクはGoの慣用的な方法

break を使う場合、どの文に影響を与えるかを常に確認すべき
