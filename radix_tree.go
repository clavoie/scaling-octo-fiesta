package main

import (
	"strings"
)

type radixTree struct {
	roots []*radixNode
}

func newRadixTree() *radixTree {
	return &radixTree{
		roots: make([]*radixNode, 0, 16),
	}
}

func (rt *radixTree) Add(value string) {
	value = strings.TrimSpace(value)
	if value == "" {
		return
	}

	for _, root := range rt.roots {
		if root.Add(value) {
			return
		}
	}

	rt.roots = append(rt.roots, newRadixNode(value))
}

func (rt *radixTree) Print() {
	for _, root := range rt.roots {
		root.Print(0)
	}
}
