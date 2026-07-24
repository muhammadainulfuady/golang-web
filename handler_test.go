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
