package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

// 和 func print(node treeNode) 一样 不过调用的时候要 print(node)
func (node Node) Print() {
	fmt.Println(node.Value)
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Srtting value to nil node. Ignored.")
		return
	}
	node.Value = value
}

func CreateNode(value int) *Node {
	return &Node{Value: value} //返回局部变量的地址 在 go 中也可以
}
