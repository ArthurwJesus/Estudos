package main

import (
	"fmt"
	"time"
)

func main() {

	//Primera utilização do for
	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando")
		time.Sleep(time.Second)
	}
	fmt.Println(i)

	//Segunda utilização do for
	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j: ", j)
		time.Sleep(time.Second)
	}

	//Utlizando o For range em um slice ou array
	empresas := [3]string{"Uol", "G1", "Max"} //Se tirar o valor do array [] teremos um Slice

	for idx, nome := range empresas { // Se quisermos pegar apenas o nome temos que passar _,nome pois assim pedimos para ignorar o indice
		fmt.Println(idx, nome)
	}

	//Utilizando o For range em uma palavra
	for indice, letra := range "FORD KA" {
		fmt.Println(indice, string(letra)) // se pedirmos para o print mostrar apenas a letra ele irá nos trazer o codigo em ASC
	}

	//Utilziando o For range em um map
	usuario := map[string]string{
		"nome":      "Arthur",
		"sobrenome": "jesus",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, ":", valor)
	}

	//Looping infinito
	//for {
	//	fmt.Println("Executando eternamente")
	//	time.Sleep(time.Second)
	//}
}
