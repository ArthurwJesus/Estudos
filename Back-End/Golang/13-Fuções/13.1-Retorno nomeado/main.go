package main

import "fmt"

//Retorno nomeado irá retornar uma varivel chamada soma e outra sub
func calculos(n1, n2 int) (soma int, sub int) {
	soma = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	//criamos as duas variaveis para recever uma a soma e a outra configuração
	sum, subs := calculos(3, 2)

	fmt.Println(sum, subs)

}
