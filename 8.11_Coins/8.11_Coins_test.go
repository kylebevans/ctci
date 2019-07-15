package main

import (
	"testing"
	"fmt"
)

func TestNumCoinCombos(t *testing.T) {
	combos, numCombos := NumCoinCombos(50, 0, [4]int{0,0,0,0})
	if numCombos != 49 {
		t.Errorf("NumCoinCombos for 10 should be 4, but it's: %v\n", numCombos)
		for _, v := range combos {
			fmt.Println(v)
		}
	}
	return
}