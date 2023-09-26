package ext

// 扩展 Node
type myTreeNode struct {
	node *Node
}

// 新建一个struct，包含以前的
func (myNode *myTreeNode) postOder() {
	// 因为有可能把nil包进myTreeNode中，所以需要判断myNode.node == nil
	if myNode == nil || myNode.node == nil {
		return
	}
	// 这里是新建一个myTreeNode类型的结构体，myNode.node.Left是个指针
	left := myTreeNode{myNode.node.Left}
	left.postOder()
	right := myTreeNode{myNode.node.Right}
	right.postOder()
	myNode.node.Print()
}

func TestNodeExt() {
	var root = Node{Value: 3}
	root.Left = &Node{}
	root.Right = &Node{5, nil, nil}
	root.Right.Left = new(Node)
	root.Left.Right = createNode(8)
	root.Left.RealSetValue(7)
	pRoot := &root
	pRoot.RealSetValue(20)
	// 这里是新建一个myTreeNode类型的结构体，root并不是指针，所以要取地址
	myRoot := myTreeNode{&root}
	myRoot.postOder() // output: 8 7 0 5 20
}
