package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Shreyas-Prabhu/EmployeeDatabase/models"
	"github.com/Shreyas-Prabhu/EmployeeDatabase/router"
	"github.com/stretchr/testify/assert"
)

const insertID = 111111
const notExistentID = 222222

func TestInsertController(t *testing.T) {
	t.Run("Insert Employee record", func(t *testing.T) {
		r := router.NewRouter()
		emp := models.Employee{ID: insertID, Name: "Shreyas", Position: "Go dev", Salary: 1232431} //using dummy ID
		byteData, _ := json.Marshal(emp)
		req, _ := http.NewRequest("POST", "/employee/insert", bytes.NewBuffer(byteData))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert := assert.New(t)
		expectedBody := `{"message":"Employee inserted"}`

		assert.Equal(w.Code, 201)
		assert.Equal(w.Body.String(), expectedBody)
	})

	t.Run("Insert Employee record when id already existing", func(t *testing.T) {
		r := router.NewRouter()
		emp := models.Employee{ID: insertID, Name: "Alex", Position: "Go developer", Salary: 1232431}
		byteData, _ := json.Marshal(emp)
		req, _ := http.NewRequest("POST", "/employee/insert", bytes.NewBuffer(byteData))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert := assert.New(t)
		expectedBody := `{"message":"Employee already present with ID:111111"}`

		assert.Equal(w.Code, 409)
		assert.Equal(w.Body.String(), expectedBody)
	})

	t.Run("Returns 422 when all data is not populated", func(t *testing.T) {
		r := router.NewRouter()
		emp := models.Employee{ID: notExistentID, Position: "Go developer", Salary: 1232431}
		byteData, _ := json.Marshal(emp)
		req, _ := http.NewRequest("POST", "/employee/insert", bytes.NewBuffer(byteData))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert := assert.New(t)
		expectedBody := `{"message":"Some fields are empty"}`

		assert.Equal(w.Code, 422)
		assert.Equal(w.Body.String(), expectedBody)
	})
}

func TestGetController(t *testing.T) {
	t.Run("Get Employee Record", func(t *testing.T) {
		r := router.NewRouter()
		endpoint := fmt.Sprint("/employee/get/", insertID)
		req, _ := http.NewRequest("GET", endpoint, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert := assert.New(t)
		assert.Equal(w.Code, 200)
	})

	t.Run("Return 404 if no record found", func(t *testing.T) {
		r := router.NewRouter()
		endpoint := fmt.Sprint("/employee/get/", notExistentID)
		req, _ := http.NewRequest("GET", endpoint, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert := assert.New(t)
		assert.Equal(w.Code, 404)
	})
}

func TestUpdateController(t *testing.T) {
	t.Run("Update Employee Record", func(t *testing.T) {
		r := router.NewRouter()
		emp := models.Employee{ID: insertID, Name: "Shreyas Prabhu", Position: "Go developer", Salary: 1232431} //using dummy ID
		byteData, _ := json.Marshal(emp)
		endpoint := fmt.Sprint("/employee/update/", insertID)
		req, _ := http.NewRequest("PUT", endpoint, bytes.NewBuffer(byteData))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert := assert.New(t)

		assert.Equal(w.Code, 200)
	})

	t.Run("Return 404 if no record found", func(t *testing.T) {
		r := router.NewRouter()
		emp := models.Employee{ID: notExistentID, Name: "Shreyas Prabhu", Position: "Go developer", Salary: 1232431} //using dummy ID
		byteData, _ := json.Marshal(emp)
		endpoint := fmt.Sprint("/employee/update/", notExistentID)
		req, _ := http.NewRequest("PUT", endpoint, bytes.NewBuffer(byteData))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert := assert.New(t)

		assert.Equal(w.Code, 404)
	})
}

func TestDeleteController(t *testing.T) {
	t.Run("Return 404 if Employee Record not found", func(t *testing.T) {
		r := router.NewRouter()
		endpoint := fmt.Sprint("/employee/delete/", notExistentID)
		req, _ := http.NewRequest("DELETE", endpoint, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert := assert.New(t)
		expectedBody := `{"message":"Employee does not exist with ID:222222"}`

		assert.Equal(w.Code, 404)
		assert.Equal(w.Body.String(), expectedBody)
	})

	t.Run("Delete Employee record", func(t *testing.T) {
		r := router.NewRouter()
		endpoint := fmt.Sprint("/employee/delete/", insertID)
		req, _ := http.NewRequest("DELETE", endpoint, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert := assert.New(t)
		expectedBody := `{"message":"Employee deleted with ID:111111"}`

		assert.Equal(w.Code, 200)
		assert.Equal(w.Body.String(), expectedBody)
	})
}

func TestPaginationController(t *testing.T) {
	t.Run("Return 200 for successful paginated list", func(t *testing.T) {
		r := router.NewRouter()
		req, _ := http.NewRequest("GET", "/employee/getList?page=1&size=2", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert := assert.New(t)

		assert.Equal(w.Code, 200)
	})
}
