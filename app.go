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

func getTemplate() []byte {
  var buf bytes.Buffer

  tmpl, err := slim.ParseFile("templates/layout.slim")

	if err != nil { return []byte("There was a problem parsing the layout") }

  err = tmpl.Execute(&buf, slim.Values{
    "pages": []Page{
      {Name: "Google", Url: "https://google.com"},
      {Name: "Facebook", Url: "https://facebook.com"},
    },
  })

  if err != nil { return []byte("There was a problem displaying the template") }

  return []byte(buf.String())
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
  var templateData []byte

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

  templateData = getTemplate()

	w.Header().Set("Content-Length", fmt.Sprint(len(templateData)))
	fmt.Fprint(w, string(templateData))
}

func main() {
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
