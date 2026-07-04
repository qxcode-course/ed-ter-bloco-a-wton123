package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}
func preorder(n *Node, str *[]string){
	if n == nil {
		*str = append(*str, "#")
		return 
	}
	*str = append(*str, strconv.Itoa(n.Value) )
	preorder(n.Left, str)
	preorder(n.Right, str)
}

func insert(node *Node, value int)*Node{
	if node == nil {
			n := &Node{Value: value,
	          Left: nil,
			  Right: nil,}
		return n
	}

	if node.Value < value {
		node.Right = insert(node.Right, value)
	}
	if node.Value > value{
		node.Left = insert(node.Left, value)
		
	}
	
	return node
	
}

func BstInsert(values []int) *Node {
	var root *Node 
	for _, v := range values{
		root = insert(root, v)
	}
	return root
	
}

// Dica: crie um vetor compartilhado e vá preenchendo conforme anda na recursão
// Depois use o strings.Join para gerar o serial
func Serialize(root *Node) string {
	str := []string{}
	preorder(root, &str)
	return strings.Join(str," ")

}

// -----------------------------------------------------------------------------------
func BShow(node *Node, history string) {
	if node != nil && (node.Left != nil || node.Right != nil) {
		BShow(node.Left, history+"l")
	}
	for i := 0; i < len(history)-1; i++ {
		if history[i] != history[i+1] {
			fmt.Print("│   ")
		} else {
			fmt.Print("    ")
		}
	}
	if history != "" {
		if history[len(history)-1] == 'l' {
			fmt.Print("╭───")
		} else {
			fmt.Print("╰───")
		}
	}
	if node == nil {
		fmt.Println("#")
		return
	}
	fmt.Println(node.Value)
	if node.Left != nil || node.Right != nil {
		BShow(node.Right, history+"r")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	values := make([]int, 0, len(parts))
	for _, elem := range parts {
		v, err := strconv.Atoi(elem)
		if err == nil {
			values = append(values, v)
		}
	}
	root := BstInsert(values)
	BShow(root, "") // Chama a função de impressão formatada
	fmt.Println(Serialize((root)))
}
