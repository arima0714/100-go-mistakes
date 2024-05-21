# No.15: コードのドキュメントがない

公開されている要素はすべて文書化されなければならない。構造体、インターフェース、関数のどれであっても。

慣習といて公開される要素の名前で始まるコメントを追加する。

```
package main

// Customer は顧客を表す
type Customer struct{}

// ID は顧客 ID を返す
//
// Deprecated: この関数は、非推奨です。代わりに GetID を使ってください
func (c Customer) ID() string { return "ID: " }

// GetID は顧客 ID を返す
func (c Customer) GetID() string { return "" }

func main() {
	cus := Customer{}
	print(cus.GetID)
}

```

例は上記の通り。また、パッケージにはパッケージドキュメントをつけよう。パッケージドキュメントはパッケージと同じ名前の関連するファイルか、`doc.go` のような特定のファイルに書くべき。パッケージドキュメントの記入例は下記の通り。パッケージ宣言に隣接していないコメントは省略される。したがって `著作権~~~` はパッケージドキュメントとして表示されない。

```
// 著作権に関するあれこれなど

// Package main はメイン関数
package main
```