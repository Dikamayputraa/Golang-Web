package handler

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
	"Web/entity"
)

func HomeHendler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))

	if err != nil {
		log.Println(err)

		http.Error(w, "Error is happening, keep calm!", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"title":   "I'am learning Golang Web.",
		"content": "I'am learning GOlang Web From BWA.",
	}

	tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)

		http.Error(w, "Error is happening, keep calm!", http.StatusInternalServerError)
		return
	}

	// w.Write([]byte("Hello, ini halaman pertama Golang"))

}

func ProductHendler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	//validasi karakter yang tidk sesuai di url
	idNumb, err := strconv.Atoi(id)
	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}
	// fmt.Fprintf(w, "Product Id : %d", idNumb)

	// data := map[string]interface{}{
	// 	"content": idNumb,
	// }

	// data := entity.Product{ID: 1, Nama : "mobilio", Price: 2000000, Stok: 10}
	data := []entity.Product{
		{ID: 1, Nama : "mobilio", Price: 2000000, Stok: 7},
		{ID: 2, Nama : "Expander", Price: 2300000, Stok: 11},
		{ID: 3, Nama : "Innova", Price: 300000, Stok: 2},
	}

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm!", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm!", http.StatusInternalServerError)
		return
	}

}

func Form(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		tmpl,err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm!", http.StatusInternalServerError)
			return
		}
		
		err = tmpl.Execute(w, nil)
		if err != nil{
			log.Println(err)
			http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
			return
		}
		 return
	}

	http.Error(w, "Error is happening, keep calm!", http.StatusInternalServerError)

}

func Process(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		err := r.ParseForm()
		if err != nil{
			log.Println(err)
			http.Error(w, "Error is happening, keep calm!", http.StatusInternalServerError)
			return
		}

		name := r.Form.Get("name")
		pesan := r.Form.Get("pesan")

		data := map[string]interface{}{
			"name"	: name,
			"pesan"	: pesan,
		}
		
		tmpl, err := template.ParseFiles(path.Join("views", "result.html"), path.Join("views", "layout.html"))
		if err != nil{
			log.Println(err)
			http.Error(w, "Error is happening, keep calm!", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil{
			log.Println(err)
			http.Error(w, "Error is happening, keep calm!", http.StatusInternalServerError)
			return
		}
		return
	}
	http.Error(w, "Error is happening, keep calm!", http.StatusInternalServerError)

}

func PostGet(w http.ResponseWriter, r *http.Request){
	method := r.Method

	switch method{
	case "GET":
		w.Write([]byte("Ini adalah GET"))
	case "POST":
		w.Write([]byte("Ini adalah POST"))
	default:
		http.Error(w, "Error is hapining, keep cal,!", http.StatusBadRequest)
	}
}

func HelloHendler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to MyGolang"))
}

func ContactHendler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, ini halaman pertama Golang"))
}
