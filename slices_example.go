package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
   image := make([][]uint8, dy)
   for i:=0; i < dy; i++ {
     image[i] = make([]uint8, dx)
   }
   for x := 0; x < dx; x++ {
     for y := 0; y < dy; y++ {
	    // v := x*y
	    v := (x+y) / 2.0
	    // v := x^y
        image[y][x] = uint8(v)
     }
   }

   return image
}

func main() {
	pic.Show(Pic)
}

