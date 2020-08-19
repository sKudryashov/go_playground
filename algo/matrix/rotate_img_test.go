package matrix

import "testing"

func Test_rotate(t *testing.T) {
	type args struct {
		mtrx [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				mtrx: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.mtrx)
		})
	}
}
