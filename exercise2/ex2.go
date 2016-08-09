package main

import (
   "fmt"
)

func main() {
  try_pointers()
  try_struct()
  try_array()
  try_slice()
}

func try_slice() {
   s1 := []int {1, 2, 3, 4, 5, 6, }
   s2 := make([]int, 5, 10)
   for i := 0; i < len(s2); i++ {
      s2[i] = i*i
   }
   fmt.Println("s1 =", s1)
   fmt.Println("s2 =", s2)
   
   var s_nil []int	// A nil slice
   sa := append(s_nil, 1,2,3,6)
   fmt.Println("sa =", sa, "len=", len(sa), "cap=", cap(sa))
   for i, v := range sa {
      fmt.Printf("sa[%d] = %d\n", i, v)
   }
   for i := range sa {
      sa[i] = 1 << uint(i)
   }
   for _, v := range sa {
      fmt.Printf(" %d", v)
   }
   fmt.Println()
   pixels := getPixels2(10,20)
   for i, v := range pixels {
      fmt.Println(i, v)
   }
}

func getPixels1(w, h int) [][]uint8 {
   var pix [][]uint8
   for x := 0; x < w; x++ {
      pix = append(pix, make([]uint8, h))
      for y := 0; y < h; y++ {
         pix[x][y] = uint8(x * y)
      }
   }
   return pix
}

func getPixels2(w, h int) [][]uint8 {
   pix := make([][]uint8, w)
   for x := 0; x < w; x++ {
      pix[x] = make([]uint8, h)
      for y := 0; y < h; y++ {
         pix[x][y] = uint8(x + y)
      }
   }
   return pix
}

func try_pointers() {
   i := 100
   pi := &i
   var p *int
   p = &i
   fmt.Printf("*pi = %v\n", *p)
   *pi = 200
   fmt.Printf("i = %v\n", i)   
}

type Point2d struct {
   x, y float64
}

func try_struct() {
   p1 := Point2d { 10.0, 20.0 }
   fmt.Printf("p1(%v, %v)\n", p1.x, p1.y)

   var (
      p2 Point2d
      p3 = &Point2d { y:333, x:211 }
   )
   p2.x = p1.x
   p2.y = p1.y + 13.33
   fmt.Printf("p2(%v, %v)\n", p2.x, p2.y)
   fmt.Printf("p3(%v, %v)\n", p3.x, p3.y)
}

func try_array() {
   var strings [2] string
   strings[0] = "String 0"
   strings[1] = "String 1"
   fmt.Println(strings[0], "and", strings[1])
   
   integers := [5] int { 10, 20, 30 }
   fmt.Println("integers[...] =", integers, "array len =", len(integers), integers[2:4])
   
}
