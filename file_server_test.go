package golang_web

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	fmt.Println("Run in http://localhost:8080")

	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}

}

//go:embed resources
var resources embed.FS

func TestFileServerGolangEmbed(t *testing.T) {
	fmt.Println("Run in http://localhost:8080")

	sub, err := fs.Sub(resources, "resources")
	if err != nil {
		return
	}
	fileServer := http.FileServer(http.FS(sub))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err = server.ListenAndServe()
	if err != nil {
		return
	}

}
