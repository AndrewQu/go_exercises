package main

import (
   "fmt"
   "os"
   "strconv"
   "time"
   "github.com/AndrewQu/go_exercises/tree"
)

var num_operations int = 10

func main() {
//   num_operations := 10
   if len(os.Args) > 1 {
      n, err := strconv.Atoi(os.Args[1])
      if err != nil {
         fmt.Printf("Cannot convert input to int, using default num_operations = %d\n",num_operations)
      } else {
         num_operations = n
      }
   }

   try_select()
   try_select_default()
   try_tree()   
}

func try_select_default() {
   tick := time.Tick(1000 * time.Millisecond)
   boom := time.After(time.Duration(num_operations) * 500 * time.Millisecond)
   
   fmt.Printf("Ticking")   
   for {
      select {
         case <- tick :
            fmt.Printf(" ticked\nTicking")
         case <- boom :
            fmt.Println("Boom and exit")
            return
         default :
            fmt.Printf(".")
            time.Sleep(50 * time.Millisecond)
      }
   }
}

func try_select() {
   ch := make(chan int, 10)
   quit := make(chan int)
   
   go func() {   
      fmt.Printf("Started:")
      for i := 0; i < num_operations; i++ {
         fmt.Printf(" %d", <- ch)
         if (i + 1) % 10 == 0 {
            fmt.Println("")
         }
      }
      quit <- 1	// Send quit instruction
   }()

   fibonacci(ch, quit)
}

func fibonacci( ch, quit (chan int) ) {
   x, y := 0, 1
   for {
      select {
      case ch <- x :
         x, y = y, x + y
      case <- quit :
         fmt.Println("\nQuit received.")
         return
      }
   }
}

func try_tree() {
   // COmpare if 2 trees are the same
   tree_values1 := make (chan int, 10)
   tree_values2 := make (chan int, 10)
   done := make(chan struct{})
   
   comp_state := 2  // 2=not set, 1=same, 0=differeent
   
   go func () {
      for {
//         fmt.Println("Retriving v1")
         v1, ok1 := <- tree_values1
//         fmt.Println("Retrieved v1 =", v1, "Retriving v2")
         v2, ok2 := <- tree_values2
//         fmt.Println("Retrieved v2 =", v2)
         if v1 != v2 || ok1 != ok2 {
            comp_state = 0	// Not the same
            break
         }
         if ok1 == false && ok1 == ok2 {
            comp_state = 1;	// The same
            break
         }
      }
      close(done)
   } ()
   
   tree1 := tree.NewTree(1, 64)
   tree2 := tree.NewTree(1, 64)

   go func () {
      tree1.StartWalk(tree_values1, done)
   } ()
   
   tree2.StartWalk(tree_values2, done)
   
   for ; comp_state == 2; {
      fmt.Println("Waiting signal to end...")
      time.Sleep(50 * time.Millisecond)
   }
   
   if comp_state == 1 {
      fmt.Println("2 trees are the same!")
   } else if (comp_state == 0) {
      fmt.Println("2 trees are different!")
   } else {
      fmt.Println("State unknown")
   }
   
}

