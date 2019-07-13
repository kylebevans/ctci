package main

import (
	"strings"
	"fmt"
)

func Parens(p int) []string {
// ()
// ()() (())
// (()()) ()()() ((())) ()(()) (())()
// ((()())) (()()()) (((()))) (()(())) ((())()) ()(()()) ()()()() ()((())) ()()(()) ()(())() (()())() ((()))() (())()()
	var combos []string
	if p == 0 {
		return combos
	}
	if p == 1 {
		combos = append(combos, "()")
		return combos
	}
	pMinusCombos := Parens(p-1)
	combosDone := make(map[string]bool)
	for _, v := range pMinusCombos {
		comboCandidate1 := strings.Join([]string{"()", v}, "")
		if !combosDone[comboCandidate1] {
			combos = append(combos, comboCandidate1)
			combosDone[comboCandidate1] = true
		}
		
		comboCandidate2 := strings.Join([]string{v, "()"}, "")
		if !combosDone[comboCandidate2] {
			combos = append(combos, comboCandidate2)
			combosDone[comboCandidate2] = true
		}

		comboCandidate3 := strings.Join([]string{"(", v, ")"}, "")
		if !combosDone[comboCandidate3] {
			combos = append(combos, comboCandidate3)
			combosDone[comboCandidate3] = true
		}
	}
	return combos
	
}

func main() {
	parenCombos := Parens(4)
	fmt.Println(parenCombos)
	return
}