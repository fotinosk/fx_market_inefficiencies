package graph

import (
	"fmt"
	utils "fxIneff/src/api"
	matrix "fxIneff/src/conversion_matrix"

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
	returns := 1.0
	current_node := start_point

	for _, next_node := range path[1:] {
		rate := vertices[current_node][next_node]
		returns = returns * rate

		current_node = next_node
	}

	return returns, nil
}

func getGreedyNext(
	start_node string,
	current_node string,
	already_visited []string,
	currency_matrix matrix.ConversionMatrix,
) string {
	// find X to maximize current node -> X -> start node

	best_next := start_node
	best_val := 1.0

	for _, currency := range currency_matrix.GetCurrencies() {
		if slices.Contains(already_visited, currency) {
			continue
		}

		alpha := currency_matrix.GetAlpha(start_node, start_node, currency)
		if alpha > best_val {
			best_val = alpha
			best_next = currency
		}
	}
	return best_next
}

func TraversalAlgorithm(vertices map[string]map[string]float64, start_node string) []string {
	/*
		greedy depth first algorithm
		step 1: starting from the start currency scan all other currencies and go to the one with the biggest alpha
			store the visited node so you don't repeat it later
		step 2: repeat step 1 until the biggest alpha is returning to the start currency
		If it ever goes negative or down then stop

		This means that the max length can be the number of currencies you have

		Note: later on, add condition to backtrack and explore other paths

	*/

	visited_nodes := make([]string, 0, len(vertices))

	var currency_matrix matrix.ConversionMatrix
	currency_matrix.PopulateConversionMatrix(vertices)

	current_node := getGreedyNext(start_node, start_node, visited_nodes, currency_matrix)
	visited_nodes = append(visited_nodes, current_node)


	i := 0
	for i < 100 { // will never get to 100 since we can only visit each node once
		current_node = getGreedyNext(start_node, current_node, visited_nodes, currency_matrix)
		visited_nodes = append(visited_nodes, current_node)
		
		if current_node == start_node {
			break
		}
		i++
	}
	full_path := append([]string{start_node}, visited_nodes...)
	fmt.Println(TraverseGraph(full_path))
	return visited_nodes
}
