package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {

  zn := 1.0
  for i := 0; i < 10; i++ {
     z := zn - (zn*zn - x) / (2*zn)
	 zn = z
  }
  return zn
}

func main() {
	fmt.Println(Sqrt(1.5))
}
