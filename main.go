package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Selamat datang di bro di dunia virtual  gw!!!")
	})

	http.ListenAndServe(":80", nil)
}
