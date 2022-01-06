package customer_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/santosdavid/go-api-v2/infrastructure/ui/gin/handler/customer"
	"github.com/stretchr/testify/assert"
)

// TODO: mover para outro lugar
func setUpRouter() *gin.Engine {
	r := gin.Default()

	handler, _ := customer.New()
	handler.Make(r)

	return r
}

func TestCreateCustomerError(t *testing.T) {
	tests := []struct {
		testName string
		body     string
		expected map[string]interface{}
	}{
		{
			"email invalid and short passsword",
			`{"payday": 1, "email": "testetest.com.br", "password": "123"}`,
			map[string]interface{}{
				"Email":    "Email must be a valid email address",
				"Password": "Password must be at least 6 characters in length",
			},
		},
		{
			"empty body",
			``,
			map[string]interface{}{},
		},
		{
			"empty json",
			`{}`,
			map[string]interface{}{
				"Email":    "Email is a required field",
				"Password": "Password is a required field",
				"PayDay":   "PayDay is a required field",
			},
		},
	}

	router := setUpRouter()

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {

			var jsonData = []byte(tt.body)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/customers", bytes.NewBuffer(jsonData))
			router.ServeHTTP(w, req)

			var body struct {
				Msg    string
				Errors map[string]interface{}
			}

			bodyBytes, _ := ioutil.ReadAll(w.Body)

			json.Unmarshal(bodyBytes, &body)

			assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
			assert.Equal(t, tt.expected, body.Errors)
		})
	}
}

func TestCreateCustomerWithSucess(t *testing.T) {
	var jsonData = []byte(`{"payday": 5, "email": "test@test.com.br", "password": "123456"}`)

	router := setUpRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/customers", bytes.NewBuffer(jsonData))

	router.ServeHTTP(w, req)

	var body struct {
		ID string
	}

	bodyBytes, _ := ioutil.ReadAll(w.Body)

	if err := json.Unmarshal(bodyBytes, &body); err != nil {
		t.Fatal("error to parse json")
	}

	assert.Equal(t, http.StatusCreated, w.Code)
}
