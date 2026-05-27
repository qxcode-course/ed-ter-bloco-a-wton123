package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

 func (q *Queue[T]) Enqueue(value T){
	node := &Node[T]{value , nil}
	if q.size == 0{
		q.head = node
		q.tail = node
		q.size++
		return
	}
	q.tail.next = node
	q.tail = node
    q.size++
 }
 func (q *Queue[T]) Dequeue() (T, bool){
	var z T
	if q.size == 0{
		
		return  z ,false
	}else if q.size == 1{
		
		q.head = nil
		q.tail = nil
    	q.size--
		return z,true
	}
	v := q.head.Value
	node := q.head.next
	q.head = node
    q.size--
	return v,true
 }
 func (q *Queue[T]) Peek() (T, bool){
	var z T
	if q.size == 0{
		
		
		return  z,false
	}
	v := q.head.Value
	return v ,true
 }
 func (q *Queue[T]) Size() int{
    return q.size
 }
 func (q *Queue[T]) IsEmpty() bool{
	if q.size == 0{
		return true
		
	}else{
		return  false
	}
   
 }
 func (q *Queue[T]) Clear(){
	q.size= 0
	q.tail = nil
	q.head = nil
 }


type Node[T any] struct {
	Value T
	next  *Node[T]
}

type Queue[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) String() string {
	result := "["
	for n := q.head; n != nil; n = n.next {
		if n != q.head {
			result += ", "
		}
		result += fmt.Sprintf("%v", n.Value)
	}
	return result + "]"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	queue := NewQueue[int]()

	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println("$" + line)
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}
		switch args[0] {
		case "end":
			break
		case "show":
			fmt.Println(queue)
		case "push":
			 for _, arg := range args[1:] {
			 	value, _ := strconv.Atoi(arg)
			 	queue.Enqueue(value)
			 }
		case "pop":
			 if _, ok := queue.Dequeue(); !ok {
			 	fmt.Println("falha: fila vazia")
			 }
		case "peek":
			 if value, ok := queue.Peek(); ok {
			 	fmt.Println(value)
			 } else {
			 	fmt.Println("falha: fila vazia")
			 }
		default:
			fmt.Println("Unknown command:", args[0])
		}
	}
}
