package main

import "fmt"

type student struct{
  name string
  age  int
}

func (s student) getStudent(){
  fmt.Println(s.name)
  fmt.Println(s.age)
}

func main(){
  var s = student{"chandra",19}

  s.getStudent()
}
