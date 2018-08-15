package main

import "strings"

type radixNode struct {
	key      string
	children []*radixNode
}

func newRadixNode(value string) *radixNode {
	return &radixNode{
		key:      value,
		children: make([]*radixNode, 0),
	}
}

func (rn *radixNode) Add(value string) bool {
	if value == "" {
		return false
	}
	if rn.key == value {
		return true
	}

	keyContainsValue := strings.HasPrefix(rn.key, value)
	valueContainsKey := strings.HasPrefix(value, rn.key)
	if keyContainsValue == false && valueContainsKey == false {
		return false
	}

	splitIndex := strings.Index(rn.key, value)
}
