package product_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lucasti79/bgw4-put-patch-delete/internal/domain"
	"github.com/lucasti79/bgw4-put-patch-delete/internal/product"
	"github.com/lucasti79/bgw4-put-patch-delete/tests/utils"
	"github.com/stretchr/testify/assert"
)

func Test_MySqlRepositoryWithSqlMock_GetOne_Mock(t *testing.T) {
	productId := 1
	db, mock := utils.InitSqlMockDatabase(t)
	defer db.Close()

	columns := []string{
		"id",
		"name",
		"price",
		"quantity",
		"type",
	}

	rows := sqlmock.NewRows(columns)
	rows.AddRow(productId, "", 0.0, 0, "")

	mock.
		ExpectQuery(regexp.QuoteMeta(product.GetById)).
		WithArgs(productId).
		WillReturnRows(rows)

	repository := product.NewRepository(db)
	result, err := repository.GetByID(productId)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, productId, result.Id)
}

func Test_MySqlRepositoryWithSqlMock_Store_Mock(t *testing.T) {
	createProduct := domain.Product{
		Id: 1,
	}

	db, mock := utils.InitSqlMockDatabase(t)
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta(product.Create)).
		WithArgs(
			createProduct.Name,
			createProduct.Price,
			createProduct.Quantity,
			createProduct.Type,
		).WillReturnResult(sqlmock.NewResult(1, 1))

	repository := product.NewRepository(db)
	repository.UpdateOrCreate(&createProduct)
}

func Test_MySqlRepositoryWithSqlMock_Delete_Mock(t *testing.T) {
	productId := 1
	db, mock := utils.InitSqlMockDatabase(t)
	defer db.Close()

	mock.ExpectPrepare(regexp.QuoteMeta(product.Delete))

	mock.ExpectExec(regexp.QuoteMeta(product.Delete)).
		WithArgs(productId).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repository := product.NewRepository(db)
	err := repository.Delete(productId)

	assert.NoError(t, err)
}
