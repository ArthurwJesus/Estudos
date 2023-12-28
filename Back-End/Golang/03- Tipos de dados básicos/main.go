package main

import "fmt"

func main() {
	var numero int16 = 10000

	//var numero2 int8 = 1000 -> aqui acontece um overflow pois o tamanho é maior
	// var numero2 uint8 = -8 -> aqui o erro é referente ao uint que não permite valores com sinal

	fmt.Println(numero)

	var numero2 float32 = 313212.32
	fmt.Println(numero2)

	numero3 := 32.2321
	fmt.Println(numero3)

}
