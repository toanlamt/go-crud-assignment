package services

import (
	"errors"
	"strings"
	"github.com/toanlamt/go-crud-assignment/internal/models"
	"github.com/toanlamt/go-crud-assignment/internal/repositories"
)

type ProductService struct {
	repository *repositories.ProductRepository
}

func NewProductService(repository *repositories.ProductRepository) *ProductService {
	return &ProductService{
		repository: repository,
	}
}

func (s *ProductService) GetAll() ([]models.Product, error) {
	return s.repository.GetAll()
}

func (s *ProductService) GetByID(id int) (*models.Product, error) {
	return s.repository.GetByID(id)
}

func (s *ProductService) Create(req models.CreateProductRequest) (*models.Product, error) {
	if strings.TrimSpace(req.Name) == "" {
		return nil, errors.New("name is required")
	}

	if len(req.Name) < 3 {
		return nil, errors.New("name must be at least 3 characters")
	}

	if req.Price <= 0 {
		return nil, errors.New("price must be greater than 0")
	}

	if req.Quantity < 0 {
		return nil, errors.New("quantity must be greater than or equal to 0")
	}

	product := &models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Quantity:    req.Quantity,
	}

	if err := s.repository.Create(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) Update(id int, req models.UpdateProductRequest) (*models.Product, error) {

	if strings.TrimSpace(req.Name) == "" {
		return nil, errors.New("name is required")
	}

	if len(req.Name) < 3 {
		return nil, errors.New("name must be at least 3 characters")
	}

	if req.Price <= 0 {
		return nil, errors.New("price must be greater than 0")
	}

	if req.Quantity < 0 {
		return nil, errors.New("quantity must be greater than or equal to 0")
	}

	product, err := s.repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.Quantity = req.Quantity

	err = s.repository.Update(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) Delete(id int) error {
	return s.repository.Delete(id)
}