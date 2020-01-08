package main

import (
  "fmt"
  "html/template"
  "net/http"
)

func main(){
  http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
    var data = map[string]string{
      "name":"chandra"
    }

    var t, err = template.ParseFiles("index.html")
    if err != nil{
      fmt.Println(err.Error())
      return
    }

    t.Execute(w, data)
  })

  fmt.Println("8000")
  http.ListenAndServe(":8000",nil)

}
