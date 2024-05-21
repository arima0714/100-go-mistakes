// 著作権に関するあれこれなど

// Package main はメイン関数
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
