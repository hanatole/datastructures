package main

import (
	"container/list"
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

	queue := list.New()
	queue.PushBack(root)

	i := 1

	for queue.Len() > 0 && i < len(data) {
		element := queue.Front()
		queue.Remove(element)

		node := element.Value.(*Node)

		if i < len(data) && data[i] != "NULL" {
			value, err := strconv.Atoi(data[i])
			if err == nil {
				node.left = NewNode(value)
				queue.PushBack(node.left)
			}
		}
		i++

		if i < len(data) && data[i] != "NULL" {
			value, err := strconv.Atoi(data[i])
			if err == nil {
				node.right = NewNode(value)
				queue.PushBack(node.right)
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

	queue := list.New()
	queue.PushBack(n)

	for queue.Len() > 0 {
		element := queue.Front()
		queue.Remove(element)

		node := element.Value.(*Node)
		node.process()

		if node.left != nil {
			queue.PushBack(node.left)
		}

		if node.right != nil {
			queue.PushBack(node.right)
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
