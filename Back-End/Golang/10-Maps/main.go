package main

import "fmt"

func main() {
	// [tipo da chave]tipo do valor
	usuario := map[string]string{
		"nome":      "Arthur",
		"sobrenome": "Jesus",
	}

	fmt.Println(usuario)

	fmt.Println(usuario["nome"]) //Acessar o nome

	usuario2 := map[string]map[string]string{
		"nome": {
			"PrimeiroNome": "Cleiton",
			"SegundoNome":  "Rasta",
		},
		"curso": {
			"NomeCurso": "ADS",
			"Campus":    "1",
		},
	}

	fmt.Println(usuario2)

	//deletar uma chave do map
	delete(usuario2, "curso")
	fmt.Println(usuario2)

	//adicionar uma chave
	usuario2["endereco"] = map[string]string{
		"Bairro": "Sol Nascente",
		"Rua":    "Solaris",
		"Numero": "81",
	}
	fmt.Println(usuario2)

}
