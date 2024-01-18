package main

import "fmt"

func main() {

	//Primeira declaração de array
	var array1 [5]int
	array1[0] = 12

	// [12 0 0 0 0]
	fmt.Println(array1)

	//Declarando array e ja passando os valores
	array2 := [5]string{"Arthur", "Lucas", "Patrick", "Carlos", "Gustavo"}

	fmt.Println(array2)

	//Array Flexivel, quando passamos ... dizemos ao go que o valor do array é quantidade de posições existente

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	//SLICE
	slice1 := []int{10, 30, 40, 20, 50}
	fmt.Println(slice1)
	slice1 = append(slice1, 60) //Adicionando um novo valor ao slice
	fmt.Println(slice1)

	//Slices podem ser uma parte de um array que ja exista, ou seja apontam para aquele array(como se fosse um ponteiro)
	slice2 := array2[1:3]
	fmt.Println(slice2)

	//Trocando o valor do array e verificando se mudou o Slice
	array2[1] = "jorge"
	fmt.Println(slice2)

	//ARRAY INTERNOS
	slice3 := make([]string, 3, 5)
	fmt.Println(slice3)                                //Ver slice
	fmt.Println("Tamanho Slice3: ", len(slice3))       //Ver tamamho
	fmt.Println("Capacidade do Slice3: ", cap(slice3)) //Ver capacidade

	//Adicionando um novo valor no slice
	slice3 = append(slice3, "Arthur", "Lucas", "Patrick", "Carlos", "Gustavo")
	fmt.Println(slice3) //Ver slice

	slice3 = append(slice3, "Jorge")                   //Estourando o tamanho setado do slice
	fmt.Println("Tamanho Slice3: ", len(slice3))       //Ver tamamho
	fmt.Println("Capacidade do Slice3: ", cap(slice3)) //Ver capacidade

	//Se passarmos para o slice somento o numero de possições o tamanho vai ser o mesmo valor das possições
	slice4 := make([]float32, 5)
	fmt.Println("Tamanho Slice3: ", len(slice4))
	fmt.Println("Capacidade do Slice3: ", cap(slice4))
}
