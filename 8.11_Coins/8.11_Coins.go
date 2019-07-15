package main


func NumCoinCombos(n int, currentCoin int, currentCoins [4]int) ([][4]int, int) {
// while current number of current denom is <= n % current denom
// base case is: add a coin, total = n
// add a kind of coin and recursively call NumCoinCombo with that set

	var combos [][4]int
	denoms := [4]int{25, 10, 5, 1}
	total := 25 * currentCoins[0] + 10 * currentCoins[1] + 5 * currentCoins[2] + currentCoins[3]
	if currentCoin == 3 {
		currentCoins[3] = n - total
		combos = append(combos, currentCoins)
		return combos, len(combos)
	}
	for i := 0; i <= (n - total) / denoms[currentCoin]; i++ {
		var currentCombos [][4]int
		currentCoins[currentCoin] = i
		total := 25 * currentCoins[0] + 10 * currentCoins[1] + 5 * currentCoins[2] + currentCoins[3]
		if total == n {
			combos = append(combos, currentCoins)
			return combos, len(combos)
		}
		currentCombos, _ = NumCoinCombos(n, currentCoin+1, currentCoins)
		combos = append(combos, currentCombos...)
			
	}
	return combos, len(combos)
}
