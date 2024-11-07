package products_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lucasti79/bgw4-put-patch-delete/tests/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ProductTestSuite struct {
	suite.Suite
}

func (suite *ProductTestSuite) SetupTest() {
	utils.RegisterTxDbDatabase(suite.T())
}

// FIRST

func (suite *ProductTestSuite) TestProductsGetAll_Success() {
	// given
	db, err := utils.InitTxDbDatabase(suite.T())
	defer db.Close()
	assert.Nil(suite.T(), err)

	server := utils.CreateServerTest(suite.T(), db)
	req, res := createRequestTest(suite.T(), http.MethodPut, "/api/products/1", `
		{
			"name": "product 1",
			"type": "product type 1",
			"quantity": 10,
			"price": 20.00
		}	
	`)
	req = utils.WithUrlParam(suite.T(), req, "id", "1")

	server.ServeHTTP(res, req)
	assert.Equal(suite.T(), http.StatusOK, res.Code)
	req, res = createRequestTest(suite.T(), http.MethodGet, "/api/products", "")

	// when
	server.ServeHTTP(res, req)

	// then
	assert.Equal(suite.T(), http.StatusOK, res.Code)
}

func TestProductsGetAll(t *testing.T) {
	// buscar todos os produtos
	suite.Run(t, new(ProductTestSuite))
}

func createRequestTest(t *testing.T, method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	t.Helper()
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")
	return req, httptest.NewRecorder()
}
