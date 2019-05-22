package main

import (
  "encoding/json"
  "net/http"
  "fmt"
  "html"
  "time"

  "github.com/gorilla/mux"
)

type Todo struct {
    Name      string    `json:"name"`
    Completed bool      `json:"completed"`
    Due       time.Time `json:"due"`
}

type Todos []Todo

func Index(resp http.ResponseWriter, req *http.Request)  {
  resp.Write([]byte("ok"))
}


func Article(resp http.ResponseWriter,req *http.Request)  {
  vars := mux.Vars(req)
  title := vars["title"]
  page := vars["page"]

	fmt.Fprintf(resp, "Anda sendang mengakses - Artikel dengan judul: %s on page %s\n %s", title, page, html.EscapeString(req.URL.Path))

}

func TodoIndex(resp http.ResponseWriter, req *http.Request) {
    todos := Todos{
        Todo{Name: "Write presentation"},
        Todo{Name: "Host meetup"},
    }

    if err := json.NewEncoder(resp).Encode(todos); err != nil {
        panic(err)
    }
}

func TodoShow(resp http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    todoId := vars["todo-id"]
    fmt.Fprintln(resp, "Todo show:", todoId)
}


func Hello(resp http.ResponseWriter, req *http.Request)  {
  fmt.Fprintf(resp, "Hello, %q", html.EscapeString(req.URL.Path))
}
