package main

import "fmt"

func main() {
	var texto1 string = "texto1"
	texto2 := "texto2"

	fmt.Println(texto1)
	fmt.Println(texto2)

	var (
		texto3 string = "texto3"
		texto4 string = "texto4"
	)

	fmt.Println(texto3)
	fmt.Println(texto4)

	texto5, texto6 := "texto5", "texto6"

	fmt.Println(texto5)
	fmt.Println(texto6)

	texto5, texto6 = texto6, texto5

	fmt.Println("recebeu o valor do 6 :", texto5)
	fmt.Println("recebeu o valor do 5 :", texto6)
}
