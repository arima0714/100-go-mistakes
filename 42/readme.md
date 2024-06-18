# No.42: どちらのレシーバ型を使うべきかわかっていない

* 値またはポインタのレシーバを使うことは、性能以外の条件で決定されるため

基本のコードは下記の通り。ポインタで渡されたものはいじったら反映される。値で渡されたものはいじっても反映されない。

```golang
package main

import "fmt"

type customer struct {
	balance float64
}

func (c customer) addv(v float64) {
	c.balance += v
}

func (c *customer) addp(v float64) {
	c.balance += v
}

func main() {
	c := customer{balance: 100.0}
	c.addv(50.)
	fmt.Printf("balance: %.2f\n", c.balance) // 100.00

	c.addp(50.)
	fmt.Printf("balance: %.2f\n", c.balance) // 150.00
}

```

* レシーバがポインタでなければならない条件
    * メソッドがレシーバを変更する必要がある
    * メソッドのレシーバがコピーできないフィールドを含む
* レシーバがポインタであるべき条件
    * レシーバが大きなオブジェクトである
* レシーバが値でなければならない条件
    * レシーバの不変性を矯正する
    * レシーバがマップ、関数、チャネル
* レシーバが値であるべき条件
    * レシーバが変更する必要のないスライス
    * レシーバが小さな配列や可変なフィールドを持たず値型である

下記の例のように値レシーバがポインタフィールドから参照される構造体を持つ時は値レシーバで渡されたも内部の値をいじれる

```golang
package main

import "fmt"

type customer struct {
	data *data
}

type data struct {
	balance float64
}

func (c customer) add(operation float64) {
	c.data.balance += operation
}

func main() {
	c := customer{data: &data{balance: 100}}
	c.add(50.)
	fmt.Printf("balance: %f\n", c.data.balance)
}

```

このように値レシーバでも工夫すると中身をいじれるようになる。しかし、わかりやすさのためには、ポインタレシーバを使う方が良いかもしれない。
