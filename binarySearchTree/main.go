package binarySearchTree

import (
	"fmt"
	"golang.org/x/tour/tree"
	"math/rand"
	"time"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	return false
}

func ConcurrentBinaryTree() {
	rand.Seed(time.Now().UnixNano())

	ch1 := make(chan int)
	t1 := tree.New(10)
	fmt.Println("Tree 1: ", t1)
	go Walk(t1, ch1)
	fmt.Println("Walking tree 1:")
	for i := 0; i < 10; i++ {
		v := <-ch1
		fmt.Println(v)
	}
}
