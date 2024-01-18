package main

import "fmt"

func calc(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

//Juntando uma função nomeada com uma variatica
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	total := calc(1, 2, 5, 3, 2, 1, 3)
	fmt.Println(total)

	escrever("O numero é: ", 1, 62, 31, 41, 25, 6, 7, 8, 9)
}
