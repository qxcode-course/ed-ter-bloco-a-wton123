package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"errors"
)

type Stack[T any] struct {
	data []T
}
func (s *Stack[T])Push(value T){
	s.data = append(s.data,value)
}
func (s *Stack[T])Pop() error{
	er := errors.New("stack is empty")
	if len(s.data) == 0 {
		return er
	}
	
	//s.data[len(s.data)-1] = 0 
	s.data = s.data[:len(s.data)-1]
	return nil
}
func (s *Stack[T])Peek() (T, error){
	var t T
	r := errors.New("stack is empty")
	if len(s.data) == 0 {
		return t,r
	}
	t = s.data[len(s.data) - 1]
	return t,nil
}
func (s *Stack[T])IsEmpty() bool{
	if len(s.data) == 0 {
		return true
	}else{
		return false }
}
func (s *Stack[T])Size() int{
	return len(s.data)
}
func (s *Stack[T])Clear(){
	s.data = s.data[:0]
}

func NewStack[T any](cap int) *Stack[T] {
	 v := make([]T, 0, cap)
	return &Stack[T]{data: v}
	
	
}
 func (s *Stack[T]) String() string {
 	output := ""
 	for i := range cap(s.data) {
		if i != 0 {
			output += ", "
 		}
 		if i < len(s.data) {
			output += fmt.Sprintf("%v", s.data[i])
		} else {
			output += "_"
		}
	}
	return output
 }

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	 v := NewStack[int](10)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			 cap, _ := strconv.Atoi(parts[1])
			 v = NewStack[int](cap)
		case "push":
			 for _, part := range parts[1:] {
			 	value, _ := strconv.Atoi(part)
			 	v.Push(value)
			 }
		case "debug":
			 fmt.Println(v)
		case "top":
			 top, err := v.Peek()
			 if err != nil {
			 	fmt.Println(err)
			 } else {
			 	fmt.Println(top)
			 }
		case "size":
			 fmt.Println(v.Size())
		case "pop":
			 err := v.Pop()
			 if err != nil {
			 	fmt.Println(err)
			 }
		case "clear":
			 v.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
