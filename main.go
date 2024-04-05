package main

import (
	"fmt"
	"log"
	"net/http"
)

// Creating formHandler Function
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err:%v", err)
		return
	}
	fmt.Fprintf(w, "POST request Successfull")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address= %s\n", address)

}

// Creating Hellohandler Function
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Error not found", http.StatusFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not Supported", http.StatusNotFound)
		return

	}
	fmt.Fprintf(w, "hello GoLangByLaddu!")
}

// Main Function
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	//Starting Server port 8082
	fmt.Printf("Starting server at port  8080\n")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal(err)
	}
}
