package products

import (
	"github.com/stretchr/testify/assert"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"bytes"
)

func createServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	motor := NewRepository()
	service := NewService(motor)
	handler := NewHandler(service)

	router := gin.Default()
	router.GET("/api/v1/products", handler.GetProducts)

	return router
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder){
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")

	return req, httptest.NewRecorder()
}

func TestGetAllBySellerFunctional(t *testing.T) {
	// Arrange
	var resp []Product
	r := createServer()
	req, rr := createRequestTest(http.MethodGet, "/api/v1/products?seller_id=123", "")

	// Act
	r.ServeHTTP(rr, req)

	// Assert
	assert.Equal(t, 200, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.True(t, len(resp) > 0)
}

func TestGetAllBySellerFail(t *testing.T) {
	// Arrange
	var resp []Product
	r := createServer()
	req, rr := createRequestTest(http.MethodGet, "/api/v1/products", "")

	// Act
	r.ServeHTTP(rr, req)

	// Assert
	assert.Equal(t, 400, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.NotNil(t, err)
}
