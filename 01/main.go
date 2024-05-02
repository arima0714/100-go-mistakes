package main

func retInput(i int) int {
	return i
}

func main() {
	var target int
	if true {
		temp := retInput(1)
		target = temp
	} else {
		temp := retInput(2)
		target = temp
	}
	println(target)
}
