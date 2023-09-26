package matrix

import (
	"fmt"
	"fxIneff/src/parallel/api"
	"sort"
)

type Matrix struct {
	// convention: if I want to exchange A to B, then I will find the row representing A and find the exchange rate for B
	currency_names []string
	conversions    [][]float64
}

func (mat *Matrix) PopulateMatrix() {
	// sort the keys for consistency
	vertices := api.GenerateConversionGraph()
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

func (mat Matrix) GetExchangeRate(from string, to string) (float64, error) {
	ind_from := -1
	ind_to := -1

	for ind, val := range mat.currency_names {
		if val == from {
			ind_from = ind
		}
		if val == to {
			ind_to = ind
		}
	}

	if ind_from == -1 || ind_to == -1 {
		return 0.0, fmt.Errorf("provided currency is not in exchange list")
	}
	return mat.conversions[ind_from][ind_to], nil
}

func (mat Matrix) GetAlpha(base_currency string, currennt_currency string, next_currency string) float64 {
	// returns current -> next -> base
	current_to_next, _ := mat.GetExchangeRate(currennt_currency, next_currency)
	next_to_base, _ := mat.GetExchangeRate(next_currency, base_currency)

	return current_to_next * next_to_base
}

func (mat Matrix) GetCurrencies() []string {
	return mat.currency_names
}