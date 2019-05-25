package main

import (
	"fmt"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func CreateNode(val int) *Node {
	node := new(Node)
	node.value = val
	node.left = nil
	node.right = nil
	return node
}

func LCA(root *Node, val1 int, val2 int) *Node {
	if root == nil {
		return nil
	}
	if root.value == val1 || root.value == val2 {
		return root
	}
	root.left = LCA(root.left, val1, val2)
	root.right = LCA(root.right, val1, val2)
	if root.left != nil && root.right != nil {
		return root
	}
	if root.left == nil && root.right == nil {
		return nil
	}
	if root.left != nil {
		return root.left
	} else {
		return root.right
	}
}

func main() {
	root := CreateNode(1)
	root.left = CreateNode(2)
	root.right = CreateNode(3)
	root.left.left = CreateNode(4)
	root.left.right = CreateNode(5)
	root.right.left = CreateNode(6)
	root.right.right = CreateNode(7)
	root.right.right.right = CreateNode(8)
	fmt.Println(LCA(root, 5, 8).value)
}
