package main

import "fmt"

func main() {
	var h, p, f, d int

	fmt.Scan(&h, &p, &f, &d)

	pos := f

	for {
		pos = pos + d

		if pos > 15 {
			pos = 0
		}

		if pos < 0 {
			pos = 15
		}

		if pos == h {
			fmt.Println("S")
			break
		}

		if pos == p {
			fmt.Println("N")
			break
		}
	}
}

 