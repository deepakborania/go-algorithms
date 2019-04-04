package unionfind

// Adapted from https://algs4.cs.princeton.edu/15uf/UF.java.html

// UnionFind represents the union-find/disjoint-sets data type.
// No path compression/halving for ease of understanding. Easy to implement yourself.
type UnionFind struct {
	parent []int //parent[i]=parent of i
	rank   []int //rank[i] = rank of tree rooted  at i(never more than 31)
	count  int   //Connected components count
}

func New(n int) *UnionFind {
	uf := UnionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
		count:  n,
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.rank[i] = 0
	}
	return &uf
}

// Union merges components containing p and q
func (uf *UnionFind) Union(p, q int) {
	rootP := uf.Find(p)
	rootQ := uf.Find(q)
	if rootP == rootQ {
		return
	}
	if uf.rank[rootP] < uf.rank[rootQ] {
		uf.parent[rootP] = rootQ
	} else if uf.rank[rootQ] < uf.rank[rootP] {
		uf.parent[rootQ] = rootP
	} else {
		uf.parent[rootQ] = rootP
		uf.rank[rootP]++
	}
	uf.count--
}

// Find returns the component identifier for the component containing the element
func (uf UnionFind) Find(p int) int {
	uf.validate(p)
	for p != uf.parent[p] {
		p = uf.parent[p]
	}
	return p
}

// Connected checks if two elements are in same component
func (uf UnionFind) Connected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

// Count returns the number of components
func (uf UnionFind) Count() int {
	return uf.count
}

func (uf UnionFind) validate(p int) {
	if p < 0 || p >= len(uf.parent) {
		panic("invalid request")
	}
}
