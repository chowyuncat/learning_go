package main

import (
	"fmt"
	"net/http"
	"os"
)

type apiHandler struct{}

func (apiHandler) ServeHTTP(_ http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	fmt.Printf("BasicAuth: ok=%v, username=%s, password=%s\n", ok, username, password)

	fmt.Println("/api")

	fmt.Println("Shutting down server")
	os.Exit(0)
}

func main() {
	addr := os.Args[1] // usually just a port, e.g. ":8080"

	mux := http.NewServeMux()
	mux.Handle("/api", apiHandler{})

	// return 404 for everything else
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		fmt.Println("/ is the home page")
	})

	http.ListenAndServe(addr, mux)
}
