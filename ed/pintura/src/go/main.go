package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calcilhas(image [][]int, i, j, co, nvc int){
	if i < 0 || j < 0 || i >= len(image) || j >= len(image[0]){
		return 
	}
	if image[i][j] != co {
		return
	}

	image[i][j] = nvc

	calcilhas(image, i+1 ,j , co, nvc)
	calcilhas(image, i ,j+1, co, nvc )
	calcilhas(image, i-1 ,j , co, nvc)
	calcilhas(image, i ,j-1 ,co, nvc)

}
// Não modifique a assinatura da função floodFill
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	//
	co := image[sr][sc]
	//if image[sr][sc] == 
	//_ := image
	if co == color{
		return image
	}
	calcilhas(image, sr ,sc ,co, color)
	return image
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	parts := strings.Fields(line)
	nl, _ := strconv.Atoi(parts[0])
	nc, _ := strconv.Atoi(parts[1])

	image := make([][]int, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		rowStr := strings.Fields(scanner.Text())
		row := make([]int, nc)
		for j := 0; j < nc; j++ {
			row[j], _ = strconv.Atoi(rowStr[j])
		}
		image[i] = row
	}

	// Lê sr, sc e color
	scanner.Scan()
	params := strings.Fields(scanner.Text())
	sr, _ := strconv.Atoi(params[0])
	sc, _ := strconv.Atoi(params[1])
	color, _ := strconv.Atoi(params[2])

	result := floodFill(image, sr, sc, color)

	// Imprime a matriz resultante
	for _, row := range result {
		for j, val := range row {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(val)
		}
		fmt.Println()
	}
}
