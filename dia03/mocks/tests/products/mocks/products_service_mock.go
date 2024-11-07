package mocks

import (
	"github.com/lucasti79/bgw4-put-patch-delete/internal/domain"
	"github.com/stretchr/testify/mock"
)

type ProductsRepositoryMock struct {
	mock.Mock
}

func (service *ProductsRepositoryMock) Get() (p []domain.Product, err error) {
	args := service.Called()
	arg0, ok := args.Get(0).([]domain.Product)
	if !ok {
		return []domain.Product{}, args.Error(1)
	}

	return arg0, args.Error(1)
}

func (service *ProductsRepositoryMock) GetByID(id int) (p *domain.Product, err error) {
	return nil, nil
}
func (service *ProductsRepositoryMock) UpdateOrCreate(p *domain.Product) (err error) {
	args := service.Called(p)
	return args.Error(0)
}
func (service *ProductsRepositoryMock) Update(id int, p *domain.Product) (err error) {
	return nil
}
func (service *ProductsRepositoryMock) Delete(id int) (err error) {
	return nil
}
