package main

import "fmt"

func main(){
  var message = make(chan string)

  var whoami = func(){
    message <- "chandra"
  }

  go whoami()

  fmt.Println(<-message)
}
