package handler

import (
	"bytes"
	"clase3_parte2/internal/domain"
	"clase3_parte2/test/mocks"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer(mockService mocks.MockService) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	p := NewProduct(&mockService)

	r := gin.Default()

	pr := r.Group("/products")
	pr.POST("", p.Store())
	pr.GET("", p.GetAll())
	pr.PUT("/:id", p.UpdateAll())

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func TestGetAllStatusOk(t *testing.T) {
	// arrange
	mockService := mocks.MockService{
		DataMock: []domain.Product{},
		Error:    "",
	}
	var resp []domain.Product
	r := createServer(mockService)
	req, rr := createRequestTest(http.MethodGet, "/products", "")
	// act
	r.ServeHTTP(rr, req)
	// assert
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, mockService.DataMock, resp)
}

func TestGetAllStatusInternalServerError(t *testing.T) {
	// arrange
	mockService := mocks.MockService{
		DataMock: []domain.Product{},
		Error:    "error ocurred",
	}
	var resp map[string]string
	expected := map[string]string{
		"error": "something went wrong",
	}
	r := createServer(mockService)
	req, rr := createRequestTest(http.MethodGet, "/products", "")
	// act
	r.ServeHTTP(rr, req)
	// assert
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, 500, rr.Code)
	assert.Equal(t, expected, resp)
}

func TestStoreStatusCreated(t *testing.T) {
	// arrange
	mockService := mocks.MockService{
		DataMock: []domain.Product{},
		Error:    "",
	}
	expected := domain.Product{
		ID:    0,
		Name:  "Tester",
		Type:  "Funcional",
		Count: 10,
		Price: 9,
	}
	var resp domain.Product
	r := createServer(mockService)
	req, rr := createRequestTest(http.MethodPost, "/products", `{
		"nombre": "Tester","tipo": "Funcional","cantidad": 10,"precio": 9
		}`)
	// act
	r.ServeHTTP(rr, req)
	// assert
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, 201, rr.Code)
	assert.Equal(t, expected, resp)
}


func TestPutOk(t *testing.T) {
	//arrange
	beforeProduct := domain.Product{
		ID:    1,
		Name:  "Mi Producto",
		Type:  "Imaginario",
		Count: 10,
		Price: 100,
	}
	afterUpdate := domain.Product{
		ID:    1,
		Name:  "Mi Producto",
		Type:  "Imaginario",
		Count: 1,
		Price: 90,
	}
	mockService := mocks.MockService{
		DataMock: []domain.Product{
			beforeProduct,
		},
		Error:    "",
	}
	r := createServer(mockService)
	req, rr := createRequestTest(http.MethodPut, "/products/1", `{
		"nombre": "Mi Producto","tipo": "Imaginario","cantidad": 1,"precio": 90
		}`)
	//act
	var resp domain.Product
	r.ServeHTTP(rr, req)
	//assert
	assert.Equal(t, http.StatusOK, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, afterUpdate, resp)
}

func TestUpdateOk(t *testing.T) {
	//arrange
	beforeProduct := domain.Product{
		ID:    1,
		Name:  "Mi Producto",
		Type:  "Imaginario",
		Count: 10,
		Price: 100,
	}
	afterUpdate := domain.Product{
		ID:    1,
		Name:  "Mi Producto",
		Type:  "Imaginario",
		Count: 1,
		Price: 90,
	}
	mockService := mocks.MockService{
		DataMock: []domain.Product{
			beforeProduct,
		},
		Error:    "",
	}
	r := createServer(mockService)
	req, rr := createRequestTest(http.MethodPut, "/products/1", `{
		"nombre": "Mi Producto","tipo": "Imaginario","cantidad": 1,"precio": 90
		}`)
	//act
	var resp domain.Product
	r.ServeHTTP(rr, req)
	//assert
	assert.Equal(t, http.StatusOK, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, afterUpdate, resp)
}

func TestDeleteOk(t *testing.T) {
	//arrange
	beforeProduct := []domain.Product{
		{
			ID:    1,
			Name:  "Mi Producto",
			Type:  "Imaginario",
			Count: 10,
			Price: 100,
		},
		{
			ID:    2,
			Name:  "Producto 2",
			Type:  "Imaginario",
			Count: 19,
			Price: 85,	
		},
	}
	afterUpdate := domain.Product{
		ID:    1,
		Name:  "Mi Producto",
		Type:  "Imaginario",
		Count: 1,
		Price: 90,
	}
	mockService := mocks.MockService{
		DataMock: beforeProduct,
		Error:    "",
	}
	r := createServer(mockService)
	req, rr := createRequestTest(http.MethodDelete, "/products/2", "")
	//act
	r.ServeHTTP(rr, req)
	//assert
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, afterUpdate, mockService.DataMock, "Was not deleted")
}