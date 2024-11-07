package product_test

import (
	"github.com/lucasti79/bgw4-put-patch-delete/internal/domain"
	"github.com/lucasti79/bgw4-put-patch-delete/internal/product"
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

func (s *ProductTestSuite) Test_MySqlRepositoryWithTxDb_Store_Mock() {
	var p domain.Product

	db, err := utils.InitTxDbDatabase(s.T())
	defer db.Close()
	assert.NoError(s.T(), err)

	repository := product.NewRepository(db)

	err = repository.UpdateOrCreate(&p)
	assert.NoError(s.T(), err)

	assert.True(s.T(), p.Id > 0)
}

func (s *ProductTestSuite) Test_MySqlRepositoryWithTxDb_Delete_Mock() {
	var p domain.Product

	db, err := utils.InitTxDbDatabase(s.T())
	defer db.Close()
	assert.NoError(s.T(), err)

	repository := product.NewRepository(db)

	err = repository.UpdateOrCreate(&p)
	assert.NoError(s.T(), err)

	assert.True(s.T(), p.Id > 0)

	err = repository.Delete(p.Id)
	assert.NoError(s.T(), err)
}
