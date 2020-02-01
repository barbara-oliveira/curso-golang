package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func horaCerta(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:03:05") //estranhamente isso é a formatação
	fmt.Fprintf(w, "<h1> Hora certa: %s </h1>", s)
}

func main() {
	http.HandleFunc("/horacerta", horaCerta)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
