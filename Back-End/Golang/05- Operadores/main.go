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
	fmt.Println("And: True && False", true && false)
	fmt.Println("Ou: true || false", true || false)
	fmt.Println("Inversão: !true", !true)

	//Unários
	num := 10
	num++ // Ou num += VALOR DESEJADO
	fmt.Println("Num++: ", num)

	num-- // Ou num -=VALOR DESEJADO
	fmt.Println("Num--: ", num)

	num *= 3
	fmt.Println("Num *=3: ", num)

	num /= 2
	fmt.Println("Num/=2: ", num)

	num %= 3
	fmt.Println("Num %=2: ", num)

}
