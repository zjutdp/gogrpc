package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values

// from the tree to the channel ch.

func Walk(t *tree.Tree, ch chan int) {

	if t != nil {

		ch <- t.Value

		Walk(t.Left, ch)

		Walk(t.Right, ch)

	}

	close(ch)

}

// Same determines whether the trees

// t1 and t2 contain the same values.

func Same(t1, t2 *tree.Tree) bool {

	ch1 := make(chan int)

	ch2 := make(chan int)

	go Walk(t1, ch1)

	go Walk(t2, ch2)

	for i := range ch1 {

		j, ok := <-ch2

		if i != j || !ok {

			fmt.Println(i, j)

			return false

		}

	}

	return true

}

func main() {

	t1 := tree.New(2)

	t2 := tree.New(2)

	fmt.Println(t1)

	fmt.Println(t2)

	fmt.Println(Same(t1, t2))

}
