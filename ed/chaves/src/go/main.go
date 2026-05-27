package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)
func (q *Queue[T])partida(g1 int, g2 int){
	var t1 , t2 T
	t1 = q.Dequeue()
	t2 = q.Dequeue()
	if g1 > g2 {
		q.Enqueue(t1)
	}else{
		q.Enqueue(t2)
	}
}
func main() {
    scanner := bufio.NewScanner(os.Stdin)
	_ = scanner
	QQ := NewQueue[string]()
	QQ.Enqueue("A")
	QQ.Enqueue("B")
	QQ.Enqueue("C")
	QQ.Enqueue("D")
	QQ.Enqueue("E")
	QQ.Enqueue("F")
	QQ.Enqueue("G")
	QQ.Enqueue("H")
	QQ.Enqueue("I")
	QQ.Enqueue("J")
	QQ.Enqueue("K")
	QQ.Enqueue("L")
	QQ.Enqueue("M")
	QQ.Enqueue("N")
	QQ.Enqueue("O")
	QQ.Enqueue("P")

	for i := 0; i < 15 ; i++{
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		//fmt.Println("$" + line)
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		
		
		r1, _ := strconv.Atoi(args[0])
		r2, _ := strconv.Atoi(args[1])
		QQ.partida(r1,r2)

	}
	campeao := QQ.String()
	fmt.Println(campeao)
}
