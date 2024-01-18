package main

import "fmt"

//Estrutura básica Switch
func diaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	}
	return "Dia da semana inválido!"
}

//Outra forma de trabalhar com switch

func diaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda-feira"
	case numero == 3:
		return "Terça-feira"
	case numero == 4:
		return "Quarta-feira"
	case numero == 5:
		return "Quinta-feira"
	case numero == 6:
		return "Sexta-feira"
	case numero == 7:
		return "Sábado"
	}
	return "Dia da semana inválido!"
}

func main() {

	//Estrutura básica IF
	numero := 7

	if numero > 15 {
		fmt.Println("É maior que 15")
	} else {
		fmt.Println("É menor ou igual a 15")
	}

	//Iniciando uma variavel dentro do if
	if numero2 := numero; numero2 > 0 {
		fmt.Println("Número é maior que 0")
	} else if numero2 < 0 {
		fmt.Println("Número é menor que 0")
	} else {
		fmt.Println("Número é igual a 0")
	}

	dia := diaSemana(5)
	fmt.Println(dia)

	dia2 := diaSemana2(7)
	fmt.Println(dia2)
}
