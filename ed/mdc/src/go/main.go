package main

import (
	"fmt"
)

func mdc(a, b int) int {
	if a == 0{
		return b
	}else if b == 0{
		return a
	}
	return mdc(b, a%b)
	//Escreva A na forma do resto do quociente (A = B⋅Q + R), ou seja R = A % B.
   // Encontre o MDC(B,R) usando o Algoritmo Euclidiano, já que MDC(A,B) = MDC(B,R)

	return 0
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(mdc(a, b))
	
}
