package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
	"strconv"
)

type Vector struct {
	data     []int
	size     int
	capacity int
}

func (v *Vector)Status() string{
	return fmt.Sprintf("size:%d capacity:%d", v.size, v.capacity)

}
func (v *Vector)Reserve(newCapacity int){
	if newCapacity <= v.capacity {
		return 
	}
	v2 := make([]int, newCapacity)
	for i := 0 ; i < v.size; i++{
		v2[i] = v.data[i]
	}
	v.data = v2
	v.capacity = newCapacity
}

func (v *Vector)PushBack(value int ){
	if v.size == v.capacity{
		if v.capacity == 0 {
			v.Reserve(1)
		}else{
			v.Reserve(v.capacity * 2)
		}
	
	}
     v.data[v.size] = value
	v.size++
}

func NewVector(capacity int) *Vector {
	return &Vector{
		data:     make([]int, capacity), // nunca use len(data) ou cap(data) ou qq método do go de manipulação de array
		size:     0,
		capacity: capacity,
	}
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

func (v *Vector) At(index int) (int, error){
	if index < 0 || index >= v.size {
		return 0, fmt.Errorf("index out of range")
	}
	return v.data[index], nil
}

 func  (v *Vector)Set(index int, value int) error {
		if value < 0 || index >= v.size{
			return fmt.Errorf("index out of range")
		}
		v.data[index] = value
		return nil
 }

 func (v *Vector)Clear() {
       t := v.size
	   for i := 0 ; i < t ; i++{
		  v.size--
	   }
 }

 func (v *Vector) PopBack()(int, error){
	//' Remove e retorna o último elemento do vetor. Retorna um erro se o vetor estiver vazio
	if v.size <= 0{
		return 0 ,fmt.Errorf("vector is empty")
	}
	penu := v.size - 1
	valor := v.data[penu]
	v.size--
	return valor , nil
 } 

 func (v *Vector)Insert(index int, value int) error{  
	 //Insere um valor no índice especificado, deslocando os elementos conforme necessário.
	 if 0 >= index && index >= v.size{
		return fmt.Errorf("vector is empty")
	}
	if v.capacity == v.size {
		if v.capacity == 0{
			v.Reserve(1)
		}else{
			v.Reserve(v.capacity * 2)
		}

	}
	for i := v.size; i > index; i-- {
    v.data[i] = v.data[i-1]
    }
	v.data[index] = value
    v.size++
     return nil
 }

func (v *Vector)Erase(index int) error  {   
	         //' Remove o elemento no índice especificado, deslocando os elementos conforme necessário.
	if v.size <= index || v.size <= 0 {
		return fmt.Errorf("index out of range")
	}

	for i := index ; i < v.size - 1 ; i++{
		v.data[i] = v.data[i + 1]
	}
	v.size--
	return nil
}

  func (v *Vector)IndexOf(value int) int  {
		//' Retorna o índice da primeira ocorrência do valor especificado, ou -1 se não for encontrado
		for i := 0 ; i < v.size ; i++{
			if value == v.data[i]{
				return i
			}
		}
		return -1
  }    
  
  func  (v *Vector)Contains(value int) bool {
		//  ' Verifica se o valor especificado existe no vetor
		for i := 0 ; i < v.size ; i++{
			if value == v.data[i]{
				return true
			}
		}
		return false
  }    
  
  func Slice(start int, end int) *Vector {
  /*   ' Retorna um novo vetor que é uma fatia do vetor original do índice start até o índice end (exclusivo).
 ' O método deve lidar com índices negativos e índices que excedem o tamanho do vetor de forma circular
' Ele não deve criar um novo bloco de memória para os elementos, mas sim compartilhar a mesma memória do vetor original*/

  }


func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewVector(0)
	_ = v
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
			 value, _ := strconv.Atoi(parts[1])
			 v = NewVector(value)
		case "push":
			 for _, part := range parts[1:] {
			 	value, _ := strconv.Atoi(part)
			 	v.PushBack(value)
			 }
		case "show":
			fmt.Printf("[%s]\n", Join(v.data[:v.size], ", "))
		case "status":
			fmt.Println(v.Status())
		case "pop":
			numero , err := v.PopBack()
			if err != nil {
				fmt.Println(err)
			 }else{
				_ = numero
			 }
		case "insert":
			index, _ := strconv.Atoi(parts[1])
		    value, _ := strconv.Atoi(parts[2])
			err := v.Insert(index, value)
			if err != nil {
			  fmt.Println(err)
			}
		case "erase":
			 index, _ := strconv.Atoi(parts[1])
			 err := v.Erase(index)
			 if err != nil {
			 	fmt.Println(err)
			 }
		case "indexOf":
			 value, _ := strconv.Atoi(parts[1])
			 index := v.IndexOf(value)
			 fmt.Println(index)
		case "contains":
			 value, _ := strconv.Atoi(parts[1])
			 if v.Contains(value) {
			 	fmt.Println("true")
			 } else {
			 	fmt.Println("false")
			 }
		case "clear":
			 v.Clear()
		case "capacity":
			// fmt.Println(v.Capacity())
		case "get":
			 index, _ := strconv.Atoi(parts[1])
			 value, err := v.At(index)
			 if err != nil {
			 fmt.Println(err)
			 } else {
			 	fmt.Println(value)
			 }
		case "set":
				index, _ := strconv.Atoi(parts[1])
				value, _ := strconv.Atoi(parts[2])
				err := v.Set(index, value)
				if err != nil {
				fmt.Println(err)
				 }
				
		case "reserve":
			 newCapacity, _ := strconv.Atoi(parts[1])
			 v.Reserve(newCapacity)
		case "slice":
			 start, _ := strconv.Atoi(parts[1])
			 end, _ := strconv.Atoi(parts[2])
			 slice := v.Slice(start, end)
			 fmt.Println(slice)
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
