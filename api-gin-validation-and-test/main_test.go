package main

import (
	"api-with-gin/controllers"
	"api-with-gin/database"
	"api-with-gin/models"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func RouteSetupTest() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	routes := gin.Default()

	return routes
}

func CreateStudentMock() {
	student := models.Student{Name: "Jonh Doe", CPF: "12345678901", RG: "123456789"}

	database.DB.Create(&student)

	ID = int(student.ID)
}

func RemoveStudentMock() {
	var student models.Student

	database.DB.Delete(&student, ID)
}

func TestStatusCodeVerifyGreetingRouteWithParam(t *testing.T) {
	r := RouteSetupTest()
	r.GET("/:name", controllers.Gretting)

	req, _ := http.NewRequest("GET", "/davi", nil)

	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "should be equals")

	mockResponse := `{"APY says:":"Hello davi, are you ok ?"}`

	responseBody, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, mockResponse, string(responseBody))
}

func TestIndexHandler(t *testing.T) {
	database.ConnectDB()

	CreateStudentMock()

	defer RemoveStudentMock()

	r := RouteSetupTest()
	r.GET("/students", controllers.Index)

	req, _ := http.NewRequest("GET", "/students", nil)

	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestGetByCPFHandler(t *testing.T) {
	database.ConnectDB()

	CreateStudentMock()

	defer RemoveStudentMock()

	r := RouteSetupTest()
	r.GET("/students/cpf/:cpf", controllers.GetByCPF)

	req, _ := http.NewRequest("GET", "/students/cpf/12345678901", nil)

	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestFindByIdHandler(t *testing.T) {
	database.ConnectDB()

	CreateStudentMock()

	defer RemoveStudentMock()

	r := RouteSetupTest()
	r.GET("/students/:id", controllers.FindById)

	searchPath := "/students/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("GET", searchPath, nil)

	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	var studentMock models.Student

	json.Unmarshal(response.Body.Bytes(), &studentMock)

	assert.Equal(t, "Jonh Doe", studentMock.Name)
	assert.Equal(t, "12345678901", studentMock.CPF)
	assert.Equal(t, "123456789", studentMock.RG)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestDeleteHandler(t *testing.T) {
	database.ConnectDB()

	CreateStudentMock()

	r := RouteSetupTest()
	r.DELETE("/students/:id", controllers.Delete)

	searchPath := "/students/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("DELETE", searchPath, nil)

	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestUpdateHandler(t *testing.T) {
	database.ConnectDB()

	CreateStudentMock()
	defer RemoveStudentMock()

	r := RouteSetupTest()
	r.PATCH("/students/:id", controllers.Edit)

	student := models.Student{Name: "Jonh Doe", CPF: "42345678901", RG: "123456777"}

	jsonParsed, _ := json.Marshal(student)

	path := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(jsonParsed))

	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	var studentMock models.Student

	json.Unmarshal(response.Body.Bytes(), &studentMock)

	assert.Equal(t, "42345678901", studentMock.CPF)
	assert.Equal(t, "123456777", studentMock.RG)
}
