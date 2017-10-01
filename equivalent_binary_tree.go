package main

import (
"golang.org/x/tour/tree"
"fmt"
)

func doWalk(t *tree.Tree, ch chan int) {
   if t != nil {
     doWalk(t.Left, ch)
     ch <- t.Value
     doWalk(t.Right, ch)
   }
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
   // fmt.Printf("%v", t)
   doWalk(t, ch)
   close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
  ch1 := make(chan int)
  ch2 := make(chan int)
  go Walk(t1, ch1)
  go Walk(t2, ch2)
  for x := range ch1 {
    y := <- ch2
	// fmt.Printf("%v %v\n", x, y)
	if(x != y) {
	   return false
	}
  }
  return true
}

func main() {
  ch := make(chan int)
  go Walk(tree.New(1), ch)
  for x := range ch {
    fmt.Printf("%v ", x)
  }
    fmt.Printf("\n")
  s1 := Same(tree.New(1), tree.New(1))
  s2 := Same(tree.New(1), tree.New(2))
  fmt.Printf("%v\n", s1)
  fmt.Printf("%v\n", s2)
}
