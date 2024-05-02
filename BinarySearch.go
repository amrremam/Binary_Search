package main

import "fmt"

var count int


// Node of type struct
type Node struct {
	Root int
	Left *Node
	Right *Node
}


//Insert Method
func (n *Node) insert(k int){
	if n.Root < k {   // MoveRight
		if n.Right == nil{
			n.Right = &Node{Root: k}
		}else {
			n.Right.insert(k)
		}

	}else if n.Root > k{  // MoveLeft
		if n.Left == nil{
			n.Left = &Node{Root: k}
		}else {
			n.Left.insert(k)
		}
	}
}


// Search Method
func (n *Node) search(k int) bool{
	count++

	if n == nil{
		return false
	}

	if n.Root < k{
		return n.Right.search(k)
	}else if n.Root > k{
		return n.Left.search(k)
	}
	return true
}


func main(){
	tree := &Node{Root: 100}

	tree.insert(120)
	tree.insert(200)
	tree.insert(60)
	tree.insert(90)
	tree.insert(30)

	fmt.Println(tree)
	fmt.Println(tree.search(30))
	fmt.Println(count)
}
