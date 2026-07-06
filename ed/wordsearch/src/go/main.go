package main

import (
	"bufio"
	"fmt"
	"os"
	//"text/template"
)

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	_, _ = grid, word
	n := len(grid)
	m := len(grid[0])

	var dfs func(int, int, int)bool 

	dfs = func(l,c,idx int)bool{
		if idx == len(word){
			return true
		}

		if l < 0 || l >= n || c < 0 || c >= m {
			return false
		}
		if grid[l][c] != word[idx]{
			return false
		}
		tem := grid[l][c]
		grid[l][c] = '#'

		ok := dfs(l-1,c,idx+1) ||
		dfs(l+1,c,idx+1) ||
		dfs(l,c-1,idx+1) ||
		dfs(l,c+1,idx+1)
		
		grid[l][c] = tem
		return  ok
	}
	for l := 0 ; l < n ; l++{
		for c := 0 ; c < m ; c++{
			if grid[l][c] == word[0]{
				if dfs(l,c,0){
					return true
				}
				
			}
	    }
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
