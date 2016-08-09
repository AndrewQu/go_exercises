package main

import (
   "fmt"
   "time"
   "math/rand"
   "math"
)

var kk int = 3

func main() {
   fmt.Println("math.Sqrt(16)=", math.Sqrt(16))
   fmt.Println("math.Sin(30) =", math.Sin(30.0 * math.Pi / 180.0))
   
   // calling locally defined fun add(x,y)
   fmt.Println("calling locally defined add(20,30) =", add(20, 30))
   
   // Call func that returns 2 strings
   str1, str2 := ret_strings()
   fmt.Println("Returned str1 =", str1)
   fmt.Println("Returned str2 =", str2)
   
   t := time.Now()
   fmt.Printf("Hello world, connected to playground!\n");
   fmt.Println("The time is ", t);
   
   // Generate random numbers
   rand.Seed(int64(t.Minute() * 60 + t.Second()))
   fmt.Printf("5 Random numbers :")
   for i := 0; i < 5; i++ {
      fmt.Printf(" %v", rand.Intn(1000))
   }
   fmt.Printf("\n")
   
   // for is the only loop statement in go. Equivalent while(...) {...}
   fmt.Printf("kk from %v to 9 :", kk)
   for kk < 10 {
      fmt.Printf(" %v", kk)
      kk++
   }
   fmt.Printf("\n")
   
   // Call square root function
   fmt.Println("sqrt(2.0) =", sqrt(2.0), " (approximate)")
   
   // Cal switch function
   when_is_saturday()
   
}

func add(x, y int) int {
   return x + y
}

func ret_strings() (str1, str2 string) {
   str1 = "This is the first returned string 1"
   str2 = "This is the second returned string 2"
   return
}

func sqrt(x float64) float64 {
   z := 1.0
   for i := 0; i < 100; i++ {
      if dz := (z * z - x) / 2.0 / z; math.Abs(dz) < 0.00001 {
         break
      } else {
         z = z - dz
      }
      fmt.Println("z =", z)
   }
   return z
}

func when_is_saturday() {
   today := time.Now().Weekday()
   switch days := time.Saturday - today; days {
      case 0:
         fmt.Printf("Today")
      case 1:
         fmt.Printf("Tomorrow")
      case 2:
         fmt.Printf("The day after tomorrow")
      default:      
         fmt.Printf("%v days later", int(days))      
   }
   fmt.Println(" is Saturday")
   
}
