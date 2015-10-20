package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
    "encoding/json"
)

type Person struct {
  Name string
  Tags []string
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()  //Parse url parameters passed, then parse the response packet for the POST body (request body)
                   // attention: If you do not call ParseForm method, the following data can not be obtained form
    fmt.Println(r.Form)
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello!") // write data to response
}

func person_json(w http.ResponseWriter, r *http.Request){
  person := Person{"Ilya", []string{"ruby", "golang", "angular"}}

  js, err := json.Marshal(person)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method)
    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.html")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
        fmt.Println("username:", r.Form["username"])
        fmt.Println("password:", r.Form["password"])
    }
}

func main() {
    http.HandleFunc("/", sayhelloName)
    http.HandleFunc("/login", login)
    http.HandleFunc("/person.json", person_json)
    err := http.ListenAndServe(":9898", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
