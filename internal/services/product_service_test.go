package services

import (
	"github.com/toanlamt/go-crud-assignment/internal/models"
	"github.com/toanlamt/go-crud-assignment/internal/repositories"
	"testing"
)

var _ repositories.ProductRepository = (*fakeRepository)(nil)
	
type fakeRepository struct {
	product *models.Product
	err     error
}

func (f *fakeRepository) GetByID(id int) (*models.Product, error) {
	return f.product, f.err
}

func (f *fakeRepository) Create(product *models.Product) error {
	return nil
}

func (f *fakeRepository) GetAll() ([]models.Product, error) {
	return nil, nil
}

func (f *fakeRepository) Update(product *models.Product) error {
	return nil
}

func (f *fakeRepository) Delete(id int) error {
	return nil
}

func TestGetByIDSuccess(t *testing.T) {

	repo := &fakeRepository{
		product: &models.Product{
			ID: 1,
			Name: "Keyboard",
		},
	}

	service := NewProductService(repo)

	product, err := service.GetByID(1)

	if err != nil {
		t.Fatal(err)
	}

	if product.Name != "Keyboard" {
		t.Errorf("expected Keyboard, got %s", product.Name)
	}
}