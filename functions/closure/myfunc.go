package main

import "fmt"

/*
- Go only allows named function inside function
- Inner function can get and change variable from outer function
*/
func nextInt(start uint) func() uint {
	initValue := start
	return func() uint {
		initValue++
		return initValue
	}
}

func main() {
	nextNumber := nextInt(3)
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
}
