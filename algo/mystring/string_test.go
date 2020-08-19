package mystring_test

import (
	"testing"

	"github.com/sKudryashov/algo/mystring"
)

func TestStringWindow(t *testing.T) {
	strOne := "adbokskdjjdkfjdkjf"
	// an := "bko"
	strTwo := []rune(strOne)
	for i := 0; i < len(strOne); i++ {
		substr := strTwo[1:3]
		println(" str two ", string(strTwo[i]))
		println("substr ", string(substr))
		println("strTwo", string(strTwo[:i]))
	}
}

func TestMatchingPairs(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				s: "abcd",
				t: "adcb",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mystring.MatchingPairs(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("MatchingPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
