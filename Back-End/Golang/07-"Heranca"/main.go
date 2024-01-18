package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {

	pessoa1 := pessoa{"Arthur", "Jesus", 22}

	fmt.Println(pessoa1)

	estudante1 := estudante{pessoa1, "Ciência da computação", "UPF"}
	fmt.Println(estudante1)
	//Com isso  conseguimos acessar os campos apenas com estudante1.nome e não estudante1.pessoa.estudante1

}
