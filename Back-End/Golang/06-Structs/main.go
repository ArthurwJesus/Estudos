package main

import "fmt"

//Primeira struct criada
//type livro struct {
//	nome  string
//	preco float32
//	autor string
//
//}

type livro struct {
	nome   string
	preco  float32
	autor  string
	genero genero
}

type genero struct {
	genero string
}

func main() {

	//Primeira maneira de popular uma Struct passando o nome do campo e o valor
	var livro1 livro

	genero2 := genero{"Aventura"}

	livro1.nome = "As aventuras de Rin"
	livro1.preco = 27.98
	livro1.autor = "Runing"
	livro1.genero = genero2

	fmt.Println(livro1)

	//Segunda maneria de popular uma Struct, utilizando apenas os valores
	genero1 := genero{"Guerra"}
	livro2 := livro{"As cronicas de Narnia", 10.99, "James", genero1}

	fmt.Println(livro2)

	//Maneira de popular apenas alguns dos campos da Struct

	livro3 := livro{nome: "Turma da monica"}
	fmt.Println(livro3)

}
