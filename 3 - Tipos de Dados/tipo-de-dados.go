package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = 100000000000
	numero2 := 100000000
	fmt.Println(numero, numero2)

	var numero3 uint32 = 1000
	fmt.Println(numero3)

	// INT32 = rune
	var numero4 rune = 123456
	fmt.Println(numero4)

	// BYTE = INT8
	var numero5 byte = 123
	fmt.Println(numero5)

	// FLOAT
	var real1 float32 = 123.45
	fmt.Println(real1)

	var real2 float64 = 12311111.45
	fmt.Println(real2)

	real3 := 12345.67
	fmt.Println(real3)

	// STRINGS
	var caracteres string = "texto"
	fmt.Println(caracteres)

	caracteres2 := "Texto2"
	fmt.Println(caracteres2)

	char := 'B'
	fmt.Println(char)

	var texto string
	fmt.Println(texto)

	// Booleano
	var booleano1 bool = true
	fmt.Println(booleano1)

	// Erro
	var erro error = errors.New("Erro Interno")
	fmt.Println(erro)
}
