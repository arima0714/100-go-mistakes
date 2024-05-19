# No.10: 型埋め込みで起こりえる問題点を意識していない

```
package main

import "fmt"

type Foo struct {
	Bar // 埋め込みフィールド
}

type Bar struct {
	Baz int
}

func main() {
	foo := Foo{}
	fmt.Printf("foo = %d\n", foo)
	fmt.Printf("foo.Bar = %d\n", foo.Bar)
	fmt.Printf("foo.Baz = %d\n", foo.Baz)
}

```

上記プログラムの出力は下記の通り

```
foo = {{0}}
foo.Bar = {0}
foo.Baz = 0
```

`foo.Baz` で値を参照できている。これが、意図しないフィールドを参照できてしまってしまうことがある。これを型埋め込みと呼ぶ。

基本的に、型埋め込みは使わなくても何とかなるが、便利だから使ってしまう。

下記のルールを守るようにすると、お決まりの退屈なコードを書かなくて済む。
* フィールドへのアクセスを簡単にする記法だとして使うなら、素直にフィールドを使いなさい
* 外部から隠蔽したいフィールドやメソッドを型埋め込みで利用できるようにしない
