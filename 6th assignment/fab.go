package main

import "fmt"

func printFibonacciSeries(num int) {
	x := 0
	y := 1
	z := y
	fmt.Printf("Series is: %d %d", x, y)
	for true {
		z = y
		y = x + y
		if y >= num {
			fmt.Println()
			break
		}
		x = z
		fmt.Printf(" %d", y)
	}
}

func main() {
	printFibonacciSeries(10)
	printFibonacciSeries(16)
	printFibonacciSeries(100)
}
