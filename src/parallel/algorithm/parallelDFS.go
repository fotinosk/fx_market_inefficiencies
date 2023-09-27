// package algorithm
package main

import (
	"fmt"
	"fxIneff/src/parallel/matrix"
	"fxIneff/src/parallel/utils"
	"time"

	"golang.org/x/exp/slices"
)

type AlphaDict struct {
	name  string
	alpha float64
}

func initialize_path(start_node string, node string, matrix matrix.Matrix, channel chan AlphaDict) {
	alpha := matrix.GetAlpha(start_node, start_node, node)
	dict := AlphaDict{node, alpha}
	channel <- dict
}

func getGreedyNext(
	start_node string,
	current_node string,
	already_visited []string,
	currency_matrix matrix.Matrix,
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

func explore_path(current_path []string, matrix matrix.Matrix, channel chan []string) {
	current_node := current_path[len(current_path)-1]
	base_currency := current_path[0]

	next := getGreedyNext(base_currency, current_node, current_path[1:], matrix)
	current_path = append(current_path, next)

	if next == base_currency {
		channel <- current_path
	} else {
		go explore_path(current_path, matrix, channel)
	}
}

func traverseGraph(path []string, matrix matrix.Matrix) (float64) {
	start_point := path[0]
	returns := 1.0
	current_node := start_point

	for _, next_node := range path[1:] {
		rate, _ := matrix.GetExchangeRate(current_node, next_node)
		returns = returns * rate
		current_node = next_node
	}

	return returns
}

func ParallelDFS(start_node string) ([]string, float64) {
	var ConvMatrix matrix.Matrix
	ConvMatrix.PopulateMatrix()
	currencies := ConvMatrix.GetCurrencies()

	// start by ordering the next steps according to alpha
	alpha_map := make(map[string]float64, len(currencies))
	channel := make(chan (AlphaDict))

	for _, curr := range currencies {
		go initialize_path(start_node, curr, ConvMatrix, channel)
	}

	counter := 0
	for counter < len(currencies) {
		i := <-channel
		alpha_map[i.name] = i.alpha
		counter++
	}

	alpha_map = utils.SortMapByValue(alpha_map)

	path_channel := make(chan([]string))

	for key := range alpha_map {
		path := []string{start_node, key}
		go explore_path(path, ConvMatrix, path_channel)
	}

	// how to know when the channel is empty
	// for res := range path_channel {
	// 	ret := traverseGraph(res, ConvMatrix)
	// 	fmt.Println(res, ret)
	// }
	for {
		select {
			case res := <- path_channel:
				ret := traverseGraph(res, ConvMatrix)
				fmt.Println(res, ret)
			case <- time.After(10 * time.Millisecond):
				fmt.Println("0.01sec passed")
				return []string{"abc"}, 0.1
		}
	}

	// return []string{"abc"}, 0.1
}

func main() { ParallelDFS("eur") }
