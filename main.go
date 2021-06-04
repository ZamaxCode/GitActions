package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprintf(w, "Pagina no encontrada")
	} else {
		template.Execute(w, nil)
	}
}

func main() {

	http.HandleFunc("/", index)

	fmt.Println("La aplicacion está corriendo en el puerto 3000")
	fmt.Println("Entre al siguiente link para ingresar: localhost:3000")
	http.ListenAndServe(":3000", nil)
}
