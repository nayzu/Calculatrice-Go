package main

import (
	"fmt"
	"os"
)

func main() {

	var n1, n2 float64
	var operator string

	fmt.Print("Premier nombre: ")
	fmt.Scan(&n1)

	fmt.Print("Opérateur: ")
	fmt.Scan(&operator)

	fmt.Print("Deuxième nombre: ")
	fmt.Scan(&n2)

	var result float64

	switch operator {
	case "+":
		result = n1 + n2
	case "-":
		result = n1 - n2
	case "*":
		result = n1 * n2
	case "/":
		if n1 == 0 {
			fmt.Print("Divison de 0 impossible !")
			os.Exit(1)
		}
		result = n1 / n2
	default:
		fmt.Print("Nombre ou opérateur invalide !")
		os.Exit(1)
	}

	fmt.Printf("\nRésultat de votre calcul: %f", result)
}
