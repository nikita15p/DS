package main

import "fmt"

type Node struct{
	value int
	left *Node
	right *Node
}

func CreateNode(val int)  *Node{

	node := new(Node)
	node.value = val
	node.left =nil
	node.right=nil
	return node
}

func LCA(root *Node,val1 int, val2 int)int{

	templ,tempr := root,root
//fmt.Println(temp.right.value)
	ch := make(chan int)

	go func() {

		for {
			if templ.left != nil {
				if templ.left.value == val1 || templ.left.value == val2 {
					break
				}
				templ = templ.left

			}
		}
		ch<-templ.value

	}()

	go func(){
		for {
			if tempr.right != nil {
				if tempr.right.value == val1 || tempr.right.value == val2 {
					break
				}
				tempr = tempr.right
			}else{
				tempr = root
			}

		}
		ch<-tempr.value
	}()

	return <-ch
}

func main(){

	root := CreateNode(1)
	root.left = CreateNode(2)
	root.right = CreateNode(3)
	root.left.left = CreateNode(4)
	root.left.right = CreateNode(5)
	root.right.left = CreateNode(6)
	root.right.right =  CreateNode(7)

	fmt.Println(LCA(root,2,3))

}
