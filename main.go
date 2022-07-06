package main

import (
	"fmt"
	"log"
	"net/http"
)

func formhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() error %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successfull\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address= %s\n", address)
}

func hellohandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {
		http.Error(w, "404 page not found\n", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not supported\n", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "hello!!")

}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/hello", hellohandler)
	http.HandleFunc("/form", formhandler)

	fmt.Printf("Starting the server at port 8080.......\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
