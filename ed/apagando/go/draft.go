package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	fila := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&fila[i])
	}

	var M int
	fmt.Scan(&M)

	// Usando map como conjunto
	saidas := make(map[int]bool)
	for i := 0; i < M; i++ {
		var x int
		fmt.Scan(&x)
		saidas[x] = true
	}

	// Filtrando resultado
	resultado := make([]int, 0)
	for _, pessoa := range fila {
		if !saidas[pessoa] {
			resultado = append(resultado, pessoa)
		}
	}

	for i, v := range resultado {
	if i > 0 {
		fmt.Print(" ")
	}
	fmt.Print(v)
}
fmt.Println(" ")
}
