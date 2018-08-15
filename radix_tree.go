package main

import (
	"fmt"
	"io"
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

func (rt *radixTree) Write(w io.Writer) error {
	_, err := fmt.Fprint(w, "{")
	if err != nil {
		return err
	}

	for index, root := range rt.roots {
		err = root.Write(w)
		if err != nil {
			return err
		}

		if len(rt.roots)-1 > index {
			_, err = fmt.Fprint(w, ",")
			if err != nil {
				return err
			}
		}
	}

	_, err = fmt.Fprintf(w, "}")
	return err
}
