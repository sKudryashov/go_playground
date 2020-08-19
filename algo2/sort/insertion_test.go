package sort_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"

	mysort "github.com/sKudryashov/algo/sort"
)

func TestInsertionSort(t *testing.T) {
	sortedArr := mysort.LaunchInsertionSort()
	spew.Dump(sortedArr)
}
