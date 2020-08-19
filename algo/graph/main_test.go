package graph_test

import "testing"

type GraphTest struct {
	set string
}

func (GraphTest) GetInt() int {
	return 1
}

func NewGraphTestEsc() *GraphTest {
	return &GraphTest{
		set: "testset2",
	}
}

func NewGraphTest() int {
	gp := GraphTest{
		set: "testset",
	}
	gt := new(GraphTest)
	gt.GetInt()
	gg := NewGraphTestEsc()
	gg.GetInt()
	gg.set = "int"
	return gp.GetInt()
}

type GraphTestInt int

func NewGraphTestInt() GraphTestInt {
	return 3
}

type GraphSlice []int

func NewGraphTestSlice() int {
	gs := GraphSlice{}
	gs = append(gs, []int{1, 2, 3}...)
	return gs[0]
}

func BenchmarkTypes(b *testing.B) {
	b.Run("benchmark type struct", func(bb *testing.B) {
		bb.ReportAllocs()
		for i := 0; i < bb.N; i++ {
			NewGraphTest()
		}
	})
	b.Run("benchmark type int", func(bb *testing.B) {
		bb.ReportAllocs()
		for i := 0; i < bb.N; i++ {
			NewGraphTestInt()
		}
	})
	b.Run("benchmark type slice", func(bb *testing.B) {
		bb.ReportAllocs()
		for i := 0; i < bb.N; i++ {
			NewGraphTestSlice()
		}
	})
}
