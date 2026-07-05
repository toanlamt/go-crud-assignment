package services

import "github.com/toanlamt/go-crud-assignment/internal/models"

type ProductService interface {
    Create(models.CreateProductRequest) (*models.Product, error)

    GetAll() ([]models.Product, error)

    GetByID(id int) (*models.Product, error)

    Update(id int, req models.UpdateProductRequest) (*models.Product, error)

    Delete(id int) error
}