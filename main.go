package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "<h1> Welcome to my nightmare </h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1> <p>To get in touch, email me at <a href=\"mailto:epiq@epiq.com\">epiq@epiq.com</a>.</p>")

}

// custom routing
func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		//	TODO: handle page not found error
	}

}

func main() {
	http.HandleFunc("/", pathHandler)
	fmt.Println("starting the server on :3000...")
	http.ListenAndServe(":3000", nil)
}
