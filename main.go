package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Selamat datang, Salam Sukses dan Jangan LUPA BAHAGIA HARI INI!!!")
	})

	http.ListenAndServe(":80", nil)
}
