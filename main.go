package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/thing.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		fmt.Fprint(w, `{"get": "shwifty"}`)
	})
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", nil)
}
