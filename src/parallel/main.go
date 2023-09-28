package main

import (
	"fmt"
	"fxIneff/src/parallel/algorithm"
	"time"
)

func main() {
	start := time.Now()
	p, r := algorithm.ParallelDFS("eur")
	elapsed := time.Since(start)
	fmt.Println(p, r, elapsed)
}
