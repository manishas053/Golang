package main

import "fmt"

func main(){
  var da[10]int
  var db[10]int
  s := 0
  t := 0
  a := 0
  b := 0
  m := 0
  n := 0
  pos := 0
  count := 0
  fmt.Println("Start and end points of Sam's house: ")
  fmt.Scanln(&s)
  fmt.Scanln(&t)
  fmt.Println("Position of apple and orange tree: ")
  fmt.Scanln(&a)
  fmt.Scanln(&b)
  fmt.Println("No. of apples and oranges: ")
  fmt.Scanln(&m)
  fmt.Scanln(&n)
  fmt.Println("d values of each apple: ")
  for i := 0; i < m; i++ {
    fmt.Scanln(&da[i])
  }
  fmt.Println("d values of each orange: ")
  for j := 0; j < n; j++ {
    fmt.Scanln(&db[j])
  }
  for i := 0; i < m; i++ {
    if da[i] < 0 {
      pos = a - da[i]
    } else {
      pos = a + da[i]
      if (pos >= s) && (pos <= t) {
        count ++
      }
    }
  }
  fmt.Println(count)
  count = 0
  pos1 := 0
  for j := 0; j < n; j++ {
    if db[j] < 0 {
      pos1 = b - db[j]
      if (pos1 >= s) && (pos1 <= t) {
        count ++
      }
    } else {
      pos1 = b + da[j]
    }
  }
  fmt.Println(count)
}
