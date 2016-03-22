// Copyright © 2016 Ryan D <ryan0x44.com>
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello API")
	})
	http.ListenAndServe(":8080", r)
}
