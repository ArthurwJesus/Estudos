package main

import "fmt"

func main() {
	var valor1 int = 10
	var valor2 int = valor1

	fmt.Println(valor1, valor2)

	valor1++
	fmt.Println(valor1, valor2)

	//Ponteiro
	var valor3 int  //valor
	var valor4 *int //Endereço de memoria

	valor3 = 5       //setando o valor pra variavel
	valor4 = &valor3 // setando o endereço de memoria para o valor 4 para que sempre possua o mesmo valor da outra variavel

	fmt.Println("Valor e endereço de memoria: ", valor3, valor4)

	fmt.Println("\n Valor  e desreferencia do endereço de memoria: ", valor3, *valor4) // o *valor4 aqui faz com que o go vá ate o endereço de memoria e busque apenas o valor

	valor3 *= 2
	fmt.Println("\n Valor e endereço de memoria depois da atualização de valor: ", valor3, valor4)

	fmt.Println("\n Valor  e desreferencia do endereço de memoria do novo valor: ", valor3, *valor4)

}
