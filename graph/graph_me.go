package main

import "fmt"

type NodeMe struct {
	ID    int
	Nodes []*NodeMe
}

func NewNodeMe(id int) *NodeMe {
	return &NodeMe{
		ID:    id,
		Nodes: []*NodeMe{},
	}
}

func (n *NodeMe) AddPath(toNode *NodeMe) {
	n.Nodes = append(n.Nodes, toNode)
}

type GraphMe struct {
	nodeMap map[int]*NodeMe
	Start   *NodeMe
}

func NewGraphMe(id int) *GraphMe {
	m := make(map[int]*NodeMe)
	m[id] = NewNodeMe(id)
	return &GraphMe{
		nodeMap: m,
		Start:   m[id],
	}
}

func (g *GraphMe) AddNode(from, to int) {
	fromNode, ok := g.nodeMap[from]
	if !ok {
		fromNode = NewNodeMe(from)
		g.nodeMap[from] = fromNode
	}
	toNode, ok := g.nodeMap[to]
	if !ok {
		toNode = NewNodeMe(to)
		g.nodeMap[to] = toNode
	}
	fromNode.AddPath(toNode)
}

func (g *GraphMe) HasCycle() bool {
	//dfs
	stack := make([]int, len(g.nodeMap))
	sIdx := 0
	visited := make([]bool, len(g.nodeMap))
	walking := make([]bool, len(g.nodeMap))

	stack[sIdx] = g.Start.ID
	sIdx++

	for sIdx != 0 {
		pnID := stack[sIdx-1]
		visited[pnID] = true
		walking[pnID] = true
		//sIdx--
		fmt.Println(pnID, visited, walking)
		visitCount := 0
		for _, node := range g.nodeMap[pnID].Nodes {
			if !visited[node.ID] {
				stack[sIdx] = node.ID
				sIdx++
				visitCount++
			} else if walking[node.ID] {
				fmt.Println("detect")
			}
		}
		if visitCount == 0 {
			walking[pnID] = false
			sIdx--
		}
	}
	return true
}

func main() {
	graph := NewGraphMe(0)
	graph.AddNode(0, 1)
	graph.AddNode(0, 2)
	graph.AddNode(2, 3)
	graph.AddNode(3, 5)
	graph.AddNode(5, 4)
	graph.AddNode(4, 2)

	fmt.Println(graph.HasCycle())
}
