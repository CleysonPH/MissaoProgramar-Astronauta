package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		p1, p2, p3, p4, p5 int
		r                  string
	)
	fmt.Println("Telefonou para a vítima?")
	fmt.Print("(sim/não): ")
	fmt.Scanf("%s", &r)
	p1 = verificaResposta(r)

	fmt.Println("Esteve no local do crime?")
	fmt.Print("(sim/não): ")
	fmt.Scanf("%s", &r)
	p2 = verificaResposta(r)

	fmt.Println("Mora perto da vítima?")
	fmt.Print("(sim/não): ")
	fmt.Scanf("%s", &r)
	p3 = verificaResposta(r)

	fmt.Println("Devia para a vítima?")
	fmt.Print("(sim/não): ")
	fmt.Scanf("%s", &r)
	p4 = verificaResposta(r)

	fmt.Println("Já trabalhou com a vítima?")
	fmt.Print("(sim/não): ")
	fmt.Scanf("%s", &r)
	p5 = verificaResposta(r)

	veredito := p1 + p2 + p3 + p4 + p5
	if veredito >= 5 {
		fmt.Println("Assassino")
	} else if veredito >= 3 {
		fmt.Println("Cúmplice")
	} else if veredito >= 2 {
		fmt.Println("Suspeita")
	} else {
		fmt.Println("Inocente")
	}
}

func verificaResposta(resposata string) int {
	if strings.ToLower(resposata) == "s" || strings.ToLower(resposata) == "sim" {
		return 1
	}
	return 0
}
