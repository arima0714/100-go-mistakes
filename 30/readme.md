# No.30: range ループで要素がコピーされることを無視する

下記のプログラムを参考にループに要素がコピーされてしまうこと。ループでの要素の更新方法を学べる。

```golang
package main

import "fmt"

type account struct {
	balance float32
}

func main() {
	accounts := []account{
		{balance: 100.},
		{balance: 200.},
		{balance: 300.},
	}
	for _, a := range accounts {
		a.balance += 1000 // a はコピーなので元の配列には影響しない
	}
	fmt.Println(accounts) // [{100} {200} {300}]

	for i := range accounts { // 値を更新する方法１
		accounts[i].balance += 1000
	}
	fmt.Println(accounts) //[{1100} {1200} {1300}]

	for i := 0; i < len(accounts); i++ { // 値を更新する方法２
		accounts[i].balance += 1000
	}
	fmt.Println(accounts) // [{2100} {2200} {2300}]

	accounts2 := []*account{ // そもそもポインタへのスライスを使う
		{balance: 100.},
		{balance: 200.},
		{balance: 300.},
	}
	for _, a := range accounts2 {
		a.balance += 1000
	}
	fmt.Print("[")
	for _, a := range accounts2 {
		fmt.Printf("{%v} ", a.balance)
	}
	fmt.Print("]\n")
	// [{1100} {1200} {1300}]
}
```

raneg ループ内の値要素はコピーであるということを忘れない。要素にアクセスするときはインデックス経由でやろう。
