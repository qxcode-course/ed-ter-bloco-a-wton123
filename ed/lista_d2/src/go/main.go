package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)
type Node struct {
	Value int
	next  *Node
	prev  *Node
	root  *Node
}

type LList struct{
	root *Node
}

func NewLList()*LList{
	 root := &Node{}

    root.next = root
    root.prev = root

    return &LList{
        root: root,
    }
}

func (ll *LList)String()string{
	n := ll.root.next
	str := "["
	ftemoq := true
	for ll.root != n {
		if !ftemoq {
			str += ", "
		}
		str += fmt.Sprint(n.Value)
		ftemoq = false
		n = n.next
	}
    str += "]"
	return str
}

func (ll *LList) PushBack(value int) {
	n := &Node{
		Value: value,
		next:  ll.root,
		prev:  ll.root.prev,
		root:  ll.root,
	}

	ll.root.prev.next = n
	ll.root.prev = n
}

func (ll *LList) PushFront(value int) {
	n := &Node{
		Value: value,
		next:  ll.root.next,
		prev:  ll.root,
		root:  ll.root,
	}

	ll.root.next.prev = n
	ll.root.next = n
}

 func (ll *LList)PopBack(){
	c:= ll.root.next 
	if ll.root == c{
		return
	}
	s := c.next
	ll.root.next = s
	s.prev = ll.root
	c.next = nil
	c.prev = nil
 }
  func (ll *LList)PopFront(){
		u := ll.root.prev 
	if ll.root == u{
		return
	}
	p := u.prev
	ll.root.prev = p
	p.prev = ll.root
	u.next = nil
	u.prev = nil
 }
 	
func (ll *LList) Front() *Node {
	if ll.root.next == ll.root {
		return nil
	}
	return ll.root.next
}

func (ll *LList) Back() *Node {
	if ll.root.prev == ll.root {
		return nil
	}
	return ll.root.prev
}
func (n *Node) Next() *Node{
	if n.next == n.root{
		return nil
	}
    return n.next
}
func (n *Node) Prev()*Node{
	if n.prev == n.root{
		return nil
	}
	return n.prev
}

func (ll *LList) Clear(){
	ll.root.next = ll.root
	ll.root.prev = ll.root
}

func (ll *LList) Size()int{
	qtd := 0
	node := ll.root.next
	for node != ll.root{
		   
           qtd += 1
		   node = node.next
		}
	return qtd
}
func (ll *LList)Search(value int) *Node {
	node := ll.root.next
	for node != ll.root{
		   if node.Value == value {
					return node
		   }
           
		   node = node.next
		}
	return nil
}
func (ll *LList) Insert(n *Node, value int ){
	if n.root == nil {
		return 
	}
	n1 := &Node{
		Value: value,
		next: n,
		prev: n.prev,
		root: ll.root,
	}
	n.prev.next = n1
	n.prev = n1
}

func (ll *LList) Remove(node *Node) *Node {
	if node == nil || node == ll.root {
		return nil
	}
	m := node.next

	node.prev.next = node.next
	node.next.prev = node.prev

	node.next = nil
	node.prev = nil

	if m == ll.root {
		return nil
	}

	return m
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
		case "walk":
			 fmt.Print("[ ")
			 for node := ll.Front(); node != nil; node = node.Next() {
			 	fmt.Printf("%v ", node.Value)
			 }
			 fmt.Print("]\n[ ")
			 for node := ll.Back(); node != nil; node = node.Prev() {
			 	fmt.Printf("%v ", node.Value)
			 }
			 fmt.Println("]")
		case "replace":
			 oldvalue, _ := strconv.Atoi(args[1])
			 newvalue, _ := strconv.Atoi(args[2])
			 node := ll.Search(oldvalue)
			 if node != nil {
			 	node.Value = newvalue
			 } else {
			 	fmt.Println("fail: not found")
			 }
		case "insert":
			 oldvalue, _ := strconv.Atoi(args[1])
			 newvalue, _ := strconv.Atoi(args[2])
			 node := ll.Search(oldvalue)
			 if node != nil {
			 	ll.Insert(node, newvalue)
			 } else {
			 	fmt.Println("fail: not found")
			 }
		case "remove":
			 oldvalue, _ := strconv.Atoi(args[1])
			 node := ll.Search(oldvalue)
			 if node != nil {
			 	ll.Remove(node)
			 } else {
			 	fmt.Println("fail: not found")
			 }
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
