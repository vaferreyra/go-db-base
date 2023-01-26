package store

import (
	"errors"

	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
)

var (
	ErrNotFound   = errors.New("There is not product with that id")
	ErrInternal   = errors.New("Something internal has wrong")
	ErrDuplicated = errors.New("Product already exists")
)

type StoreInterface interface {
	// Read devuelve un producto por su id
	Read(id int) (domain.Product, error)
	// Create agrega un nuevo producto
	Create(product domain.Product) error
	// Update actualiza un producto
	Update(product domain.Product) error
	// Delete elimina un producto
	Delete(id int) error
	// Exists verifica si un producto existe
	Exists(codeValue string) bool
}
