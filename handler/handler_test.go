package handler_test

import (
	"encoding/json"
	"fmt"

	"net/http"

	"net/http/httptest"

	"strings"

	"testing"

	"todo/authentication"
	"todo/handler"

	"todo/models"

	"todo/repository"

	"github.com/gofiber/fiber/v2"

	"github.com/stretchr/testify/assert"
)

var Token string

// testing sign up API
func TestSignup(t *testing.T) {
	db := repository.SetupTestDB()
	c, _ := db.DB()
	defer c.Close()
	app := fiber.New()
	database := handler.Database{Database: db}
	app.Post("/signup", database.Signup)

	t.Run("user already exist", func(t *testing.T) {
		user := models.Information{
			Username:    "vignesh",
			Email:       "kjjk@example.com",
			Password:    "password123",
			PhoneNumber: "1234567890",
		}
		requestBody, _ := json.Marshal(user)
		req := httptest.NewRequest("POST", "/signup", strings.NewReader(string(requestBody)))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusForbidden, resp.StatusCode)
		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, float64(403), response["status"])

	})

	t.Run(" username field missing", func(t *testing.T) {
		user := models.Information{
			Username:    "",
			Email:       "ak@example.com",
			Password:    "password123",
			PhoneNumber: "1234567890",
		}
		requestBody, _ := json.Marshal(user)
		req := httptest.NewRequest("POST", "/signup", strings.NewReader(string(requestBody)))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusForbidden, resp.StatusCode)
		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, float64(403), response["status"])

	})
	t.Run(" email format wrong", func(t *testing.T) {
		user := models.Information{
			Username:    "santhosh",
			Email:       "@example.com",
			Password:    "password123",
			PhoneNumber: "1234567890",
		}
		requestBody, _ := json.Marshal(user)
		req := httptest.NewRequest("POST", "/signup", strings.NewReader(string(requestBody)))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, float64(400), response["status"])

	})
	t.Run(" password charactor should have minimum 8 characters", func(t *testing.T) {
		user := models.Information{
			Username:    "santhosh",
			Email:       "santhosh@example.com",
			Password:    "1234567",
			PhoneNumber: "1234567890",
		}
		requestBody, _ := json.Marshal(user)
		req := httptest.NewRequest("POST", "/signup", strings.NewReader(string(requestBody)))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusForbidden, resp.StatusCode)
		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, float64(403), response["status"])

	})

	t.Run("phone numbers should have 10 digits", func(t *testing.T) {
		user := models.Information{
			Username:    "narayanan",
			Email:       "narayanan@example.com",
			Password:    "12345678",
			PhoneNumber: "123456789",
		}
		requestBody, _ := json.Marshal(user)
		req := httptest.NewRequest("POST", "/signup", strings.NewReader(string(requestBody)))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, float64(400), response["status"])

	})
	t.Run("sign up successful", func(t *testing.T) {
		user := models.Information{
			Username:    "mohan",
			Email:       "blastmohanadd@example.com", //change email while running this test case
			Password:    "1234!2@--",
			PhoneNumber: "8270574897",
		}
		requestBody, _ := json.Marshal(user)
		req := httptest.NewRequest("POST", "/signup", strings.NewReader(string(requestBody)))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, float64(200), response["status"])

	})
}

// testing login API
func TestLogin(t *testing.T) {
	db := repository.SetupTestDB()
	c, _ := db.DB()
	defer c.Close()
	app := fiber.New()
	database := handler.Database{Database: db}
	app.Post("/login", database.Login)
	t.Run("invalid email format", func(t *testing.T) {
		user := models.Information{
			Email:    "@gmail.com",
			Password: "1234!2@--",
		}
		requestBody, _ := json.Marshal(user)
		req := httptest.NewRequest("POST", "/login", strings.NewReader(string(requestBody)))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, float64(400), response["status"])

	})
	t.Run("email not registered", func(t *testing.T) {
		user := models.Information{
			Email:    "gk@example.com",
			Password: "1234!2@--",
		}
		requestBody, _ := json.Marshal(user)
		req := httptest.NewRequest("POST", "/login", strings.NewReader(string(requestBody)))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, float64(400), response["status"])

	})
	t.Run("password not matching", func(t *testing.T) {
		user := models.Information{
			Email:    "ngk@example.com",
			Password: "1234!2@---",
		}
		requestBody, _ := json.Marshal(user)
		req := httptest.NewRequest("POST", "/login", strings.NewReader(string(requestBody)))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusForbidden, resp.StatusCode)
		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, float64(403), response["status"])

	})
	t.Run("login successful", func(t *testing.T) {
		user := models.Information{
			Email:    "muthuvel@example.com",
			Password: "1234!2@--",
		}
		requestBody, _ := json.Marshal(user)
		req := httptest.NewRequest("POST", "/login", strings.NewReader(string(requestBody)))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, float64(200), response["status"])
		a, b := response["token"].(string)
		if b {
			Token = a
		}
		fmt.Println("token", Token)
	})

}

