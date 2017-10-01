package main

import (
	"fmt"
)

func Sqrt(x float64) (float64, error) {

  if x < 0 {
    e := ErrNegativeSqrt(x)
    return x, e
  }

  zn := 1.0
  for i := 0; i < 10; i++ {
     z := zn - (zn*zn - x) / (2*zn)
	 zn = z
  }
  return zn, nil
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
