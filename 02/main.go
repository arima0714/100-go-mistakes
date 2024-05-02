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
