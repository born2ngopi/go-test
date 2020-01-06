package main

import "fmt"

func main(){
  
  var getArea = func (s int) int {
    return s * s
  }

  result := getArea(5)

  fmt.Println(result)
}
