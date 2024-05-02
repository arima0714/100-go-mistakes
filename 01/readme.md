# No1: 意図しない変数シャドウイング

下記の出力は 0 になる。変数 `target` のスコープが違うので、変更したはずが全く変更されておらず、初期化状態のままになっている。git 

```
package main

func retInput(i int) int {
	return i
}

func main() {
	var target int
	if true {
		target := retInput(1)
		target = retInput(target)
	} else {
		target := retInput(2)
		target = retInput(target)
	}
	println(target)
}

```

一時変数をかませても良いね。

```
package main

func retInput(i int) int {
	return i
}

func main() {
	var target int
	if true {
		temp := retInput(1)
		target = temp
	} else {
		temp := retInput(2)
		target = temp
	}
	println(target)
}
```

