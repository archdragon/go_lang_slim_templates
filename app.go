package main

import (
	"fmt"
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
  var templateError string

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
  tmpl, err := slim.ParseFile("templates/layout.slim")

	if err != nil {
		templateError = "There was a problem parsing the layout"
	} else {
    err = tmpl.Execute(&buf, slim.Values{
      "pages": []Page{
        {Name: "Google", Url: "https://google.com"},
        {Name: "Facebook", Url: "https://facebook.com"},
      },
    })
    if err != nil {
      templateError = "There was a problem displaying the template"
    }
  }

  if templateError == "" {
    data = []byte(buf.String())
  } else {
    data = []byte(templateError)
  }

	w.Header().Set("Content-Length", fmt.Sprint(len(data)))
	fmt.Fprint(w, string(data))
}

func main() {
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
