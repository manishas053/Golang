package main

import "fmt"

func main() {
  var op, row int
  num := 1
  fmt.Println("Inverted(1) or straight(2) : ")
  fmt.Scanln(&op)
  fmt.Println("No. of rows : ")
  fmt.Scanln(&row)
  switch(op) {
  case 1:
    space := 1
    for(i := row; i >= 1; i--){
      for(j := 1; j <= space; j++){
        fmt.Println(" ")
      }
      for(j = 1; j <= i; j++){
        fmt.Printf("%d ", num)
        num ++
      }
      fmt.Println()
      space ++
    }
    break

  case 2:
    space := row
    for(i := 1; i <= row; i++){
      for(j := 1; j < space; j++){
        fmt.Println(" ")
      }
      for(j = 1; j <= i; j++){
        fmt.Printf("%d ", num)
        num++
      }
      fmt.Println()
      space --
    }
    break
    }
  }
