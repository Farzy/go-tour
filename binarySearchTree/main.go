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
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		v1 := <-ch1
		v2 := <-ch2
		if v1 != v2 {
			return false
		}
	}
	return true
}

func ConcurrentBinaryTree() {
	rand.Seed(time.Now().UnixNano())

	ch1 := make(chan int)
	t1 := tree.New(1)
	fmt.Println("Tree 1: ", t1)
	go Walk(t1, ch1)
	fmt.Println("Walking tree 1:")
	for i := 0; i < 10; i++ {
		v := <-ch1
		fmt.Println(v)
	}

	t2 := tree.New(1)
	fmt.Println("Tree 2: ", t2)

	fmt.Println("Compare t1 & t2:", Same(t1, t2))

	t3 := tree.New(2)
	fmt.Println("Tree 3: ", t3)

	fmt.Println("Compare t1 & t3:", Same(t1, t3))
}
