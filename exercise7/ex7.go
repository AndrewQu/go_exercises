package main

import (
   "fmt"
   "sync"
   "time"
   "reflect"
)

func main() {
   try_mutex()
   try_panic_recover()
   try_type_compare()
}

func try_type_compare() {
   v1 := 100
   v2 := 100.0
   fmt.Printf("%%Tv1=%T  %%Tv2=%T\n", v1, v2)
   fmt.Printf("TypeOf(v1)=%v  TypeOf(v2)=%v\n", reflect.TypeOf(v1), reflect.TypeOf(v2))
   if reflect.TypeOf(v1) == reflect.TypeOf(v2) {
      fmt.Println("v1, 2 types are the same")
   } else {
      fmt.Println("v1, 2 types are different")
   }   
}

func try_panic_recover() {
   defer func () {
      error := recover()
      if reflect.TypeOf(error).String() == "runtime.errorString" {
         fmt.Printf("Recover: %v \n", error)
      } else {
         fmt.Printf("Revocer: unknown panic error.\n")
      }
   }()
   
   arr := []int {3, 4, 5, 6, 7, 8}
   for i := 0; i < 50; i++ {
      fmt.Println(arr[i])
   }
}

type SafeCounter struct {
   v map[string] int
   mux sync.Mutex
}

func try_mutex() {
   counter := SafeCounter { v : make(map[string] int) }
   for i := 0; i < 100; i++ {
      go counter.Increment("Cloud")
   }
   time.Sleep( time.Second )

   v, ok := counter.Count("Cloud")
   fmt.Println("Count =", v, ok)
   
   v1, ok := counter.Count("Class")
   fmt.Println("Count =", v1, ok)
}

// Increase the counter for the given key
func (c *SafeCounter) Increment(key string) {
   c.mux.Lock()
   c.v[key]++
   c.mux.Unlock()
}

func (c *SafeCounter) Count(key string) (int, bool) {
   c.mux.Lock()
   
   defer c.mux.Unlock()
   v, ok := c.v[key]
   return v, ok
}

