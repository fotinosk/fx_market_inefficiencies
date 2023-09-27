package utils

import (
	"sort"
	"golang.org/x/exp/slices"

)

func SortMapByValue(inp map[string]float64) map[string]float64 {
    var entries []struct {
        Key   string
        Value float64
    }

    // Populate the slice with map entries
    for key, value := range inp {
        entries = append(entries, struct {
            Key   string
            Value float64
        }{key, value})
    }

    // Sort the slice by values
    sort.Slice(entries, func(i, j int) bool {
        return entries[i].Value < entries[j].Value
    })

    // Create a new map and populate it with sorted entries
    sortedMap := make(map[string]float64, len(entries))
    for _, entry := range entries {
        sortedMap[entry.Key] = entry.Value
    }

    return sortedMap
}


func RemoveSliceElements(list1 []string, list2 []string) []string {
	// remove the elements of list2 from list1
	var rslice = make([]string, len(list1))


	for _, element := range(list1) {
		if !slices.Contains(list2, element) {
			rslice = append(rslice, element)
		}
	}
	return rslice
}