package main

import "fmt"

func main() {
	s := newStack()
	s.push(&node{3})
	s.push(&node{5})
	s.nodes[0].printVal()
	s.push(&node{7})
	s.push(&node{9})
	fmt.Println(s.pop(), s.pop(), s.pop(), s.pop())

}

type node struct {
	value int
}

type stack struct {
	nodes []*node
	count int
}

func (n *node) printVal() int {
	fmt.Println((*n).value)
	return 5
}

func newStack() *stack {

	return &stack{}
}

func (s *stack) push(n *node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

func (s *stack) pop() *node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

// func (s *stack) printStack(){

// }
