package main

import (
	"fmt"

	"github.com/ronething/learngo/tree"
)

type myTreeNode struct {
	node *tree.Node
}

// 后续遍历
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node
	//fmt.Println(root)
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	root.Print()
	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	fmt.Println()
	root.Traverse()

	//var pRoot *tree.Node
	//fmt.Println(pRoot) // 这里 pRoot 是 nil
	//pRoot.SetValue(200)
	//pRoot = &root
	//pRoot.SetValue(300)
	//pRoot.Print()

	//pRoot := &root
	//pRoot.print()
	//pRoot.setValue(100)
	//pRoot.print()

	//nodes := []treeNode{
	//	{value: 3},
	//	{},
	//	{6, nil, &root},
	//}
	//fmt.Println(nodes)

	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()

}
