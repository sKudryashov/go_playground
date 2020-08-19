package matrix

import "testing"

func TestGetNumOfIslands(t *testing.T) {
	type args struct {
		islands [][]string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				islands: [][]string{
					{"1", "1", "0", "0", "0"},
					{"1", "1", "0", "0", "0"},
					{"0", "0", "1", "0", "0"},
					{"0", "0", "0", "1", "1"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetNumOfIslands(tt.args.islands)
		})
	}
}
