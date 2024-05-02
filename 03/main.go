package main

var a = func() int {
	println("var")
	return 0
}()

func init() {
	println("init")
}

func main() {
	println("main")
}
