package main

import "fmt"

type customer struct {
	id string
	// operatoins []float64
}

func main() {
	cust1 := customer{id: "x"}
	cust2 := customer{id: "x"}
	fmt.Println(cust1 == cust2) // true
}
