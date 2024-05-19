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
