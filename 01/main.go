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
