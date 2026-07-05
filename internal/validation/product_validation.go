package validation

import (
    "errors"
    "strings"

    "github.com/toanlamt/go-crud-assignment/internal/models"
)

func ValidateCreateProduct(
    req models.CreateProductRequest,
) error {

    if strings.TrimSpace(req.Name) == "" {
        return errors.New("name is required")
    }

    if len(req.Name) < 3 {
        return errors.New("name must be at least 3 characters")
    }

    if req.Price <= 0 {
        return errors.New("price must be greater than 0")
    }

    if req.Quantity < 0 {
        return errors.New("quantity must be greater than or equal to 0")
    }

    return nil
}