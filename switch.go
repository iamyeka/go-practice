package main

import (
	"fmt"
)

const (
	B float64 = 1 << (10 * iota)
	KB
)

func main() {
	//fmt.Println(B)
	//fmt.Println(KB)

	switch a := 1; {
	case a >= 0:
		fmt.Println("a >= 0")
		fallthrough
	case a >= 1:
		fmt.Println("a >= 1")
	default:
		fmt.Println("a <= 0")
	}
}
