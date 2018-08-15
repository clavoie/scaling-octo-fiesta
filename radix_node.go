package main

import (
	"fmt"
	"strings"
)

type radixNode struct {
	key      string
	children []*radixNode
	isTerm   bool
}

func newRadixNode(value string) *radixNode {
	return &radixNode{
		key:      value,
		children: make([]*radixNode, 0),
		isTerm:   false,
	}
}

func (rn *radixNode) Add(value string) bool {
	if value == "" {
		return false
	}
	if rn.key == value {
		return true
	}

	// fmt.Printf("'%v' :: '%v'\n", rn.key, value)
	if rn.key[0] != value[0] {
		return false
	}

	// simple case, one string is completely contained by another
	if strings.HasPrefix(rn.key, value) {
		newChild := newRadixNode(rn.key[len(value):])
		newChild.children = rn.children
		// replace this node in place
		rn.key = value
		rn.children = []*radixNode{newChild}
		rn.isTerm = true
		return true
	}

	if strings.HasPrefix(value, rn.key) {
		rn.addChild(value[len(rn.key):])
		rn.isTerm = true
		return true
	}

	// complex case, one string is partially contained by the other
	splitIndex := 0
	minStr := rn.key
	maxStr := value
	isMaxKey := false

	if len(rn.key) > len(value) {
		minStr = value
		maxStr = rn.key
		isMaxKey = true
	}

	for index := range minStr {
		splitIndex = index
		if minStr[index] != maxStr[index] {
			// splitIndex -= 1
			break
		}
	}

	// fmt.Printf("newKey: %v, childMin: %v, childMax: %v\n", minStr[:splitIndex], minStr[splitIndex:], maxStr[splitIndex:])
	newKey := minStr[:splitIndex]
	childMin := newRadixNode(minStr[splitIndex:])
	childMax := newRadixNode(maxStr[splitIndex:])

	if isMaxKey {
		childMax.children = rn.children
	} else {
		childMin.children = rn.children
	}

	rn.key = newKey
	rn.children = []*radixNode{childMin, childMax}
	rn.isTerm = false
	return true
}

func (rn *radixNode) addChild(value string) {
	for _, child := range rn.children {
		if child.Add(value) {
			return
		}
	}

	rn.children = append(rn.children, newRadixNode(value))
}

func (rn *radixNode) Print(indent int) {
	termStr := ""
	if rn.isTerm {
		termStr = "*"
	}

	fmt.Printf("%v%v%v\n", strings.Repeat(" ", indent), rn.key, termStr)

	for _, child := range rn.children {
		child.Print(indent + 2)
	}
}
