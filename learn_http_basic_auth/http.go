package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

const EXPECTED_USERNAME = "myusername"
const EXPECTED_PASSWORD = "mypassword"

var shutdown chan any = make(chan any)

type apiHandler struct{} // state could be held here

func (apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		fmt.Println("Shutting down server")
		shutdown <- struct{}{}
	}()

	username, password, ok := r.BasicAuth()
	fmt.Printf("BasicAuth: ok=%v, username=%s, password=%s\n", ok, username, password)

	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if username != EXPECTED_USERNAME || password != EXPECTED_PASSWORD {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("\n"))
		fmt.Printf("Bad username or password\n")
		return
	}

	w.Write([]byte("Here is /api\n"))
}

func main() {
	addr := os.Args[1] // usually just a port, e.g. ":8080"

	mux := http.NewServeMux()
	mux.Handle("/api", apiHandler{})

	// return 404 for everything else
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			fmt.Println("Shutting down server")
			shutdown <- struct{}{}
		}()

		http.NotFound(w, r)
		return
	})

	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		err := srv.ListenAndServe()
		fmt.Printf("server returned %v\n", err)
	}()

	<-shutdown
	srv.Shutdown(context.Background())
}
