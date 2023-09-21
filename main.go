package main

import (
	"fmt"
	utils "fxIneff/src/api"

	"fxIneff/src/graph"
	// matrix "fxIneff/src/conversion_matrix"
)

func main() {
	fmt.Println("running main package")

	vertices := utils.Generate_nodes()
	graph.TraversalAlgorithm(vertices, "eur")
}
