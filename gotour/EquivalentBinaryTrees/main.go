package main

import (
	"fmt"
	"reflect"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		close(ch)
		return
	}
	ch <- t.Value

	if t.Left != nil {
		Walk(t.Left, ch)
	}
	if t.Right != nil {
		Walk(t.Right, ch)
	}

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 100)
	go Walk(t1, ch1)
	ch2 := make(chan int, 100)
	go Walk(t2, ch2)

	res1, res2 := make([]int, 0), make([]int, 0)
	exit := false
	for !exit {
		select {
		case v := <-ch1:
			fmt.Println("ch1", v)
			res1 = append(res1, v)
		default:
			exit = true
			break
		}
	}

	exit = false

	for !exit {
		select {
		case v := <-ch2:
			fmt.Println("ch2", v)
			res2 = append(res2, v)
		default:
			exit = true
			break
		}
	}

	return reflect.DeepEqual(res1, res2)
}

func main() {
	// fmt.Println("1", Same(tree.New(1), tree.New(1)))
	// fmt.Println("2", Same(tree.New(1), tree.New(2)))

	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for v := range ch {
		fmt.Println("main", v)
	}
}
