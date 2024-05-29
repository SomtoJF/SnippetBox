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

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view/{id}", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Running on localhost 4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}