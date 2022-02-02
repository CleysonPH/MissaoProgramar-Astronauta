package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	frase := scanner.Text()
	fraseReversa := ""
	for _, letra := range frase {
		fraseReversa = string(letra) + fraseReversa
	}
	fmt.Println(fraseReversa)
}
