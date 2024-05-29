package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(res http.ResponseWriter, req *http.Request){
	res.Write([]byte("Hello from Snippetbox"))
}

func snippetView(res http.ResponseWriter, req *http.Request){
	snippetId, err := strconv.Atoi(req.PathValue("id"))

	if err != nil || snippetId < 1 {
		http.NotFound(res, req) 
		return 
	}

	msg := fmt.Sprintf("Display a specific snippet with ID %d...", snippetId) 
	res.Write([]byte(msg))
}

func snippetCreate(res http.ResponseWriter, req *http.Request) {
	 res.Write([]byte("Display a form for creating a new snippet...")) 
}

func snippetCreatePost(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte("Posting a new snippet..."))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create/post", snippetCreatePost)

	log.Print("Running on localhost 4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}