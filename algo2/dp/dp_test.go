package dp_test

import (
	"testing"

	"github.com/sKudryashov/algo/dp"
)

func TestCoinComb(t *testing.T) {
	dp.LaunchCoinComb()
}

func TestLaunchCoin(t *testing.T) {
	dp.LaunchCoinMax()
}

func TestKnapsack(t *testing.T) {
	dp.LaunchKnapsack()
}

func TestLaunchStair(t *testing.T) {
	dp.LaunchStair()
}

func TestLaunchWIS(t *testing.T) {
	dp.LaunchWIS()
}

func TestTrap(t *testing.T) {
	dp.LaunchTrap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
}
