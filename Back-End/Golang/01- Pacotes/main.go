package main

import (
	"fmt"
	"github.com/badoux/checkmail"
	"modulo/auxiliar"
)

func main() {
	fmt.Println("Escrevendo da main.go \n")
	auxiliar.RetornaTexto()
	err := checkmail.ValidateFormat("estudos@gmail.com")
	fmt.Println(err)
}
