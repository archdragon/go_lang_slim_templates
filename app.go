package main

import (
	"fmt"
  // "io/ioutil"
  "bytes"
	"log"
	"net/http"
  "github.com/mattn/go-slim"
)

type Page struct {
  Name string
  Url string
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
  var buf bytes.Buffer
  var data []byte

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
  tmpl, err := slim.ParseFile("templates/layout.slim")

	if err != nil {
		data = []byte("404 - no layout")
	} else {
    err = tmpl.Execute(&buf, slim.Values{
      "pages": []Page{
        {Name: "Google", Url: "http://google.com"},
      },
    })
    if err != nil {
      data = []byte("404 - layout parse")
    }
  }

  if data == nil {
    data = []byte(buf.String())
  }

	w.Header().Set("Content-Length", fmt.Sprint(len(data)))
	fmt.Fprint(w, string(data))
}

func main() {
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
