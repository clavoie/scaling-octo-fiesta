package main

type radixTree struct {
	roots []*radixNode
}

func newRadixTree() *radixTree {
	return &radixTree{
		roots: make([]*radixNode, 0, 16),
	}
}

func (rt *radixTree) Add(value string) {

}
