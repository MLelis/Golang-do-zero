package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"Marcus", "Jordhan", 20, 181}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "UFS"}
	fmt.Println(e1)
	fmt.Println(e1.nome)
	fmt.Println(e1.altura)
}
