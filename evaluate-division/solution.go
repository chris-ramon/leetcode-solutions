// # Complexity
// - Time complexity: `O(N)`, where N is the number of equations.
// - Space complexity: `O(N)`, where N is the size of the map used to store the graph.

type Edge struct {
	Value float64
	Node  *Node
}

type Node struct {
	Key      string
	Outbound map[string]*Edge
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := map[string]*Node{}

	for _, equation := range equations {
		x, y := equation[0], equation[1]
		if node := findNodeByKey(graph, x); node == nil {
			addNode(graph, newNode(x))
		}
		if node := findNodeByKey(graph, y); node == nil {
			addNode(graph, newNode(y))
		}
	}

	for i, equation := range equations {
		x, y := equation[0], equation[1]
		xNode := findNodeByKey(graph, x)
		yNode := findNodeByKey(graph, y)
		value := values[i]

		xNode.Outbound[yNode.Key] = &Edge{Node: xNode, Value: 1}
		yNode.Outbound[xNode.Key] = &Edge{Node: yNode, Value: 1}

		xNode.Outbound[yNode.Key] = &Edge{Node: yNode, Value: value}
		yNode.Outbound[xNode.Key] = &Edge{Node: xNode, Value: 1 / value}
	}

	result := []float64{}
	for _, query := range queries {
		v1, v2 := query[0], query[1]
		visited := map[string]bool{v1: true}
		result = append(result, dfs(graph, v1, v2, visited))
	}

	return result
}

func dfs(graph map[string]*Node, v1 string, v2 string, visited map[string]bool) float64 {
	if _, found := graph[v1]; !found {
		return -1
	}

	node := findNodeByKey(graph, v1)

	if edge := node.Outbound[v2]; edge != nil {
		return edge.Value
	}

	for _, edge := range node.Outbound {
		if _, found := visited[edge.Node.Key]; found {
			continue
		}

		visited[edge.Node.Key] = true

		val := dfs(graph, edge.Node.Key, v2, visited)
		if val != -1 {
			return edge.Value * val
		}
	}

	return -1
}

func findNodeByKey(graph map[string]*Node, key string) *Node {
	if node, found := graph[key]; found {
		return node
	}
	return nil
}

func addNode(graph map[string]*Node, node *Node) {
	graph[node.Key] = node
}

func newNode(key string) *Node {
	return &Node{Key: key, Outbound: make(map[string]*Edge, 0)}
}
