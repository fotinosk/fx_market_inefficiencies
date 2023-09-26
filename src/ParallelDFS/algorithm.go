package parallelDfs

import (
	"fmt"
	utils "fxIneff/src/api"
	matrix "fxIneff/src/conversion_matrix"
	"golang.org/x/exp/slices"

)

// const NUM_CONCURRENT_PATHS = 4  // no. of go routines 


func getGreedyMultiple(
	start_node string,
	current_node string,
	already_visited []string,
	currency_matrix matrix.ConversionMatrix,
) []string {
	// order the possible next steps by alpha
	
	possibilities := utils.RemoveSliceElements(currency_matrix.GetCurrencies(), already_visited)
	next_step_map := map[]
}

func TraverseGraphConcurrently(vertices map[string]map[string]float64) ([]string, float64) {

}