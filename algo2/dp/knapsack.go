package dp

type knapsack struct {
	capacity int
}

type items []item

type item struct {
	value int
	size  int
}

func LaunchKnapsack() {
	p := knapsackProblem()
	println("knapsackProblem solution here is ", p)
}

func knapsackReconstruction() {}

func knapsackProblem() int {
	A := make([][]int, 6, 6)
	knap := knapsack{
		capacity: 6,
	}
	// C := 6
	// n := 4

	things := items{
		{}, // zero element
		{
			size:  4,
			value: 3,
		}, {
			size:  3,
			value: 2,
		}, {
			size:  2,
			value: 4,
		}, {
			size:  3,
			value: 4,
		},
	}
	// subs := make([]int, 6, 6)
	for i := 0; i < knap.capacity; i++ {
		subs := make([]int, 7, 7)
		A[i] = subs
		// A[0][i] = 0
	}
	thingsAmount := len(things)
	for i := 1; i < thingsAmount; i++ {
		for c := 0; c <= knap.capacity; c++ {
			firstSubProblem := A[i-1][c] // first subproblem
			if things[i].size > c {
				// if current thing is bigger than current space resulted in knapsack accretion, keep the previos item
				// in the knapsack (or none)
				A[i][c] = firstSubProblem
			} else {
				// if there is a room for new item - let's calculate the residual capacity after putting
				// this item into knapsack
				residualCapacity := c - things[i].size
				reservedValue := A[i-1][residualCapacity]
				// let's take the value of previous and prospective items and compare it to
				// the previous case
				secondSubProblem := reservedValue + things[i].value
				bestValue := max(firstSubProblem, secondSubProblem)
				A[i][c] = bestValue
			}
		}
	}
	return 1
	// return A[len(things)][knap.capacity]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//TODO: backtrack knapsack
