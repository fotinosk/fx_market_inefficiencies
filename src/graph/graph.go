package graph

import (
	"fmt"
	utils "fxIneff/src/api"

	"golang.org/x/exp/slices"
)

func has_duplicates(path []string) bool {
	sub_array := path[1:]
	for ind, val := range sub_array {
		sub_sub_array := sub_array[ind+1:]
		if slices.Contains(sub_sub_array, val) {
			return true
		}
	}
	return false
}

func TraverseGraph(path []string) (float64, error) {
	if path[0] != path[len(path)-1] {
		return 0.0, fmt.Errorf("path given is not a cycle - the path must start and end at the same currency")
	}

	has_dupl := has_duplicates(path)
	if has_dupl {
		return 0.0, fmt.Errorf("node visited mutliple times, this is forbiden")
	}

	vertices := utils.Generate_nodes()
	start_point := path[0]
	returns := 10000.0
	current_node := start_point

	for _, next_node := range path[1:] {
		rate := vertices[current_node][next_node]
		fmt.Printf("Going from node %s to %s with rate %f \n", current_node, next_node, rate)
		returns = returns * rate

		current_node = next_node
	}

	return returns, nil
}

func TraversalAlgorithm(vertices map[string]map[string]float64) []string {
	// implement here the traversal algorithm

	// step 1: find the biggest discrepancies this is O(n^2)

	var result []string
	return result
}
