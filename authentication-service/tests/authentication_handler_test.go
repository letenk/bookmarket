package tests

import (
	"authentication_service/cmd/config"
	"authentication_service/cmd/models/web"
	"authentication_service/cmd/routes"
	"authentication_service/pkg"

	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/rizkydarmawan-letenk/jabufaker"
	"github.com/stretchr/testify/assert"
)

func registerRandomUserHandler(t *testing.T) web.RegisterInput {
	// Load file .env
	godotenv.Load("../.env")
	// Open connection
	// Connection to DB
	conn := config.SetupDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}
	defer conn.Close()

	// Truncate table users before test running
	pkg.TruncateUsers(conn)

	// Use route
	router := routes.SetupRouter(conn)

	// Getting random province and regencies for city
	province := jabufaker.RandomProvince()
	regency := jabufaker.RandomRegency(province)

	// Payload
	data := web.RegisterInput{
		Fullname: jabufaker.RandomPerson(),
		Email:    jabufaker.RandomEmail(),
		Address:  jabufaker.RandomString(20),
		City:     regency,
		Province: province,
		Mobile:   "082277760694",
		Password: "password",
		Role:     "admin",
	}

	dataBody := fmt.Sprintf(`{"fullname": "%s", "email": "%s", "address": "%s", "city": "%s", "province": "%s", "mobile": "%s", "password": "%s", "role": "%s"}`, data.Fullname, data.Email, data.Address, data.City, data.Province, data.Mobile, data.Password, data.Role)

	// Create reader
	requestBody := strings.NewReader(dataBody)

	// Create request
	request := httptest.NewRequest(http.MethodPost, "http://localhost:80/api/v1/register", requestBody)
	// Added header content type
	request.Header.Add("Content-Type", "application/json")

	// Create recorder
	recorder := httptest.NewRecorder()

	// Run handler
	router.ServeHTTP(recorder, request)

	// Get response
	response := recorder.Result()

	// Read response
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// Test pass
	assert.Equal(t, 201, response.StatusCode)
	assert.Equal(t, 201, int(responseBody["code"].(float64)))
	assert.Equal(t, "success", responseBody["status"])
	assert.Equal(t, "You have successfully registered", responseBody["message"])

	return data
}

func TestRegisterHandlerSuccess(t *testing.T) {
	registerRandomUserHandler(t)
}

func TestRegisterHandlerFailed(t *testing.T) {
	// Load file .env
	godotenv.Load("../.env")
	// Open connection
	// Connection to DB
	conn := config.SetupDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}
	defer conn.Close()

	// Use route
	router := routes.SetupRouter(conn)

	t.Run("Email already exists", func(t *testing.T) {
		// Create new user
		newUser := registerRandomUserHandler(t)
		// Getting random province and regencies for city
		province := jabufaker.RandomProvince()
		regency := jabufaker.RandomRegency(province)

		// Payload
		data := web.RegisterInput{
			Fullname: jabufaker.RandomPerson(),
			Email:    newUser.Email,
			Address:  jabufaker.RandomString(20),
			City:     regency,
			Province: province,
			Mobile:   "082277760694",
			Password: "password",
			Role:     "admin",
		}

		dataBody := fmt.Sprintf(`{"fullname": "%s", "email": "%s", "address": "%s", "city": "%s", "province": "%s", "mobile": "%s", "password": "%s", "role": "%s"}`, data.Fullname, data.Email, data.Address, data.City, data.Province, data.Mobile, data.Password, data.Role)

		// Create reader
		requestBody := strings.NewReader(dataBody)

		// Create request
		request := httptest.NewRequest(http.MethodPost, "http://localhost:80/api/v1/register", requestBody)
		// Added header content type
		request.Header.Add("Content-Type", "application/json")

		// Create recorder
		recorder := httptest.NewRecorder()

		// Run handler
		router.ServeHTTP(recorder, request)

		// Get response
		response := recorder.Result()

		// Read response
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		// Test pass
		assert.Equal(t, 400, response.StatusCode)
		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "error", responseBody["status"])
		assert.Equal(t, "Register failed.", responseBody["message"])
		assert.Equal(t, "email already exist", responseBody["data"].(map[string]interface{})["errors"])
	})

	t.Run("Validation error", func(t *testing.T) {
		dataBody := fmt.Sprintf(`{"fullname": "%s", "email": "%s", "address": "%s", "city": "%s", "province": "%s", "mobile": "%s", "password": "%s", "role": "%s"}`, "", "", "", "", "", "", "", "")

		// Create reader
		requestBody := strings.NewReader(dataBody)

		// Create request
		request := httptest.NewRequest(http.MethodPost, "http://localhost:80/api/v1/register", requestBody)
		// Added header content type
		request.Header.Add("Content-Type", "application/json")

		// Create recorder
		recorder := httptest.NewRecorder()

		// Run handler
		router.ServeHTTP(recorder, request)

		// Get response
		response := recorder.Result()

		// Read response
		body, _ := io.ReadAll(response.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)

		// Test pass
		assert.Equal(t, 400, response.StatusCode)
		assert.Equal(t, 400, int(responseBody["code"].(float64)))
		assert.Equal(t, "error", responseBody["status"])
		assert.Equal(t, "Register failed.", responseBody["message"])
		assert.NotEmpty(t, responseBody["data"].(map[string]interface{})["errors"])
	})
}
