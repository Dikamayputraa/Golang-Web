package main

import (
	"log"	
	"net/http"
	"Web/handler"
)

func main() {
	mux := http.NewServeMux()

	//cara kedua buat routing
	// contactHendler := func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Contact Here!"))
	// }

	mux.HandleFunc("/", handler.HomeHendler) //root
	mux.HandleFunc("/hello", handler.HelloHendler)
	mux.HandleFunc("/product", handler.ProductHendler)
	mux.HandleFunc("/contact", handler.ContactHendler)
	mux.HandleFunc("/post-get", handler.PostGet)
	mux.HandleFunc("/form", handler.Form)
	mux.HandleFunc("/process", handler.Process)

	//load css
	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))


	//cara ketiga buat outing dengan function anonimous (tidak bernama)
	// mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Profile Page!"))
	// })

	log.Println("Starter web on port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}

