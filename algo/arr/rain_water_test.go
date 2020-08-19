package arr

import "testing"

func TestTrap(t *testing.T) {
	tests := []struct {
		name      string
		container []int
		want      int
	}{
		{
			container: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:      6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Trap(tt.container); got != tt.want {
				t.Errorf("Trap() = %v, want %v", got, tt.want)
			}
		})
	}
}
