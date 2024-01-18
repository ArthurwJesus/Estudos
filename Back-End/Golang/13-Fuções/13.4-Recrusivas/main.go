package main

import "fmt"

func fibo(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibo(posicao-2) + fibo(posicao-1)
}

func main() {

	posicao := uint(10)
	fmt.Println(fibo(posicao)) //Mostra apenas o valor final

	for i := uint(0); i < posicao; i++ {
		fmt.Println(fibo(i)) // Mostrando o caminho do fibonacci
	}
}
