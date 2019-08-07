package main

import (
	"fmt"
	"testing"
)

func TestGetMaxStackHeight(t *testing.T) {
	boxset := []Box{{1, 1, 1}, {2, 2, 2}, {3, 4, 4}, {4, 5, 5}, {5, 6, 6}, {1, 10, 10}, {2, 3, 3}, {3, 3, 3}, {4, 3, 5}, {5, 5, 5}, {4, 4, 4}, {4, 3, 1}}
	stack, height := GetMaxStackHeight(boxset)
	if height != 19 {
		t.Errorf("Stack height should be 15, but is %v\n", height)
		fmt.Printf("Stack: %v\n", stack)
	}
	fmt.Printf("Stack: %v\n", stack)
}