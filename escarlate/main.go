package main

import "fmt"

func main() {
	var p1, p2, p3 float64
	fmt.Print("Diga o preço do primeiro produto: ")
	fmt.Scanf("%f", &p1)
	fmt.Print("Diga o preço do segundo produto: ")
	fmt.Scanf("%f", &p2)
	fmt.Print("Diga o preço do terceiro produto: ")
	fmt.Scanf("%f", &p3)

	if p1 < p2 && p1 < p3 {
		fmt.Println("Compre o primeiro produto")
	} else if p2 < p1 && p2 < p3 {
		fmt.Println("Compre o segundo produto")
	} else {
		fmt.Println("Compre o terceiro produto")
	}
}
