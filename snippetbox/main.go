package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox..."))
}
func snippetViwe(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/snippet/view" {
		http.NotFound(w, r)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
	w.Write([]byte("Displayed specific snippet..."))

}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/snippet/create" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		w.Header().Set("allow", http.MethodPost)
		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed "))
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create new sippet..."))
}
func main() {
	fmt.Println("Hello World!")
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", home)
	// mux.HandleFunc("/snippet/view", snippetViwe)
	// mux.HandleFunc("/snippet/create", snippetCreate)
	http.HandleFunc("/", home)
	http.HandleFunc("/snippet/view", snippetViwe)
	http.HandleFunc("/snippet/create", snippetCreate)
	log.Println("starting server in port :4000")
	// err := http.ListenAndServe(":4000", mux)
	err := http.ListenAndServe(":4000", nil)

	log.Fatal(err)

}
