package main

import "fmt"

func square(s int)(area, keliling int){
  area = s * s
  keliling = 4 * s

  return
}

func main(){
  area, keliling := square(6)

  fmt.Printf("luas persigi %d\n keliling persegi %d\n",area, keliling)
}
