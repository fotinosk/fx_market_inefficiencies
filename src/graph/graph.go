package graph

import (
	"fmt"
	utils "fxIneff/src/api"
)

type Node struct {
	name        string
	conversions map[string]float32
}

func (n Node) print() {
	keys := make([]string, len(n.conversions)) // slight performance improvement since we know slice size
	i := 0
	for k := range n.conversions {
		keys[i] = k
		i++
	}
	fmt.Printf("Node for %s, with conversions to: %s\n\n", n.name, keys)
}

type Graph struct {
	vertices map[string]map[string]float32
}

func (graph Graph) initialize() []Node {
	// initialize the graph, ie convert the map of conversions to a list of Nodes
	edges := make([]Node, len(graph.vertices))
	ind := 0

	for key, value := range graph.vertices {
		node := Node{name: key, conversions: value}
		edges[ind] = node
		ind++
	}
	return edges
}

func TraverseGraph(path []string) (float32, error) {
	if path[0] != path[len(path)-1] {
		return 0.0, fmt.Errorf("path given is not a cycle - the path must start and end at the same currency")
	}

	// implement a check for duplicates - if it has duplicates then it has subcycles so raise error

	vertices := utils.Generate_nodes()
	start_point := path[0]
	returns := float32(1.0)
	current_node := start_point

	for _, next_node := range path[1:] {
		rate := vertices[current_node][next_node]
		fmt.Printf("Going from node %s to %s with rate %f \n", current_node, next_node, rate)
		returns = returns * rate

		current_node = next_node
	}

	return returns, nil
}

func Test_imports() {
	nodes := utils.Generate_nodes()
	fmt.Println("interal import working as expected")

	for name, conv := range nodes {
		n := Node{
			name:        name,
			conversions: conv,
		}
		n.print()
	}
}
