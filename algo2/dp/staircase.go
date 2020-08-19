package dp

func LaunchStair() {
	stairAmount := 2
	stairHeight := 4
	print("count of ways ", countWays(stairAmount, stairHeight+1))
}

func countWays(stairAmount, stairHeight int) int {
	res := make([]int, 10, 10)
	res[0] = 1
	res[1] = 1
	for stHght := 2; stHght < stairHeight; stHght++ {
		res[stHght] = 0
		// for stAm := 1; stAm <= stairAmount && stAm <= stHght; stAm++ { // just beatiful example
		for stAm := 1; stAm <= stairAmount; stAm++ {
			if stAm <= stHght {
				curStairCurStepDiff := stHght - stAm
				res[stHght] = res[stHght] + res[curStairCurStepDiff]
			}
		}
	}
	return res[stairHeight-1]
}
