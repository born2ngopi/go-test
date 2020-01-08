package main

import (
  "fmt"
  "net/http"
)

func index(w http.ResponseWriter, r *http.Request){
  fmt.Fprintln(w, "welcome to go")
}

func main(){
  http.HandleFunc("/",index)

  fmt.Println("starting server in localhost:8000")
  http.ListenAndServe(":8000",nil)
}
