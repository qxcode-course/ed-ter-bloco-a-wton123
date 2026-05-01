package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
		"strconv"
)

type Set struct {

	data []int                          
	size int                          
	capacity int  
}
func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return result.String()
}

func NewSet(capacity int) *Set {
	return &Set{
			data: make([]int, capacity),
			size: 0 ,
			capacity: capacity,
	}	
}
func (s *Set )Reserve(newCapacity int) {
	if s.capacity == 0 {
		s.capacity = 1
	}
	s2 := make([]int, newCapacity)
	for i := 0; i < s.size ; i++{
		s2[i] = s.data[i]
	}
	    s.data = s2
		s.capacity = newCapacity 
		
	
}
func (s *Set)insert(value int, index int) error {
	 if 0 > index || index < s.size{
		return fmt.Errorf("vector is empty")
	}
	if s.capacity == s.size {
		if s.capacity == 0{
			s.Reserve(1)
		}else{
			s.Reserve(s.capacity * 2)
		}
		

	}
    s.data[index] = value
    s.size++

  return nil
} 
func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	 v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
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
			 value, _ := strconv.Atoi(parts[1])
			 v := NewSet(value)
			 _ = v
		case "insert":
			 for _, part := range parts[1:] {
				_ = part
			 	value, _ := strconv.Atoi(parts[1])
				index, _ := strconv.Atoi(parts[2])
				 err := v.insert(value, index)
				if err != nil{
					fmt.Println(err)
				}
			 }
		case "show":
			fmt.Printf("[%s]\n", Join(v.data[:v.size], ", "))
		case "erase":
		    value, _ := strconv.Atoi(parts[1])
		case "contains":
			// value, _ := strconv.Atoi(parts[1])
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
