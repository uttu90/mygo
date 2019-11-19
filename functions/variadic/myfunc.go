package main

import "fmt"

// Variadic function means function without fixed parameters length
func add(agrs ...int) int {
	total := 0
	for _, val := range agrs {
		total += val
	}
	return total
}

func main() {
	fmt.Println(add(1, 2, 3))
	xs := []int{1, 2, 3}
	fmt.Println(add(xs...))
}
