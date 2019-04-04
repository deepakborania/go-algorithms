package unionfind

import "testing"

func TestUnionFind_Count(t *testing.T) {
	uf := New(10)
	uf.Union(0, 1)
	uf.Union(1, 2)
	got := uf.Count()
	if got != 8 {
		t.Errorf("Count() = %d, Expected = 8", got)
	}
	uf.Union(9, 8)
	got = uf.Count()
	if got != 7 {
		t.Errorf("Count() = %d, Expected = 7", got)
	}
}
