package dp

func LaunchCoinComb() {
	amount := 12
	coins := []int{1, 2, 5}
	combs := calculate(amount, coins)
	println(" combs ", combs[len(combs)-1])
}

func calculate(amount int, coins []int) []int {
	combinations := make([]int, amount+1, amount+1)
	combinations[0] = 1
	// for a := 0; a <= amount; a++ {
	for _, coin := range coins {
		for a := 1; a < len(combinations); a++ {
			if a >= coin {
				residual := a - coin
				combinations[a] = combinations[a] + combinations[residual]
			}
		}
	}
	// }
	return combinations
}
