// package algorithm
package main

import (
	"fmt"
	"fxIneff/src/parallel/matrix"
	"fxIneff/src/parallel/utils"	
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

	for key, value := range alpha_map {
		fmt.Printf("%s: %.16f\n", key, value)
	}

	return []string{"abc"}, 0.1
}

func main() {ParallelDFS("eur")}
