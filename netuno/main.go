package main

import "fmt"

func main() {
	var a int
	fmt.Scanf("%d", &a)
	if a%2 == 0 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
