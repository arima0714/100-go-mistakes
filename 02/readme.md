# No.2: 不必要にネストしたコード

条件分岐が複雑だったとしても、簡約できるかもしれない。if文のネストを２つまでにしようと心がけると、１つ目だけを把握することで流れを確認出来て、２つ目まで確認するとエッジケースを確認できるようなコードに仕上げやすい。

if ブロックがリターンする下記のようなコードの場合は

```
package main

func main() {
	var cond bool
	if cond {
		println("cond is true")
		return
	} else {
		println("cond is else")
	}
}
```

下記のように else を省略すべき

```
package main

func main() {
	var cond bool
	if cond {
		println("cond is true")
		return
	}
	if !cond {
		println("cond is else")
		return
	}
}

```

ネストを減らし、できるだけ早くリターンすることは、コードの可読性を向上させる良い方法。
