/*****************************************************************************
For n=5, print
1
3*2
4*5*6
10*9*8*7
11*12*13*14*15

*******************************************************************************/


package main

import "fmt"

func main(){
  var n int
  count := 0
  k := 0
  fmt.Printf( "Enter N : ")
  fmt.Scanln(&n)
  for i := 1; i <= n; i++ {
    count = k
    for j := 1; j <= i; j++ {
      if i % 2 == 0 {
        fmt.Printf("%d", count + i)
        count = count - 1
        if j != i {
          fmt.Printf("*")
        }
        k ++
      } else {
        count = count + 1
        fmt.Printf("%d", count)
        if j != i {
          fmt.Printf("*")
        }
        k ++
      }
    }
    fmt.Printf("\n")
  }
}
