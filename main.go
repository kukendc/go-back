package main

import "net/http"

func dataHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Le Tuyen"))
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/getData", dataHandler)
	http.ListenAndServe(":8000", nil)
}
