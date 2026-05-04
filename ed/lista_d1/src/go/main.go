package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)
type Node struct{
	Value int
    next *Node
    prev *Node
}

type LList struct{
	root *Node
}

func NewLList() *LList {
    root := &Node{}

    root.next = root
    root.prev = root

    return &LList{
        root: root,
    }
}

func (ll *LList) String()string{
    str := "["
	n := ll.root.next
	final := true
		for n != ll.root{
			if !final{
				str += ", "
			}	
			str +=  fmt.Sprint(n.Value)
				
			n = n.next
			final = false
		}
	
	str += "]"

	return str
}
func (ll *LList) Size()int{
	qdt := 0 
	node := ll.root.next
		for node != ll.root{
		   
           qdt += 1
		   node = node.next
		}
	
	return qdt
}

func (ll *LList)Clear(){
	ll.root.next = ll.root
	ll.root.prev = ll.root
}

func (ll *LList) PushFront(value int){
	node := &Node{
        Value: value,
        next: ll.root.next,
        prev: ll.root,
    }
	ll.root.next.prev = node
    ll.root.next = node

}

func (ll *LList)PushBack(value int){
    n := &Node{
		Value:value ,
		next: ll.root,
		prev: ll.root.prev,
	}
	ll.root.prev.next = n
    ll.root.prev = n
}

func (ll *LList) PopFront(){
	f := ll.root.next 
	if ll.root == f{
		return
	}
	s := f.next
    ll.root.next = s
    s.prev = ll.root
	f.next = nil
	f.prev = nil
	
} 

func (ll *LList)PopBack()   {
	fim := ll.root.prev
	if ll.root == fim{
		return
	}
	pefim := fim.prev
    ll.root.prev = pefim
    pefim.next = ll.root
	fim.next = nil
	fim.prev = nil
	
} 
func main() {
	scanner := bufio.NewScanner(os.Stdin)
 	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			 fmt.Println(ll.String())
		case "size":
			 fmt.Println(ll.Size())
		case "push_back":
			 for _, v := range args[1:] {
			 	num, _ := strconv.Atoi(v)
			 	ll.PushBack(num)
			 }
		case "push_front":
			 for _, v := range args[1:] {
			 	num, _ := strconv.Atoi(v)
			 	ll.PushFront(num)
			 }
		case "pop_back":
			 ll.PopBack()
		case "pop_front":
			 ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
