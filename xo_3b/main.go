package main

import (
	"fmt"
	"strings"
)

func main() {
	var letra string
	fmt.Print("Insira uma letra: ")
	fmt.Scanf("%s", &letra)

	vogais := []string{"a", "e", "i", "o", "u"}
	if isInSliceCaseInsensitive(letra, vogais) {
		fmt.Println("A letra é : Vogal")
	} else {
		fmt.Println("A letra é : Consoante")
	}
}

func isInSliceCaseInsensitive(a string, slice []string) bool {
	for _, s := range slice {
		if strings.ToLower(a) == strings.ToLower(s) {
			return true
		}
	}
	return false
}
