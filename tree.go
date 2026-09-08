package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	v     int
	left  *Node
	right *Node
}

func NewNode(v int) *Node {
	return &Node{v: v}
}

func fromText(text string) *Node {
	data := strings.Fields(text)

	if len(data) == 0 || data[0] == "NULL" {
		return nil
	}

	value, err := strconv.Atoi(data[0])
	if err != nil {
		return nil
	}

	root := NewNode(value)
	queue := []*Node{root}
	head, i := 0, 1

	for head < len(queue) && i < len(data) {
		node := queue[head]
		head++

		if i < len(data) && data[i] != "NULL" {
			value, err := strconv.Atoi(data[i])
			if err == nil {
				node.left = NewNode(value)
				queue = append(queue, node.left)
			}
		}
		i++

		if i < len(data) && data[i] != "NULL" {
			value, err := strconv.Atoi(data[i])
			if err == nil {
				node.right = NewNode(value)
				queue = append(queue, node.right)
			}
		}
		i++
	}

	return root
}

func (n *Node) process() {
	fmt.Printf("%d ", n.v)
}

func (n *Node) bfs() {
	if n == nil {
		return
	}

	queue := []*Node{n}
	head := 0

	for head < len(queue) {
		node := queue[head]
		head++
		node.process()

		if node.left != nil {
			queue = append(queue, node.left)
		}

		if node.right != nil {
			queue = append(queue, node.right)
		}
	}

	fmt.Println()
}

func (n *Node) preorder() {
	if n == nil {
		return
	}

	n.process()
	n.left.preorder()
	n.right.preorder()
}

func (n *Node) inorder() {
	if n == nil {
		return
	}

	n.left.inorder()
	n.process()
	n.right.inorder()
}

func (n *Node) postorder() {
	if n == nil {
		return
	}

	n.left.postorder()
	n.right.postorder()
	n.process()
}

func main() {
	text := "5 1 4 NULL NULL 3 6"
	root := fromText(text)

	root.bfs()

	root.preorder()
	fmt.Println()

	root.inorder()
	fmt.Println()

	root.postorder()
	fmt.Println()
}
