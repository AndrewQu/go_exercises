package main

import (
   "fmt"
   "math"
)

func main() {
   try_passfunc()
   try_func_closure()
   try_method()
   try_interface()
}

func try_interface() {
   pCenter := Point2d {5.0, 5.0}
   var iAbs IAbs = pCenter
   fmt.Println("Abs of (5,5) is", iAbs.Abs())
   p := iAbs.(Point2d)
   fmt.Println("Access Point2d through iAbs (x,y) =", p.X, p.Y)
   check_interface_type(iAbs)
   
   iAbs = myFloat(-17)
   fmt.Printf("Abs of %v is %v\n", -17, iAbs.Abs())
   p, ok := iAbs.(Point2d)
   if !ok {
      fmt.Println("iAbs now does not have type Point2d")
   } else {
      fmt.Println("iAbs still has type Point2d")
   }
   check_interface_type(iAbs)
   check_interface_type(303)
}

func check_interface_type( i interface{} ) {
   switch v := i.(type) {
      case myFloat :
         fmt.Println("iAbs.(type) is myFloat")
      case int :
         fmt.Println("An int type, v=", v)
      case Point2d :
         fmt.Println("iAbs.(type) is Point2d, p=", v)
      default :
         fmt.Println("Not implemented iAbs.(type)=", v)
   }

}

type IAbs interface {
   Abs() float64
}

func (p Point2d) Abs() float64 {
   return math.Sqrt(p.X * p.X + p.Y * p.Y)
}

func (p Point2d) String() string {
   return fmt.Sprintf("(%.2f, %.2f)", p.X, p.Y)
}

type myFloat float64

func (f myFloat) Abs() float64 {
   if f < 0 { f = -f }
   return float64(f)
}

func try_method() {
   p0 := Point2d { 20, 20 }
   fmt.Printf("Dist from %v to (30,30) is %5.2f\n", p0, p0.Distance(Point2d{ 30, 30 }))
   fmt.Println("p0 =", p0)
}

type Point2d struct {
   X, Y float64
}

func (p Point2d) Distance(p2 Point2d) float64 {
   dx := p.X - p2.X
   dy := p.Y - p2.Y
   p.X += 100	// Changes the original point if p is passed as *Point2d
   return math.Sqrt(dx * dx + dy * dy)
}

func try_func_closure() {
   func_add, func_minus := summation(), summation()
   pos, neg := 0, 0
   for i := 0; i < 10; i++ {
      pos = func_add(i); neg = func_minus(-i)
   }
   fmt.Println("Sum of 1 to 9:", pos, " Sum of -1 to -9:", neg)
}

func summation() func(int) int {
   sum := 0
   return func(a int) int { sum += a; return sum }
}

func try_passfunc() {
   fmt.Println("200+333 =", arithmetic(add, 200, 333))
   fmt.Println("200*333 =", arithmetic(multiply, 200, 333))
   fmt.Println("2200/20 =", arithmetic(divide, 2200, 20))
   fn_pow := func(a, b int) int {
      return int(math.Pow(float64(a), float64(b)))
   }
   fmt.Println("2^8 =", arithmetic(fn_pow, 2, 8))   
}

func arithmetic(fn func(int, int) int, a, b int) int {
   return fn(a,b)
}

func add(a, b int) int {
   return a + b
}

func multiply(a, b int) int {
   return a * b
}

func divide(a, b int) int {
   return a / b
}

