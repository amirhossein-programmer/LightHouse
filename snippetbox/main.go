package main

import (
	"log"
	"net/http"
)
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("displaying home page"))
}
func sinppetView(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/snippet/view" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("displaying snippet view..."))
}
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		// w.WriteHeader(405)
		// w.Write([]byte("Method not alowed"))
		// http.Error(w, "Method not alowed", 405)
		http.Error(w, http.MethodPost , http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/snippet/create" 	{
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("displaying snippet create..."))
}
func main(){
	mux:=http.NewServeMux()
	mux.HandleFunc("/",home)
	mux.HandleFunc("/snippet/view",sinppetView)
	mux.HandleFunc("/snippet/create",snippetCreate)
	log.Println("starring server on port :4000")
	err:=http.ListenAndServe(":4000",mux)
	log.Fatal(err)
}