package main

import (
	"fmt"
)

func printFila(fila []int, pos int) {
	fmt.Print("[ ")
	for i, v := range fila {
		if i == pos {
			fmt.Printf("%d> ", v)
		} else {
			fmt.Printf("%d ", v)
		}
	}
	fmt.Println("]")
}

func main() {
	var N, E int
	fmt.Scan(&N, &E)

	fila := make([]int, N)
	for i := 0; i < N; i++ {
		fila[i] = i + 1
	}

	pos := (E - 1) % len(fila)

	for len(fila) > 0 {
		printFila(fila, pos )

		if len(fila) == 1 {
			break
		}

		morte := (pos + 1) % len(fila)

		fila = append(fila[:morte], fila[morte+1:]...)

		if morte > pos {
			pos = (pos + 1) % len(fila)
		} else {
			pos = pos % len(fila)
		}
	}
}