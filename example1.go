package main

import "fmt"
import "os"

// this is a comment

func main() {
	fmt.Println("My Name is Saleem")
	fmt.Println("Hello, World")
	fmt.Println("1 + 1 = ", 1 + 1)
	fmt.Println("0.1 + 2.1 = ", 0.1 + 2.1)
	fmt.Println(len("Hello, World"))
	fmt.Println("Hello, World"[1])
	fmt.Println("Hello," + " World")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(2*2*2*2 * 2*2*2*2 * 2 - 1)
	fmt.Println(32132 * 42452)
	fmt.Println((true && false) || (false && true) || !(false && false))
	os.Exit(0)
	// unreachable code below this line
}