// testing task posting API
func TestTaskPosting(t *testing.T) {
	db := repository.SetupTestDB()
	c, _ := db.DB()
	defer c.Close()

	app := fiber.New()
	database := handler.Database{Database: db}
	app.Post("/posttask", authentication.AuthMiddleware(), database.TaskPosting)

	t.Run("missing token", func(t *testing.T) {
		task := models.TaskDetails{
			TASK_NAME: "Sample Task",
			Status:    "active",
		}
		requestBody, _ := json.Marshal(task)
		req := httptest.NewRequest("POST", "/posttask", strings.NewReader(string(requestBody)))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("valid task posting", func(t *testing.T) {
		token := Token // Generate a valid token here
		task := models.TaskDetails{
			TASK_NAME: "august month",
			Status:    "active",
		}
		requestBody, _ := json.Marshal(task)
		req := httptest.NewRequest("POST", "/posttask", strings.NewReader(string(requestBody)))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, float64(200), response["status"])
		assert.Equal(t, "Task added Successfully", response["message"])
	})
	t.Run("valid task posting", func(t *testing.T) {
		token := Token // Generate a valid token here
		task := models.TaskDetails{
			TASK_NAME: "august month",
			Status:    "active",
		}
		requestBody, _ := json.Marshal(task)
		req := httptest.NewRequest("POST", "/posttask", strings.NewReader(string(requestBody)))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, float64(200), response["status"])
		assert.Equal(t, "Task added Successfully", response["message"])
	})
	t.Run("missing task name", func(t *testing.T) {
		token := Token
		task := models.TaskDetails{
			Status: "active",
		}
		requestBody, _ := json.Marshal(task)
		req := httptest.NewRequest("POST", "/posttask", strings.NewReader(string(requestBody)))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("in status field only active or completed should be given", func(t *testing.T) {
		token := Token
		task := models.TaskDetails{
			TASK_NAME: "Sample Task",
			Status:    "invalid_status",
		}
		requestBody, _ := json.Marshal(task)
		req := httptest.NewRequest("POST", "/posttask", strings.NewReader(string(requestBody)))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("invalid token", func(t *testing.T) {
		invalidToken := "JlbWFpbCI6Im5na0BleGFtcGxlLmNvbSIsImV4cCI6MTY5MTc0MDIxNiwicGFzc3dvcmQiOiIxMjM0ITJALS0iLCJ1c2VyX2lkIjoiMzYifQ.5y5nSmRIzd6m8UigIun4RCMwNNMf0csPOGhF1Vi2MgY"
		task := models.TaskDetails{
			TASK_NAME: "Sample Task",
			Status:    "active",
		}
		requestBody, _ := json.Marshal(task)
		req := httptest.NewRequest("POST", "/posttask", strings.NewReader(string(requestBody)))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+invalidToken)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})
}

// testing task posting API for updating
func TestUpdateTask(t *testing.T) {
	db := repository.SetupTestDB()
	c, _ := db.DB()
	defer c.Close()
	app := fiber.New()
	database := handler.Database{Database: db}
	app.Put("/updatetask/:id", authentication.AuthMiddleware(), database.UpdateTask)
	t.Run("missing token", func(t *testing.T) {
		task := models.TaskDetails{
			TASK_NAME: "Sample Task",
			Status:    "active",
		}
		requestBody, _ := json.Marshal(task)
		req := httptest.NewRequest("PUT", "/updatetask/1", strings.NewReader(string(requestBody)))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("invalid token", func(t *testing.T) {
		// Generate an invalid token
		invalidToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFm11dGh1dmVsQGV4YW1wbGUuY29tIiwiZXhwIjoxNjkxNzYwNTY5LCJwYXNzd29yZCI6IjEyMzQhMkAtLSIsInVzZXJfaWQiOiIzNyJ9.nUL8-n1ayUbxRmYnoTzeGtyrni718CZ-xrMdYi86Q1I"
		task := models.TaskDetails{
			TASK_NAME: "Updated Task",
			Status:    "completed",
		}
		requestBody, _ := json.Marshal(task)
		req := httptest.NewRequest("PUT", "/updatetask/4", strings.NewReader(string(requestBody)))
		req.Header.Set("Authorization", "Bearer "+invalidToken)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("task not found", func(t *testing.T) {
		token := Token // Generate a valid token here
		task := models.TaskDetails{
			TASK_NAME: "Updated Task",
			Status:    "completed",
		}
		requestBody, _ := json.Marshal(task)
		req := httptest.NewRequest("PUT", "/updatetask/9", strings.NewReader(string(requestBody)))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})

	t.Run("successful updation", func(t *testing.T) {
		token := Token // Generate a valid token here
		task := models.TaskDetails{
			TASK_NAME: "fifty nine",
			Status:    "59",
		}
		requestBody, _ := json.Marshal(task)
		req := httptest.NewRequest("PUT", "/updatetask/59", strings.NewReader(string(requestBody)))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

// testing task deleting API
func TestDeleteTask(t *testing.T) {
	db := repository.SetupTestDB()
	c, _ := db.DB()
	defer c.Close()
	app := fiber.New()
	database := handler.Database{Database: db}
	app.Delete("/deletetask/:id", authentication.AuthMiddleware(), database.DeleteTask)
	t.Run("missing token", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/deletetask/1", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})
	t.Run("invalid token", func(t *testing.T) {
		// Generate an invalid token
		invalidToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFm11dGh1dmVsQGV4YW1wbGUuY29tIiwiZXhwIjoxNjkxNzYwNTY5LCJwYXNzd29yZCI6IjEyMzQhMkAtLSIsInVzZXJfaWQiOiIzNyJ9.nUL8-n1ayUbxRmYnoTzeGtyrni718CZ-xrMdYi86Q1I"
		req := httptest.NewRequest("DELETE", "/deletetask/6", nil) //change id while running this particular test case
		req.Header.Set("Authorization", "Bearer "+invalidToken)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("task not found", func(t *testing.T) {
		token := Token // Generate a valid token here
		req := httptest.NewRequest("DELETE", "/deletetask/9", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})

	t.Run("successful deletion", func(t *testing.T) {
		token := Token // Generate a valid token here
		req := httptest.NewRequest("DELETE", "/deletetask/58", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

// testing task posting date filtration API
func TestGetTasksByUserAndDate(t *testing.T) {
	db := repository.SetupTestDB()
	c, _ := db.DB()
	defer c.Close()
	app := fiber.New()
	database := handler.Database{Database: db}
	app.Get("/gettasksbydate", authentication.AuthMiddleware(), database.GetTasksByUserAndDate)
	t.Run("missing token", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/gettasksbydate?startDate=01-01-2023&endDate=31-12-2023&status=active", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})
	t.Run("invalid token", func(t *testing.T) {
		// Generate an invalid token
		invalidToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpCJ9.eyJlbWFm11dGh1dmVsQGV4YW1wbGUuY29tIiwiZXhwIjoxNjkxNzYwNTY5LCJwYXNzd29yZCI6IjEyMzQhMkAtLSIsInVzZXJfaWQiOiIzNyJ9.nUL8-n1ayUbxRmYnoTzeGtyrni718CZ-xrMdYi86Q1I"
		req := httptest.NewRequest("GET", "/gettasksbydate?startDate=01-08-2023&endDate=31-08-2023&status=active", nil)
		req.Header.Set("Authorization", "Bearer "+invalidToken)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	})
	t.Run("without startDate", func(t *testing.T) {
		token := Token // Generate a valid token here
		req := httptest.NewRequest("GET", "/gettasksbydate?endDate=01-08-2023&status=active", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("missing endDate", func(t *testing.T) {
		token := Token // Generate a valid token here
		req := httptest.NewRequest("GET", "/gettasksbydate?startDate=01-01-2023&status=active", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})
	t.Run("tasks by user and date", func(t *testing.T) {
		token := Token // Generate a valid token here
		req := httptest.NewRequest("GET", "/gettasksbydate?startDate=01-01-2023&endDate=31-12-2023", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
	t.Run("tasks by user, date, and status", func(t *testing.T) {
		token := Token // Generate a valid token here
		req := httptest.NewRequest("GET", "/gettasksbydate?startDate=01-01-2023&endDate=31-12-2023&status=active", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
	t.Run("in status field only active and completed should be given", func(t *testing.T) {
		token := Token // Generate a valid token here
		req := httptest.NewRequest("GET", "/gettasksbydate?startDate=01-01-2023&endDate=31-12-2023&status=activess", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})
	t.Run("in status field only active and completed should be given", func(t *testing.T) {
		token := Token // Generate a valid token here
		req := httptest.NewRequest("GET", "/gettasksbydate?status=com", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})
	t.Run("tasks by status", func(t *testing.T) {
		token := Token // Generate a valid token here
		req := httptest.NewRequest("GET", "/gettasksbydate?status=active", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
	t.Run("tasks by user, date,to check not assigned task  ", func(t *testing.T) {
		token := Token // Generate a valid token here
		req := httptest.NewRequest("GET", "/gettasksbydate?startDate=01-01-2024&endDate=31-12-2024", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})
	t.Run("tasks by user, date, and status, to check not assigned task ", func(t *testing.T) {
		token := Token // Generate a valid token here
		req := httptest.NewRequest("GET", "/gettasksbydate?startDate=01-01-2024&endDate=31-12-2024&status=completed", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})
}
