package main

import (
	"fmt"
	_ "github.com/Epiq122/go-lens/views"
	"github.com/go-chi/chi/v5"
	"html/template"
	"log"
	"net/http"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There wads an error parsing the template", http.StatusInternalServerError)
		return
	}
	tpl.Execute(w, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "templates/home.gohtml")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "templates/contact.gohtml")

}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, "templates/faq.gohtml")

}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("starting the server on :3000...")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
