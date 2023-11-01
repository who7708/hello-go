package graph

// 以下是一个用Go语言实现的构建有向图的函数。在这个实现中，图由一个map组成，其中键是节点，值是该节点所连接的节点列表。同时，点和边上都可以添加属性。请参考以下代码：

import (
	"fmt"
)

// Edge 结构体表示有向图的边
type Edge struct {
	To         string                 // 边的终点
	Attributes map[string]interface{} // 边上可以添加的属性
}

// Graph 结构体表示有向图
type Graph map[string][]Edge

// NewGraph 创建一个新的有向图
func NewGraph() Graph {
	return make(map[string][]Edge)
}

// AddNode 添加一个节点到图中
func (g Graph) AddNode(name string) {
	g[name] = []Edge{}
}

// AddEdge 添加一条边到图中
func (g Graph) AddEdge(from, to string, attributes map[string]interface{}) {
	edge := Edge{To: to, Attributes: attributes}
	g[from] = append(g[from], edge)
}

// PrintGraph 打印图的内容
func (g Graph) PrintGraph() {
	for node, edges := range g {
		fmt.Printf("Node: %s\n", node)
		for _, edge := range edges {
			fmt.Printf("    Edge To: %s, Attributes: %v\n", edge.To, edge.Attributes)
		}
	}
}

func GraphRun() {
	graph := NewGraph()

	// 添加节点
	graph.AddNode("Node1")
	graph.AddNode("Node2")
	graph.AddNode("Node3")

	// 添加边，并添加属性
	graph.AddEdge("Node1", "Node2", map[string]interface{}{"weight": 2.0})
	graph.AddEdge("Node2", "Node3", map[string]interface{}{"weight": 3.0})
	graph.AddEdge("Node3", "Node1", map[string]interface{}{"weight": 1.0})

	// 打印图的内容
	graph.PrintGraph()
}

// 这个程序创建了一个有向图，并添加了三个节点和三条边。每条边上都有一个"weight"属性。PrintGraph方法可以打印出图的内容，包括每个节点连接的所有边以及这些边的属性。
