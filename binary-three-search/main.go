package main

import (
	"fmt"
)

var node_count int

type node struct {
	key   int
	left  *node
	right *node
}

func (n *node) insert(key int) {
	if n.key == 0 {
		n.key = key

		return
	}

	if n.key < key {
		if n.right == nil {
			n.right = &node{key: key}
		} else {
			n.right.insert(key)
		}

		return
	}

	if n.key > key {
		if n.left == nil {
			n.left = &node{key: key}
		} else {
			n.left.insert(key)
		}
	}

	return
}

func (n *node) search(key int) bool {
	node_count++

	if n == nil {
		return false
	}

	if n.key == key {
		fmt.Println(n, n.left, n.right)
		return true
	}

	if key < n.key {
		return n.left.search(key)
	}

	if key > n.key {
		return n.right.search(key)
	}

	return false
}

func main() {
	keys := []int{52, 203, 19, 76, 150, 310, 7, 24, 88, 276, 456, 112, 86, 998, 5, 33}

	tree := &node{}

	for _, key := range keys {
		tree.insert(key)
	}

	tree.search(203)
}
