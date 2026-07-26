package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func PostForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	firstName := r.PostForm.Get("firstName")
	lastName := r.PostForm.Get("lastName")

	fmt.Fprintf(w, "Nama gw '%s' dan '%s'", firstName, lastName)
}

func TestPostForm(t *testing.T) {
	requestBody := strings.NewReader("firstName=ilham&lastName=bocil")
	request := httptest.NewRequest(http.MethodPost, "http://localhost/", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	encoder := httptest.NewRecorder()

	PostForm(encoder, request)

	response := encoder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
