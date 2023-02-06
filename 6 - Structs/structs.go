package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouto string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo Structs")

	var u usuario
	u.nome = "Marcus"
	u.idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua teste", 1}

	u2 := usuario{"Jordhan", 27, enderecoExemplo}
	fmt.Println(u2)

	u3 := usuario{nome: "Marquinhos"}
	fmt.Println(u3)

	u4 := usuario{idade: 23}
	fmt.Println(u4)
}
