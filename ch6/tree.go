package main

type tree struct {
	value int
	right, left *tree
}

func (node *tree) travers() {
	if node == nil {
		return
	}

	node.left.travers()
	node.Print()
	node.right.travers()
}

func (node *tree) traversFunc(func(node *tree)) {
	if node == nil {
		return
	}

	node.left.traversFunc(f)
	f(node)
	node.right.traversFunc(f)
}

func main() {
	
}
