package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
    parts := strings.Split(s, " ")
	m := make(map[string]int)
	for _, v := range parts {
	   _, ok := m[v]
	   switch {
	      case ok:
		     m[v] = m[v] + 1
		  default:
		     m[v] = 1
	   }
	}
	return m
}

func main() {
//  f("I am learning Go!") = 
//   map[string]int{"learning":1, "Go!":1, "I":1, "am":1}
// PASS
//  f("The quick brown fox jumped over the lazy dog.") = 
//   map[string]int{"jumped":1, "over":1, "the":1, "lazy":1, "The":1, "quick":1, "brown":1, "fox":1, "dog.":1}
// PASS
//  f("I ate a donut. Then I ate another donut.") = 
//   map[string]int{"Then":1, "another":1, "I":2, "ate":2, "a":1, "donut.":2}
// PASS
//  f("A man a plan a canal panama.") = 

	wc.Test(WordCount)
}
