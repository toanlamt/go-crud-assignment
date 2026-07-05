package repositories

import "github.com/toanlamt/go-crud-assignment/internal/models"

type ProductRepository interface {
	Create(product *models.Product) error

	GetAll() ([]models.Product, error)

	GetByID(id int) (*models.Product, error)

	Update(product *models.Product) error

	Delete(id int) error
}