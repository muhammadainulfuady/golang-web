package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func RedirectTo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You are being redirected!")
}

func RedirectFrom(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/redirect_to", http.StatusTemporaryRedirect)
}

func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect_from", RedirectFrom)
	mux.HandleFunc("/redirect_to", RedirectTo)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}

func RedirectToIndex(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, "Redirect Success! Hallo ", name)
}

func RedirectFromIndex(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "Harap masukkan query ?name=... terlebih dahulu")
		return
	} else {
		http.Redirect(w, r, "/redirect_to?name="+name, http.StatusTemporaryRedirect)
		return
	}
}

func TestRedirectIndex(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect_from", RedirectFromIndex)
	mux.HandleFunc("/redirect_to", RedirectToIndex)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
