package main

import "fmt"

func main(){
  var num[10]int
  var test int
  fmt.Scanln(&test)
  for i := 0; i < test; i++ {
    fmt.Scanln(&num[i])
  }
  for i := 0; i < test; i++ {
    if (num[i] % 3 == 0) || (num[i] % 7 == 0) || (num[i] % 10 == 0) {
      fmt.Println("YES")
    }else {
      fmt.Println("NO")
    }
  }
}
