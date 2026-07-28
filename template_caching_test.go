package golang_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.gohtml
var templates embed.FS

var myTemplates = template.Must(template.ParseFS(templates, "templates/*.gohtml"))

func TemplateCache(w http.ResponseWriter, r *http.Request) {
	err := myTemplates.ExecuteTemplate(w, "simple.gohtml", map[string]any{
		"Title": "Template Cache",
		"Name":  "Herlambang Bocil",
	})
	if err != nil {
		panic(err)
	}
}

func TestTemplateCache(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateCache(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
