package matrix

import (
	"fmt"
	"sort"
)

type ConversionMatrix struct {
	// convention: if I want to exchange A to B, then I will find the row representing A and find the exchange rate for B
	currency_names []string
	conversions    [][]float64
}

func (mat *ConversionMatrix) PopulateConversionMatrix(vertices map[string]map[string]float64) {
	// sort the keys for consistency
	names := make([]string, len(vertices))

	i := 0
	for k := range vertices {
		names[i] = k
		i++
	}
	sort.Strings(names)
	mat.currency_names = names // names are now sorted

	for _, row_name := range names {
		rates := vertices[row_name]
		vals := make([]float64, len(rates))

		j := 0
		for _, name2 := range names {
			vals[j] = rates[name2]
			j++
		}
		mat.conversions = append(mat.conversions, vals)
	}
}

func (mat *ConversionMatrix) GetExchangeRate(from string, to string) (float64, error) {
	ind_from := -1
	ind_to := -1

	for ind, val := range mat.currency_names {
		if val == from {
			ind_from = ind
		} else if val == to {
			ind_to = ind
		}
	}

	if ind_from == -1 || ind_to == -1 {
		return 0.0, fmt.Errorf("provided currency is not in exchange list")
	}
	return mat.conversions[ind_from][ind_to], nil
}
