package main

import "fmt"

var unit int = 9

func main() {
	fmt.Print(sum())
}

func info(O int, N int) int {
	used := N - O
	return used
}

func sum() int {
	elec := info(20, 30)
	bill := unit * elec
	return bill
}
