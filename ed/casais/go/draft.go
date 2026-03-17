package main

import "fmt"

func main() {
	var n, q, ani int
	fmt.Scan(&n)

	a := make([]int, n)         
	b := make(map[int]bool)  

	for i := 0; i < n; i++ {
		fmt.Scan(&ani)
		a[i] = ani
	}

	for i := 0; i < n; i++ {
		if b[i] {
			continue 
		}

		for j := i + 1; j < n; j++ {
			if !b[j] && a[i]+a[j] == 0 { 
				q++
				b[i] = true
				b[j] = true
				break
			}
		}
	}

	fmt.Println(q)
}