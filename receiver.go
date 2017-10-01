package main
import "fmt"

type Decimal struct {
	first float64
}


// func (receiver spec)  func_name(args...) return_type { body }

func (x Decimal) out() float64 {
	return x.first
}

func (x Decimal) fundoo() float64 {
	return 20.0
}


func main() {
	var start Decimal
	start.first = 10.8
	
	// here is the receiver magic
	show := start.out()
	fmt.Println(show)
	fmt.Println(start.fundoo())
}


/*

class MyClass():
    def fundoo(self):
        return 20.0


def MyClasss::fundoo(self, a1):
        return 20.0 * a1


start = MyClass()

start.fundoo()

*/
