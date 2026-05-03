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

func (s *Set) binarySearch(value int) int {
    left, right := 0, s.size-1

    for left <= right {
        mid := (left + right) / 2

        if s.data[mid] == value {
            return mid
        } else if s.data[mid] < value {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return -1
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
func (s *Set) reserve(newCapacity int) {
    newData := make([]int, newCapacity)

    for i := 0; i < s.size; i++ {
        newData[i] = s.data[i]
    }

    s.data = newData
    s.capacity = newCapacity
}

func (s *Set) insert(value int, index int) error {
    if index < 0 || index > s.size {
        return fmt.Errorf("index out of range")
    }

    if s.size == s.capacity {
        if s.capacity == 0 {
            s.reserve(1)
        } else {
            s.reserve(s.capacity * 2)
        }
    }

    for i := s.size; i > index; i-- {
        s.data[i] = s.data[i-1]
    }

    s.data[index] = value
    s.size++

    return nil
}
func (s *Set) Insert(value int) {
    if s.binarySearch(value) != -1 {
        return
    }

    index := 0
    for index < s.size && s.data[index] < value {
        index++
    }

    _ = s.insert(value, index)
}
func (s *Set) erase(index int) error {
    if index < 0 || index >= s.size {
        return fmt.Errorf("index out of range")
    }

    for i := index; i < s.size-1; i++ {
        s.data[i] = s.data[i+1]
    }

    s.size--
    return nil
}

func (s *Set) Erase(value int) bool {
    index := s.binarySearch(value)
    if index == -1 {
        return false
    }

    _ = s.erase(index)
    return true
}


func (s *Set) Contains(value int) bool {
    return s.binarySearch(value) != -1
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
			 v = NewSet(value)
		case "insert":
		if len(parts) < 2 {
			continue
		}

		for _, part := range parts[1:] {
			value, _ := strconv.Atoi(part)
			v.Insert(value)
		}
		case "show":
			fmt.Printf("[%s]\n", Join(v.data[:v.size], ", "))
		case "erase":
		if len(parts) < 2 {
			continue
		}

		value, _ := strconv.Atoi(parts[1])
		if !v.Erase(value) {
			fmt.Println("value not found")
		}
		case "contains":
		if len(parts) < 2 {
			continue
		}

		value, _ := strconv.Atoi(parts[1])

		if v.Contains(value) {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
