package main

import "fmt"

func main() {
	var n1, n2, n3 float64
	fmt.Print("Digite a primeira nota: ")
	fmt.Scanf("%f", &n1)
	fmt.Print("Digite a segunda nota: ")
	fmt.Scanf("%f", &n2)
	fmt.Print("Digite a terceira nota: ")
	fmt.Scanf("%f", &n3)

	avg := (n1 + n2 + n3) / 3.0
	fmt.Printf("Média: %.2f\n", avg)
	if avg < 7 {
		fmt.Println("Reprovado")
	} else if avg < 10 {
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Aprovado com Distinção")
	}
}
