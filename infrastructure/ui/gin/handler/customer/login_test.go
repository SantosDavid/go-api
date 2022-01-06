package customer_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginError(t *testing.T) {
	tests := []struct {
		testName string
		body     string
		expected map[string]interface{}
		message  string
		status   int
	}{
		{
			"empty body",
			"{}",
			map[string]interface{}{
				"Email":    "Email is a required field",
				"Password": "Password is a required field",
			},
			"Invalid Entity",
			http.StatusUnprocessableEntity,
		},
		{
			"invalid email",
			`{"email": "test.com.br", "password": "123456"}`,
			nil,
			"email or password invalid",
			http.StatusUnauthorized,
		},
	}

	router := setUpRouter()

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			var jsonData = []byte(tt.body)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))

			router.ServeHTTP(w, req)

			var body struct {
				Message string
				Errors  map[string]interface{}
			}

			bodyBytes, _ := ioutil.ReadAll(w.Body)

			json.Unmarshal(bodyBytes, &body)

			fmt.Print(body.Message)

			assert.Equal(t, tt.status, w.Code)
			assert.Equal(t, tt.expected, body.Errors)
			assert.Equal(t, tt.message, body.Message)
		})
	}
}

//TODO
func TestLoginSucess(t *testing.T) {

	body := `{"email": "test@test.com", "password": "123456"}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(body)))

	router := setUpRouter()
	router.ServeHTTP(w, req)

	// assert.Equal(t, http.StatusOK, w.Code)
	// assert body
}
