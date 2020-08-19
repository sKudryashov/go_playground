package tree

import "testing"

func TestDeepestLeavesSum(t *testing.T) {
	tests := []struct {
		name string
		in0  []int
		want int
	}{
		{
			name: "unit",
			// assume zeroes here are nil leaves
			in0: []int{1, 2, 3, 4, 5, 0, 6, 7, 0, 0, 0, 0, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeepestLeavesSum(tt.in0); got != tt.want {
				t.Errorf("DeepestLeavesSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
