package main

import (
	"bufio"
	"fmt"
	"os"
)

func calcilhas(board [][]byte, visita [][]bool, i, j int){
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]){
		return 
	}
	if board[i][j] == 'X' || visita[i][j]{
		return
	}

	visita[i][j] = true

	calcilhas(board,visita ,i+1 ,j )
	calcilhas(board,visita, i ,j+1 )
	calcilhas(board,visita, i-1 ,j )
	calcilhas(board,visita, i ,j-1 )

}
// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	if len(board) == 0 {
		return 
	}
	n := len(board)
	m := len(board[0])
	visita := make([][]bool, n)
	for i := range visita{
		visita[i] = make([]bool, m)
	}

	for i := 0; i < n ; i++{
		if board[i][0] == 'O'{
			calcilhas(board, visita ,i ,0)
		}
		if board[i][m-1] == 'O'{
			calcilhas(board, visita ,i ,m-1)
		}
	}

	for j := 0; j < n ; j++{
		if board[0][j] == 'O'{
			calcilhas(board, visita ,0 ,j)
		}
		if board[n-1][j] == 'O'{
			calcilhas(board, visita ,n-1 ,j)
		}
	}

	for i := 0; i < n ; i++{
		for j := 0; j < m ; j++{
		if board[i][j] == 'O' && !visita[i][j]{
			board[i][j] = 'X'
		}
		
	}
	}
}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
