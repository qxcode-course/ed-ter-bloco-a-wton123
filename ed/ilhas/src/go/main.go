package main

import (
	"bufio"
	"fmt"
	"os"
)
func calcilhas(grid [][]byte, i, j int){
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]){
		return 
	}
	if grid[i][j] == '0'{
		return
	}

	grid[i][j] = '0'
	calcilhas(grid, i+1 ,j )
	calcilhas(grid, i ,j+1 )
	calcilhas(grid, i-1 ,j )
	calcilhas(grid, i ,j-1 )

}

// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func numIslands(grid [][]byte) int {
	//
	if len(grid) == 0 {
		return 0
	}
	n := 0
	for i := 0; i < len(grid); i++{
		for j := 0; j < len(grid); j++{
			if grid[i][j] == '1'{
						calcilhas(grid,i,j)
						n++
			}
		

	}
	}
	return n
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
}
