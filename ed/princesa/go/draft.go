package main

import (
	"fmt"
)
func proximoVivo(j []int, pos int) int {
	n := len(j)
	for {
		pos = (pos + 1) % n
		if j[pos] != 0 {
			return pos
		}
	}
}

func printFila(jogadores []int, esp int) {
	fmt.Print("[ ")
	for i , v := range jogadores{
		 if jogadores[i] == 0{
				continue
		 }
		 if i == esp {
			v = jogadores[i]
			fmt.Print(v ,"> ")
		 }
		 if i != esp{
			fmt.Print(v ," ")
		 }
		
		
	}
	fmt.Println("]")
}

func main() {
	var n , e int 
	fmt.Scan(&n , &e)
	jogadores := make([]int, n)
	for i := 0 ; i < n ; i++{
		jogadores[i] = i + 1
	}
    pos := e - 1

	vivos := n

	for vivos > 1 {
		printFila(jogadores, pos)
		
		morre := proximoVivo(jogadores, pos)

		jogadores[morre] = 0
		vivos--

		pos = proximoVivo(jogadores, morre)
	}

	printFila(jogadores, pos)
	
}