package equal_bst

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Need this function so we can close the channel after recursive walk over tree is done (avoid deadlock)
func Walk(t *tree.Tree, ch chan int) {
	walkTree(t, ch)
	close(ch)
}

func walkTree(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walkTree(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walkTree(t.Right, ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for num1 := range ch1 {
		num2 := <-ch2
		fmt.Println(num1, num2)
		if num1 != num2 {
			return false
		}
	}

	return true
}

// Test by running: Same(tree.New(1), tree.New(1))
