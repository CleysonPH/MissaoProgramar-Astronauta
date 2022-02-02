package main

import "fmt"

func main() {
	var num1, num2 float64
	fmt.Scanf("%f/%f", &num1, &num2)
	result := num1 / num2
	fmt.Printf("%.2f\n", result)
}
