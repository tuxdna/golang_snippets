package main
import "fmt"
const a = 10
var name string = "Neo"

func main() {
	var x string = "Hello, World"
	fmt.Println(x)
	var y string
	y = "Hello, World"
	fmt.Println(y)

	z := y + " Wassup? " + f()

	fmt.Println(z)

	b := 20 + a
	fmt.Println(b)

	var input float64
	fmt.Println("Enter a number: ")
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Printf("Double of %v is %f\n", input, output)
}


func f() string {
	return name
}
