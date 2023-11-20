package main

// Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visitedNodes := make(map[*Node]*Node) // key - is old node, value - is a new node
	dfs(node, visitedNodes)

	return visitedNodes[node]
}

func dfs(node *Node, visitedNodes map[*Node]*Node) {
	if node == nil {
		return
	}
	newNode := new(Node)
	newNode.Val = node.Val

	visitedNodes[node] = newNode

	for _, nn := range node.Neighbors {
		if _, ok := visitedNodes[nn]; !ok {
			dfs(nn, visitedNodes)
		}

		newNode.Neighbors = append(newNode.Neighbors, visitedNodes[node])
	}

	return
}
