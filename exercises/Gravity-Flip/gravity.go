package main

import "fmt"

func main(){
  var n, i int
  fmt.Scanln(&n)
  var arr[n]int
  var arr2[101] = {0}int
  int arr3[n]
  for i := 0; i < n; i ++ {
    fmt.Scanln(&arr[i])
    arr3[i] = 0
  }

  flag := 1
  j := 0

  for (flag) {
    flag := 0
    for i := 0; i < n; i++ {
      if(arr[i] > 0) {
        arr2[j] ++
        arr[i] --
        flag = 1
      }
    }
    j ++
  }

  for i := j - 2; i >= 0; i -- {
    x = n - 1
    for (arr2[i]--) {
      arr3[x--]++
    }
  }

  for i := 0; i < n; i++ {
    fmt.Println(arr3[i])
  }
  fmt.Println()

}
