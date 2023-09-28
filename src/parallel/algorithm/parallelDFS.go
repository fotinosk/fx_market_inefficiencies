package algorithm

// package main

import (
	"fxIneff/src/parallel/matrix"
	"fxIneff/src/parallel/utils"
	"sync"

	"golang.org/x/exp/slices"
)

type AlphaDict struct {
	name  string
	alpha float64
}

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc() {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c SafeCounter) Value() int {
	// c.mu.Lock()
	// // Lock so only one goroutine at a time can access the map c.v.
	// defer c.mu.Unlock()
	return c.v
}

// generate the advantage for a node
func initialize_path(start_node string, node string, matrix matrix.Matrix, channel chan AlphaDict) {
	alpha := matrix.GetAlpha(start_node, start_node, node)
	dict := AlphaDict{node, alpha}
	channel <- dict
}

// select the node with the best returns in a time horizon of 1
func getGreedyNext(
	start_node string,
	current_node string,
	already_visited []string,
	currency_matrix matrix.Matrix,
) string {

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

// recursively esplore the path, creating new goroutins as you go until reaching the start node again
func explore_path(current_path []string, matrix matrix.Matrix, channel chan []string, num_channels SafeCounter) {
	current_node := current_path[len(current_path)-1]
	base_currency := current_path[0]

	next := getGreedyNext(base_currency, current_node, current_path[1:], matrix)
	current_path = append(current_path, next)

	if next == base_currency {
		channel <- current_path
	} else {
		num_channels.Inc()
		go explore_path(current_path, matrix, channel, num_channels)
	}
}

// travel a set path to get the advantage generated
func traverseGraph(path []string, matrix matrix.Matrix) float64 {
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


/*
Main algorithm
Parallelized Depth First Search (truncated)
Get all the initial possible first next steps - which means all available currencies
and the in separate go routines explore the path greedily. Then get the best possible path
*/
func ParallelDFS(start_node string) ([]string, float64) {
	var ConvMatrix matrix.Matrix
	ConvMatrix.PopulateMatrix()
	currencies := ConvMatrix.GetCurrencies()
	currencies = utils.RemoveSliceElements(currencies, []string{start_node})

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

	path_channel := make(chan ([]string))
	num_channels := SafeCounter{v: 0}

	for key := range alpha_map {
		path := []string{start_node, key}
		num_channels.Inc()
		go explore_path(path, ConvMatrix, path_channel, num_channels)
	}

	best_path := make([]string, len(currencies)+2)
	best_ret := 0.0

	for i := 0; i < num_channels.Value(); i++ {
		res := <-path_channel
		ret := traverseGraph(res, ConvMatrix)
		if ret > best_ret {
			best_path = res
			best_ret = ret
		}

	}
	return best_path, best_ret
}
