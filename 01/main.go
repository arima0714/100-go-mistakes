package main

func retInput(i int) int {
	return i
}

func main() {
	var target int
	if true {
		target = retInput(1)
	} else {
		target = retInput(2)
	}
	println(target)
}
