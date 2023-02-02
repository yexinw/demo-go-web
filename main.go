package main

import "net/http"

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello wolrd,api"))
	})
	http.HandleFunc("/api/hnicc", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello wolrd,hnicc"))
	})
	http.HandleFunc("/api/abicc", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello wolrd,abicc"))
	})
	http.HandleFunc("/api/wwwuqbuycom", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello wolrd,wwwuqbuycom"))
	})
	http.HandleFunc("/api/kuliren", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello wolrd,kuliren"))
	})
	http.ListenAndServe(":8080", nil)
}
