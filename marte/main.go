package main

import "fmt"

func main() {
	var (
		a float64 = -5 + 8*6
		b float64 = (55 + 9) % 9
		c float64 = 20 + -3*5.0/8.0
		d float64 = 5 + 15/3*2 - 8%3
	)
	fmt.Printf("a = %.2f\n", a)
	fmt.Printf("b = %.2f\n", b)
	fmt.Printf("c = %.2f\n", c)
	fmt.Printf("d = %.2f\n", d)
}
