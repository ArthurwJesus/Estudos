package main

import "fmt"

func func1() {
	fmt.Println("Opa")
}

func func2() {
	fmt.Println("Epa")
}

func resultados(n1, n2 int) bool {

	defer fmt.Println("Média sendo retornada, o status do aluno é:") //Independente de qual caso acabe a função(true ou false), antes do retorno o Defer ira ser executado.

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false

}

func main() {
	defer func1() //O Opa só vai aparecer depois que o EPA e o resultado forem mostrado
	func2()
	fmt.Println(resultados(8, 6))
}
