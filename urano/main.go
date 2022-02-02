package main

import "fmt"

func main() {
	var num int
	fmt.Print("Insira um número: ")
	fmt.Scanf("%d", &num)

	for i := 0; i <= 10; i++ {
		fmt.Printf("%d X %d = %d\n", num, i, num*i)
	}
}
