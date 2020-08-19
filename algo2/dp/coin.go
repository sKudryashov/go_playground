package dp

import (
	"fmt"
	"strconv"
)

// LaunchCoin coin launcher
func LaunchCoinMax() {
	targetValue := 5
	denominations := []int{1, 2, 3}
	coin := newCoin(denominations)
	coin.subproblem(denominations, targetValue)
}

type coin struct {
	cache         map[string]int
	denominations []int
}

func newCoin(denom []int) *coin {
	return &coin{
		cache:         make(map[string]int, 10),
		denominations: denom,
	}
}

// public class Solution {
// 	public int coinChange(int[] coins, int amount) {
// 	  int max = amount + 1;
// 	  int[] dp = new int[amount + 1];
// 	  Arrays.fill(dp, max);
// 	  dp[0] = 0;
// 	  for (int i = 1; i <= amount; i++) {
// 		for (int j = 0; j < coins.length; j++) {
// 		  if (coins[j] <= i) {
// 			dp[i] = Math.min(dp[i], dp[i - coins[j]] + 1);
// 		  }
// 		}
// 	  }
// 	  return dp[amount] > amount ? -1 : dp[amount];
// 	}
//   }

func (c *coin) subproblem(denominations []int, target int) int {
	table := make([][]int, target+1, target+1)
	// tableCoinsNeeded := make([]int, 0, target)
	for i := 0; i <= target; i++ {
		tableAxisX := make([]int, len(denominations), len(denominations))
		table[i] = tableAxisX
	}
	// tableSumTrack := make([]int, target+1, target+1)
	// coinsAmount := make(map[int]int, target)

	for amount := 1; amount <= target; amount++ {
		for d := 0; d < len(denominations); d++ {
			// var prevSum int
			println("state ", amount, ",", d+1)
			currentCoin := denominations[d]
			// if a == currentCoin {
			// 	// state 2,2; 3,3; 4,4
			// 	table[a][d] = currentCoin
			// }
			if amount >= currentCoin {
				// state 2,1 3,1
				//for case 3,1 we have 2 + 1
				// amount = 3, current coin = 2 prev coin = 1
				// if d == 0 {
				// 	prevSum = 0
				// } else {
				// 	prevSum = table[amount][d-1]
				// }
				residualCapacity := amount - currentCoin
				reservedValue := table[amount-1][residualCapacity] + currentCoin
				println("table[amount-1][d] ", table[amount-1][d])
				table[amount][d] = max(currentCoin, reservedValue)
			}
		}
	}
	return 0
}

func (c *coin) setCache(i, t, val int) {
	istr := strconv.Itoa(i)
	tstr := strconv.Itoa(i)
	key := fmt.Sprintf("%s#%s", istr, tstr)
	c.cache[key] = val
}

func (c *coin) checkCache(i, t int) (int, bool) {
	istr := strconv.Itoa(i)
	tstr := strconv.Itoa(i)
	key := fmt.Sprintf("%s#%s", istr, tstr)
	val, ok := c.cache[key]
	return val, ok
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
