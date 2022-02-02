package main

import "fmt"

func main() {
	var a, b float64
	fmt.Print("Insira o primeiro número: ")
	fmt.Scanf("%f", &a)
	fmt.Print("Insira o segundo número: ")
	fmt.Scanf("%f", &b)
	result := a * b
	fmt.Printf("%.2f X %.2f = %.2f\n", a, b, result)
}
