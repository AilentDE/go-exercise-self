package main

import (
	"fmt"
	"sort"
)

func main() {
	OriSlice := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		OriSlice = append(OriSlice, i)
	}

	newSlice := OriSlice[:5] // [0 1 2 3 4 5 6 7 8 9]

	newSlice[0] = 100 // [100 1 2 3 4 5 6 7 8 9]
	sort.Ints(newSlice) // [1 2 3 4 100 5 6 7 8 9]

	fmt.Println(
		"OriSlice:", OriSlice,
	)
}