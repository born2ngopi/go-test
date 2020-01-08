package main

import "fmt"

//func getNumbers(messages chan<- int, done chan bool){
//  for i:=0;i<1000;i++{
//    messages<- i
//  }
//  done<- true
//}

func main(){
    messages := make(chan int, 2)
    
    go func() {
        for {
            i := <-messages
            fmt.Println("receive", i)
        }
    }()

    for i := 0; i < 10; i++ {
        fmt.Println("send", i)
        messages <- i
    }
}
