package tree

import "fmt"

type Tree struct {
   Left *Tree
   Right *Tree
   Value int
}

// Create a tree of base * 100 + (0 to n) where n < 99
var i_current int = 0
var i_max int = 0
var node_to_do chan *Tree

func SetBranch() {
   t, ok := <- node_to_do
   if !ok { return }	// No more
   
   if i_current >= i_max {
      if i_current == i_max { close(node_to_do) }
      i_current++
   } else {
//      fmt.Println("Adding : ", t.Value - t.Value % 100 + i_current)
      t.Left = &Tree{ Value : t.Value - t.Value % 100 + i_current }
      i_current++
      node_to_do <- t.Left
      if i_current == i_max {
         close(node_to_do)
         i_current++
      } else {
//         fmt.Println("Adding : ", t.Value - t.Value % 100 + i_current)
         t.Right = &Tree{ Value : t.Value - t.Value % 100 + i_current }
         i_current++
         node_to_do <- t.Right
      }
   }
   SetBranch()
}

func NewTree(base, n int) Tree {
   t := Tree { Value : base * 100 }
   i_current = 1
   i_max = n
   i := uint(1)
   for ; ; i++ {
      n = n >> 1
      if n == 1 { break }
   }
   ch_len := 1 << i
   fmt.Println("Setting tree", base * 100, " n =", i_max, " ch_len =", ch_len, i)

   node_to_do = make(chan *Tree, ch_len)
   node_to_do <- &t
   
   SetBranch()
   return t
}

func walk(t *Tree, ch chan int, done chan struct{}) {
   select {
      case ch <- t.Value :
      case <- done :
         return
   }
   if t.Left != nil { walk(t.Left, ch, done) }
   if t.Right != nil { walk(t.Right, ch, done) }
}

func (t *Tree) StartWalk(ch chan int, done chan struct{}) {
   walk(t, ch, done)
   close(ch)
}

