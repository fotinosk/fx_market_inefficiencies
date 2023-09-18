package main

import (
	"fmt"
	utils "fxIneff/src/api"

	// "fxIneff/src/graph"
	matrix "fxIneff/src/conversion_matrix"
)

func main() {
	fmt.Println("running main package")

	vertices := utils.Generate_nodes()
	var matrix matrix.ConversionMatrix

	matrix.PopulateConversionMatrix(vertices)

	fx1, _ := matrix.GetExchangeRate("eur", "usd")
	fx2, _ := matrix.GetExchangeRate("usd", "eur")

	fx3, _ := matrix.GetExchangeRate("gbp", "usd")
	fx4, _ := matrix.GetExchangeRate("usd", "gbp")

	fx5, _ := matrix.GetExchangeRate("eur", "gbp")
	fx6, _ := matrix.GetExchangeRate("gbp", "eur")
	
	fmt.Println(fx1)
	fmt.Println(fx2)
	fmt.Println(fx3)
	fmt.Println(fx4)
	fmt.Println(fx5)
	fmt.Println(fx6)
}
