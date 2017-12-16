package main

import "fmt"

func main(){
  var num, rem, new1 int
  fmt.Scanln(&num)
  if num % 10 == 0 {
    fmt.Println(num)
  } else {
    rem = num % 10
    if rem >= 5 {
      new1 = num - rem
      new1 = new1 + 10
      fmt.Println(new1)
    } else {
      num = num - rem
      fmt.Println(num)
    }
  }
}
