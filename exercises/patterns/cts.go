package main

import "fmt"

func main(){
  var row, i, j, k, m int
  fmt.Println("No. of rows : ")
  fmt.Scanln(&row)
  for i := 0; i < row; i = i + 2 {
    k := 1
    for j = 0; j < row; j ++ {
      fmt.Printf("%d", row * i + k)
      k ++
      if j != row - 1 {
        fmt.Printf("*")
      }
    }
    fmt.Println()
  }

  if(row % 2 == 0) {
    m = i - 2
  }

  for i = m; i > 0; i = i - 2 {
    k = 1
    for j = 0; j < row; j ++ {
      fmt.Printf("%d", row * i + k)
      k ++
      if j != row - 1 {
        fmt.Printf("*")
      }
    }
    fmt.Println()
  }
}
