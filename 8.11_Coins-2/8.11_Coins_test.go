package main

import (
	"testing"
)

func TestMakeChange(t *testing.T) {
	comboMap := make(map[int]map[int]int)
	for i := 0; i <= 50; i++ {
		comboMap[i] = make(map[int]int)
	}
	combos := MakeChange(50, []int{25, 10, 5, 1}, 0, comboMap)
	if combos != 49 {
		t.Errorf("NumCoinCombos for 50 should be 49, but it's: %v\n", combos)
	}
	return
}
