package utils

import (
	"golang.org/x/exp/slices"
)

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