package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 2 {
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second)
	}

	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j")
		time.Sleep(time.Second)
	}

	for j := 0; j < 10; j += 5 {
		fmt.Println("Incrementando j")
		time.Sleep(time.Second)
	}

	nomes := [3]string{"João", "Davi", "Lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	nomes2 := []string{"João", "Davi", "Lucas"}

	for _, nome := range nomes2 {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Marcus",
		"sobrenome": "Jordhan",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}
}
