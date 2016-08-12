package main

import (
   "fmt"
   "math"
   "errors"
   "io"
   "strings"
   "os"
   "time"
)

func main() {
   try_error()
   try_string_reader()
   try_goroutine()
   try_channels()
}

func summation(nums []int, c (chan int)) {
   sum := 0
   for _, v := range nums {
      sum += v
   }
   c <- sum
}

func try_channels() {
   c := make(chan int)
   num1 := []int {3, 2, 4, 9, 8, 7}
   num2 := []int {11, 15, 17, 12, 18, 19}
   go summation(num1, c)
   go summation(num2, c)
   sum1 := <- c
   sum2 := <- c
   fmt.Printf("Summation result: %d + %d = %d\n", sum1, sum2, sum1 + sum2)
   
   ch := make(chan int, 2)
   go recv_and_print(ch)

   for i := 0; i < 40; i++ {
      ch <- i+1
      time.Sleep(100 * time.Millisecond)
   }
   close(ch)	// Without closing it, recv_and_print() will not exit. It won't print the
   // last Total = line. But the thread will be forced to close when the main program ends.
   time.Sleep(100 * time.Millisecond)
}

func recv_and_print( ch (chan int) ) {
   v, ok := <- ch
   if !ok { return }
   fmt.Printf("%2d ", v)
   
   i := 1
   for v = range ch {
      fmt.Printf("%2d ", v)
      if (i+1) % 20 == 0 { fmt.Println() }
      i++
   }
   fmt.Printf("Exit recv_and_print() Total=%d\n", i)
}

func try_goroutine() {
   go say("Hello 1")
   say("Hello 2")
}

func say(s string) {
   for i := 0; i<5; i++ {
      fmt.Println(s)
      time.Sleep(100 * time.Millisecond)
   }
}

func try_string_reader() {
   sr := strings.NewReader("1234567890123456789")
   buffer := make([]byte, 8)
   for {
      n, err := sr.Read(buffer)
      if err == io.EOF {
         break
      }
      fmt.Printf("Chars read : %q\n", buffer[:n])
   }
   sr = strings.NewReader("Pbatenghyngvbaf! Vg jbexrq.")
   io.Copy(os.Stdout, &rot13Reader{ sr })
   fmt.Println()
}

type rot13Reader struct {
   r io.Reader
}

func (r13 rot13Reader) Read(b []byte) (int, error) {
   if cap(b) == 0 {
      return 0, errors.New("Storage buffer is nil")
   }
   if len(b) < cap(b) {
      b = b[:cap(b)]
   }
   nr, err := r13.r.Read(b)
   if err == io.EOF {
      return 0, err
   }
   for i := 0; i < nr; i++ {
      p := b[i] - 'a'
      if p < 0 || p >= 26 {	// Not a lower case letter
         p = b[i] - 'A'		// Assume an upper case letter
      }
      if p >= 0 && p < 13 {
         b[i] += 13
      } else {
         if p >= 13 && p <= 26 {
  	    b[i] -= 13
         }
      }
   }
   return nr, err
}

func try_error() {
   v1 := Vector2d {10, 10}
   err := v1.Normalize()
   if err != nil {
      fmt.Println("v1 -", err)
   } else {
      fmt.Println("v1 =", v1)
   }
   ang, err := v1.AngleToX()
   fmt.Println("v1 angle-x =", ang, "Err =", err)
   
   v2 := Vector2d {0, 0}
   err = v2.Normalize()
   if err != nil {
      fmt.Println("v2 -", err)
   } else {
      fmt.Println("v2 =", v2)
   }
   ang, err = v2.AngleToX()
   fmt.Println("v2 angle-x =", ang, "Err =", err)
}   

var errZeroLengthVector error = errors.New("Vector length is 0")

// Find angle of the vector to the x-axis
func (v Vector2d) AngleToX() (ang float64, err error) {
   err = v.Normalize()
   if err != nil {
      err = errZeroLengthVector
      return
   }
   ang = math.Acos(v.X)
   if math.Abs(v.X) < 1.0e-20 {
      if v.Y > 0 {
         ang = math.Pi / 2.0
      } else {
         ang = math.Pi * 1.5
      }
   } else {
      if v.Y < 0 {
         ang = math.Pi * 2.0 - ang
      }
   }
   return
}

type Point2d struct {
   X, Y float64
}

type Vector2d struct {
   X, Y float64
}

type MyError struct {
   code int
}

func (err MyError) Error() string {
   switch err.code {
      case 1 :
         return "Cannot normalize vector, length=0"
      case 2 :
         return "Error code 2"
      default :
         return fmt.Sprintf("Error code :%d", err.code)
   }
}

func (v *Vector2d ) Normalize() error {
   l := math.Sqrt(v.X * v.X + v.Y * v.Y)
   if math.Abs(l) < 1.0e-20 {
      return MyError{1}
   } else {
      v.X /= l
      v.Y /= l
      return nil
   }
}
