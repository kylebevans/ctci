package main

func MakeChange(n int, coins []int, currentCoin int, storedWays map[int]map[int]int) int {
	if _, ok := storedWays[n][coins[currentCoin]]; ok {
		return storedWays[n][coins[currentCoin]]
	}
	if currentCoin >= len(coins) - 1 {
		return 1
	}
	var count int
	for i := 0 ; i * coins[currentCoin] <= n; i++ {
		currentAmount := n - i * coins[currentCoin]
		count += MakeChange(currentAmount, coins, currentCoin + 1, storedWays) 
	}
	storedWays[n][coins[currentCoin]] = count
	return count
}
