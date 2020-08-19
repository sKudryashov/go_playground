package sort_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"

	mysort "github.com/sKudryashov/algo/sort"
)

func TestMergeSort(t *testing.T) {
	arr := []int{0, 34, 7, 88 ,99 ,63}
	sortedArr := mysort.MergeSort(arr)
	spew.Dump(sortedArr)
}