package main

import "fmt"

func main() {

	//Função anonima padrão
	func() {
		fmt.Println("Olá mundo!")
	}()

	//Função anonima com passagem de parametro
	func(texto string) {
		fmt.Println(texto)
	}("Func anonima com passagem de parametro!")

	//Função anonima com retorno
	retornoFuncAnonima := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Retornando func anonima")

	fmt.Println(retornoFuncAnonima)

}
