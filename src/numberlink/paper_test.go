package main

import (
	"testing"
)

// Run benchmarks with 'go test numberlink -bench=".*"'

var gobacktests = []struct {
	ptr  []int
	sets [][]int
}{
	{[]int{1, 2, 3, 4}, [][]int{{0, 1}, {1, 1}, {2, 1}, {3, 1}}},
}
var paper = NewPaper(2, 2, []rune{0, 0, 0, 0})

func BenchmarkGoback(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range gobacktests {
			for j := range tt.sets {
				paper.setnrem(tt.ptr, tt.sets[j][0], tt.sets[j][1])
			}
			paper.goBack(0)
		}
	}
}
