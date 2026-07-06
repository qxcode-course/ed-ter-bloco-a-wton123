package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func calcilhas(matrix [][]int, mm [][]int, i, j int)int{
	
	if mm[i][j] != 0 {
		return mm[i][j]
	}
	n := len(matrix)
	m := len(matrix[0])

	maior := 1
	
	dx := []int{-1, 1,0,0}
	dy := []int{0,0,-1,1}
	for w := 0 ;w < 4; w++{
		ni := i + dx[w]
		nj := j + dy[w]

		if ni >= 0 && ni < n && nj >= 0 && nj < m && matrix[ni][nj] > matrix[i][j]{
			tam := 1 + calcilhas(matrix, mm, ni ,nj )
			if tam > maior {
				maior = tam
			}
		}	
	} 


	
	mm[i][j] = maior
	return maior
}

func longestIncreasingPath(matrix [][]int) int {
	//
	if len(matrix) == 0 {
		return 0
	}
	n := len(matrix)
	m := len(matrix[0])
	mm := make([][]int,n )

	for i := range mm {
		mm[i] = make([]int, m)
	}
	r := 0
	for i := 0 ; i < n ; i++{
		for j := 0 ; j < m ; j++{
			ta := calcilhas(matrix, mm, i ,j )
			if ta > r{
				r = ta
			}
	}
	}
	
	return r
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	parts := strings.Fields(scanner.Text())
	if len(parts) < 2 {
		return
	}
	nl, _ := strconv.Atoi(parts[0])
	nc, _ := strconv.Atoi(parts[1])

	matrix := make([][]int, nl)
	for i := 0; i < nl; i++ {
		if !scanner.Scan() {
			return
		}
		tokens := strings.Fields(scanner.Text())
		row := make([]int, nc)
		for j := 0; j < nc && j < len(tokens); j++ {
			v, _ := strconv.Atoi(tokens[j])
			row[j] = v
		}
		matrix[i] = row
	}

	fmt.Println(longestIncreasingPath(matrix))
}
