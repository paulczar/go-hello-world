package main

import (
    "net/http"
    "io/ioutil"
    "fmt"
    "github.com/dimiro1/health"
)

func loadFile(filename string) string {
    content, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println("index.html not found, serving default hello world page")
        return "<html><head><title>hello world</title></head><body>hello world!</body></html>\n"
    }
    fmt.Println("serving index.html file")
    return string(content)
}

func serve(w http.ResponseWriter, r *http.Request) {
  var content string
  if r.URL.Path[1:] != "" {
    content = fmt.Sprintf("<html><head><title>hello %s</title></head><body>hello %s!</body></html>\n", r.URL.Path[1:], r.URL.Path[1:])
  } else {
    content = loadFile("index.html")
  }
  fmt.Fprintf(w, "%s", content)
}

func hello(w http.ResponseWriter, r *http.Request) {
  content := loadFile("index.html")
  fmt.Fprintf(w, "%s", string(content))
}

func main() {
    fmt.Println("starting hello world app")
    healthHandler := health.NewHandler()
    http.Handle("/health/", healthHandler)
    http.HandleFunc("/", serve)
    http.ListenAndServe(":8080", nil)
}
