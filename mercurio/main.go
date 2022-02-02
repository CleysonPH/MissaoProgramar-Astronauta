package main

import "fmt"

func main() {
	var salary, payRisePercentage, payRise, newSalary float64
	fmt.Scanf("%f", &salary)

	if salary <= 280 {
		payRisePercentage = 20.0 / 100.0
	} else if salary <= 700 {
		payRisePercentage = 15.0 / 100.0
	} else if salary <= 1500 {
		payRisePercentage = 10.0 / 100.0
	} else {
		payRisePercentage = 5.0 / 100.0
	}

	payRise = salary * payRisePercentage
	newSalary = salary + payRise

	fmt.Printf("Salário: %.2f\n", salary)
	fmt.Printf("Percentual de aumento: %.2f%%\n", payRisePercentage*100.0)
	fmt.Printf("Valor do aumento: %.2f\n", payRise)
	fmt.Printf("Novo salário: %.2f\n", newSalary)
}
