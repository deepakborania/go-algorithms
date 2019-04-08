package main

import "testing"

func TestSolveKnapsackRecursive(t *testing.T) {
	type args struct {
		profits  []int
		weights  []int
		capacity int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case1", args{[]int{4, 5, 3, 7}, []int{2, 3, 1, 4}, 7}, 14},
		{"Case2", args{[]int{1, 6, 10, 16}, []int{1, 2, 3, 5}, 6}, 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolveKnapsackRecursive(tt.args.profits, tt.args.weights, tt.args.capacity); got != tt.want {
				t.Errorf("SolveKnapsackRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSolveKnapsackRecursive(b *testing.B) {
	p := []int{4, 5, 3, 7, 9, 2, 8, 12, 7, 1, 2}
	w := []int{2, 3, 1, 4, 4, 8, 2, 7, 3, 8, 3, 3}
	for i := 0; i < b.N; i++ {
		SolveKnapsackRecursive(p, w, 10)
	}
}

func BenchmarkSolveKnapsackMemoized(b *testing.B) {
	p := []int{4, 5, 3, 7, 9, 2, 8, 12, 7, 1, 2}
	w := []int{2, 3, 1, 4, 4, 8, 2, 7, 3, 8, 3, 3}
	for i := 0; i < b.N; i++ {
		SolveKnapsackMemoized(p, w, 10)
	}
}
func BenchmarkSolveKnapsackTabulated(b *testing.B) {
	p := []int{4, 5, 3, 7, 9, 2, 8, 12, 7, 1, 2}
	w := []int{2, 3, 1, 4, 4, 8, 2, 7, 3, 8, 3, 3}
	for i := 0; i < b.N; i++ {
		solveKnapsackTabulated(p, w, 10)
	}
}
