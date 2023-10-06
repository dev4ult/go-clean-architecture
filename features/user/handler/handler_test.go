package handler

import (
	"bytes"
	"clean_arch/features/user/dtos"
	mock "clean_arch/features/user/mocks"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {
	var e = echo.New()

	type Response struct {
		Data	any    `json:"data"`
		Message string `json:"message"`
	}

	var service = mock.MockService{}
	var handler = New(&service)

	t.Run("Success", func(t *testing.T)  {
		var service = mock.SuccessMockService{}
		var handler = New(&service)

		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		res := httptest.NewRecorder()

		e.GET("/users", handler.GetAllUsers())

		e.ServeHTTP(res, req)

		tmp := Response{}

		resData := json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "Success Get Users!", tmp.Message)
		assert.NotNil(t, tmp.Data)
		assert.Nil(t, err)
	})

	t.Run("No User Listed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		res := httptest.NewRecorder()

		e.GET("/users", handler.GetAllUsers())

		e.ServeHTTP(res, req)

		tmp := Response{}

		resData := json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "There is No User!", tmp.Message)
		assert.Empty(t, tmp.Data)
		assert.Nil(t, err)
	})

	t.Run("Error DTO Mapping", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		res := httptest.NewRecorder()

		mock.UsersResponse = append(mock.UsersResponse, dtos.UserResponse{Username: "sarbinus", Email: "sarbin@example.com"})

		e.GET("/users", handler.GetAllUsers())

		e.ServeHTTP(res, req)

		tmp := Response{}

		resData := json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, "Something Went Wrong!", tmp.Message)
		assert.Nil(t, tmp.Data)
		assert.Nil(t, err)
	})
}

func TestCreateUser(t *testing.T) {
	var e = echo.New()

	type Response struct {
		Data	any    `json:"data"`
		Message string `json:"message"`
	}

	var service = mock.MockService{}
	var handler = New(&service)

	t.Run("Success", func(t *testing.T) {
		var service = mock.SuccessMockService{}
		var handler = New(&service)

		reqData := `{"username": "sarbinus", "email": "sarbin@example.com", "password": "sarbin123"}`
		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader([]byte(reqData)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()

		e.POST("/users", handler.CreateUser())

		e.ServeHTTP(res, req)

		tmp := Response{}

		resData := json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "Success User Created!", tmp.Message)
		assert.NotNil(t, tmp.Data)
		assert.Nil(t, err)
	})
	
	t.Run("Error Binding", func(t *testing.T) {
		reqData := `{"username": 1}`
		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader([]byte(reqData)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()

		e.POST("/users", handler.CreateUser())

		e.ServeHTTP(res, req)

		tmp := Response{}

		resData := json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Nil(t, tmp.Data)
		assert.Nil(t, err)
	})

	t.Run("Error DTO Mapping", func(t *testing.T) {
		reqData := `{"name": "sarbinus", "mail": "sarbin@example.com", "pass": "sarbin123"}`
		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader([]byte(reqData)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()

		e.POST("/users", handler.CreateUser())

		e.ServeHTTP(res, req)

		tmp := Response{}

		resData := json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, "Something Went Wrong!", tmp.Message)
		assert.Nil(t, tmp.Data)
		assert.Nil(t, err)
	})

	t.Run("User Already Exist", func(t *testing.T) {
		reqData := `{"username": "sarbinus", "email": "sarbin@example.com", "password": "sarbin123"}`
		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader([]byte(reqData)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()

		e.POST("/users", handler.CreateUser())

		e.ServeHTTP(res, req)

		tmp := Response{}

		resData := json.NewDecoder(res.Result().Body)
		err := resData.Decode(&tmp)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.Equal(t, "User Has Already Exist!", tmp.Message)
		assert.Nil(t, tmp.Data)
		assert.Nil(t, err)
	})
}



