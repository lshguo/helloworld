package main

import (
	"fmt"
)

type Node struct{
	V int
	Left  *Node
	Right *Node
}
func (n *Node)AddLeft(l *Node){
	n.Left = l
}
func (n *Node)AddRight(r *Node){
	n.Right = r
}
func (n *Node)LNR(){
	if n.Left == nil && n.Right == nil{
		fmt.Printf("%d ", n.V)
		return
	}

	if l := n.Left; l != nil{
		l.LNR()
	}

	fmt.Printf("%d ", n.V)

	if r := n.Right; r != nil{
		r.LNR()
	}
}
//var a = []int{3,2,1,4,5}

func main(){
	    // a tree like this:
	    //      0
	    //   1    2
	    //  3 4  5 6
	var a []*Node
	for i := 0;i < 7;i++{
		n := new(Node)
		n.V = i
		a = append(a, n)
	}
	a[0].Left = a[1]
	a[0].Right = a[2]

	a[1].Left = a[3]
	a[1].Right = a[4]

	a[2].Left = a[5]
	a[2].Right = a[6]

	a[0].LNR()
}