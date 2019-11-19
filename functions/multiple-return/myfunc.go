package main

import "fmt"

// Simple multiple return
func mul2and3(x int) (int, int) {
	return x * 2, x * 3
}

// Named multiple return
func namedReverse(x int) (err string, value float64) {
	if x == 0 {
		err = "Can not divide by 0"
	} else {
		value = 1.0 / float64(x)
	}
	return
}

func main() {
	a, b := mul2and3(2)
	fmt.Println(a, b)
	e0, i0 := namedReverse(0)
	fmt.Println(e0, i0)
	e1, i1 := namedReverse(2)
	fmt.Println(e1, i1)
}
