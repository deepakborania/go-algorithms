package connectedComponents

import (
	"testing"
)

func TestCountComponents(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test2", args{5, [][]int{{0, 1}, {1, 2}, {3, 4}}}, 2},
		{"Test1", args{5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}}, 1},
		{"Test4", args{8, [][]int{{0, 1}, {2, 3}, {4, 5}, {6, 7}}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountComponents(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("CountComponents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountComponentsIterative(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test2", args{5, [][]int{{0, 1}, {1, 2}, {3, 4}}}, 2},
		{"Test1", args{5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}}, 1},
		{"Test4", args{8, [][]int{{0, 1}, {2, 3}, {4, 5}, {6, 7}}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountComponentsIterative(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("CountComponentsIterative() = %v, want %v", got, tt.want)
			}
		})
	}
}
