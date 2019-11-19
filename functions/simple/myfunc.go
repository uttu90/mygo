package main

import "fmt"

// Simple function to convert Farenhai to Celcius
func f2c(f float64) float64 {
	return (f - 32) * (5 / 9)
}

func main() {
	fmt.Println(f2c(32))
	fmt.Println(f2c(232))
}
