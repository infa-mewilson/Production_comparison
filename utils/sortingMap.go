package utils

import (
	"NEWGOLANG/config"
	"log"
)

func SortingMap(m map[string]config.Value) []string {

	//using a sorted slice of keys to return a map[string]int in key order.
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	log.Println("the only keys are", keys)
	return keys
}
