package main

import "fmt"

func main() {
	//ARTIMETICOS
	soma := 2 + 2
	subtracao := 4 - 1
	divisao := 10 / 2
	multiplicacao := 5 * 2
	restoDiv := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDiv)

	//ATRIBUIÇÃO
	var variavel1 string = "variavel 01"
	varivel2 := "variavel 02"
	fmt.Println(variavel1, varivel2)

	//RELACIONAIS
	fmt.Println("Maior que: ", 1 > 2)
	fmt.Println("Menor que: ", 1 < 2)
	fmt.Println("Maior igual que: ", 1 >= 2)
	fmt.Println("Menor igual que: ", 1 <= 2)
	fmt.Println("Atribuição: ", 1 == 2)
	fmt.Println("difrente de: ", 1 != 2)

	//LÓGICOS
}
