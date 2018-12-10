package tree

import "fmt"

type TreeNode struct {
	Value int
	Left, Right *TreeNode
}

//结构体赋值方法
func  Assignment() {
	var root TreeNode
	root = TreeNode{Value: 3}
	root.Left = &TreeNode{}
	root.Right = &TreeNode{5, nil, nil}
	newTree := new(TreeNode)
	fmt.Println(root)
	fmt.Println(newTree)

	sliceTree := []TreeNode{
		{3, nil, nil},
		{},
		{Value: 4},
	}

	fmt.Println(sliceTree)
}

func (node *TreeNode) createNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

func(node *TreeNode) print() {
	fmt.Print(node.Value)
}

func(node *TreeNode) traverse() {
	if node == nil {
		return
	}

	node.Left.traverse()
	node.print()
	node.Right.traverse()
}
