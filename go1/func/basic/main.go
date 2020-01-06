package main

import "fmt"

func sum(a,b int){
  fmt.Println(a + b)
}

func multiplies(a int, b float32)float32{
  //var c float32
  return float32(a) * b
}

func main(){
  sum(1,2)
  result := multiplies(2,0.5)

  fmt.Println(result)
}
