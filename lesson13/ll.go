package main

import "fmt"

type Node struct {
	Val      int
	NextNode *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Printing() {
	for node, i := l.Head, 1; node != nil; node = node.NextNode {
		fmt.Println(i, "-", node.Val)
		i++
	}
}

func (l *LinkedList) AddToHead(value int) {
	node := &Node{Val: value}
	if l.Head != nil {
		node.NextNode = l.Head
	}
	l.Head = node
}

func (l *LinkedList) AddToEnd(value int) {
	newNode := &Node{Val: value}
	node := l.Head
	for ; node.NextNode != nil; node = node.NextNode {
	}
	node.NextNode = newNode
}

func main() {
	ll := LinkedList{}
	// ll.Head = &Node{1, &Node{Val: 45, NextNode: &Node{Val: 54}}}
	ll.AddToHead(45)
	ll.AddToHead(23)

	ll.AddToEnd(9876)
	ll.Printing()
}
