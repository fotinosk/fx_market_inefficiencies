package utils

import "sort"

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

func SortMapByValues(inputMap *map[string]float64) {
	// Create a slice to store the map entries
	var entries []struct {
		Key   string
		Value float64
	}

	// Populate the slice with map entries
	for key, value := range *inputMap {
		entries = append(entries, struct {
			Key   string
			Value float64
		}{key, value})
	}

	// Sort the slice by values
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Value < entries[j].Value
	})

	// Clear the original map
	for key := range *inputMap {
		delete(*inputMap, key)
	}

	// Populate the original map with sorted entries
	for _, entry := range entries {
		(*inputMap)[entry.Key] = entry.Value
	}
}