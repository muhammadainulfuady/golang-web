package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

var datas = map[string]any{
	"Nama": "Muhammad Ilham",
	"Umur": 19,
}

var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, datas)
}

func TestHandler(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome To Rest API Dasar Dasar Bangettttttt")
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}
func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Users")
}
func productsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Product")
}
func ordersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Orders")
}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/api/v1/hello", helloHandler)
	mux.HandleFunc("/api/v1/users", usersHandler)
	mux.HandleFunc("/api/v1/products", productsHandler)
	mux.HandleFunc("/api/v1/orders", ordersHandler)
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServeMuxDua(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World")
	})
	mux.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hi")
	})
	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Images")
	})
	mux.HandleFunc("/images/thumnails/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Images Thumbnails")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}


func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Method)
		fmt.Fprintln(w, r.RequestURI)
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
