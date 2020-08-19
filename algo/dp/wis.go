package dp

import (
	"github.com/davecgh/go-spew/spew"
)

// Graph1
// (8)----(4)----(4)----(5)----(2)----(7)----(1)
// 16 obviously the answer
func LaunchWIS() {
	//setOne = (8)----(4)----(4)----(5)----(2)----(7)
	//setTwo = (8)----(4)----(4)----(5)----(2) + (1)
	srcWeights := []int{
		0, // prepend with 0 val to shift the index
		8,
		4,
		4,
		5,
		2,
		7,
		1,
	}

	calcWis(srcWeights)
}

func calcWis(srcWeights []int) {
	elementsNum := len(srcWeights)
	a := make([]int, elementsNum, elementsNum)
	a[0] = 0
	a[1] = srcWeights[1]
	for i := 2; i < elementsNum; i++ {
		firstInvI := i - 1
		secondInvI := i - 2
		firstInv := a[firstInvI]
		secondInv := a[secondInvI]
		secondInvExpr := secondInv + srcWeights[i]
		mx := max(firstInv, secondInvExpr)
		a[i] = mx
	}
	spew.Dump("a content ", a)
	// for i <= 1 {
	// }
}
