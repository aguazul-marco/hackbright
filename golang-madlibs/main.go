package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
  
	t.templ.Execute(w, r)
}

func main() {
	http.Handle("/", &templateHandler{filename: "index.html"})
	http.Handle("/greet", &templateHandler{filename: "greeting.html"})
	http.Handle("/words", &templateHandler{filename: "word.html"})
	http.Handle("/madlib", &templateHandler{filename: "madlib.html"})

	// Start the web server
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("Ran into an error:", err)
	} else {
		log.Println("Serving on http://localhost:8000")
	}
}
