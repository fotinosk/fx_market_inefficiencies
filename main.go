package main

import (
	"fmt"
	// "fxIneff/src/api"
	"fxIneff/src/graph"
)

func main() {
	fmt.Println("running main package")
	path := []string{"eur", "usd", "rub", "eur"}
	ret, err := graph.TraverseGraph(path)
	if err == nil {
		fmt.Printf("Traversal successful with returns of: %f", ret)
	} else {
		fmt.Println(err)
	}
}
