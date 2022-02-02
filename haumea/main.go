package main

import "fmt"

func main() {
	var num int
	fmt.Print("Insira um número decimal: ")
	fmt.Scanf("%d", &num)
	fmt.Printf("O número octal é: %o\n", num)
}
