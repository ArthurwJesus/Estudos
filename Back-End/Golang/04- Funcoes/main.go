package main

import "fmt"

func operations(n1, n2 int32) (int32, int32) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func sum(n1 int32, n2 int32) int32 {
	return n1 + n2
}

func main() {
	soma := sum(2, 4)
	fmt.Println("resultado da soma: ", soma)

	var minus = func(n1 int32, n2 int32) int32 {
		return n1 - n2
	}

	subtracao := minus(6, 4)
	fmt.Println("resultado da subtração: ", subtracao)

	resultsSum, resultsMinus := operations(10, 15)
	fmt.Println(resultsSum, resultsMinus)

	//utilizando apenas a soma
	resultsSoma, _ := operations(10, 15)
	fmt.Println(resultsSoma)
}
