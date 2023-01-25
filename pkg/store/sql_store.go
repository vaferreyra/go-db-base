package store

import (
	"database/sql"

	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
)

type sqlStore struct {
	Database *sql.DB
}

func NewSQLStore(db *sql.DB) *sqlStore {
	return &sqlStore{db}
}

func (store *sqlStore) Read(id int) (product domain.Product, err error) {
	query := `SELECT 
				id, name, quantity, code_value, is_published, expiration, price 
			 FROM 
			 	products 
			WHERE 
				id = ?`

	rows, err := store.Database.Query(query, id)
	if err != nil {
		return domain.Product{}, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Quantity,
			&product.CodeValue,
			&product.IsPublished,
			&product.Expiration,
			&product.Price,
		); err != nil {
			return domain.Product{}, err
		}
	}
	return
}

func (store *sqlStore) Create(product domain.Product) error {
	return nil
}

func (store *sqlStore) Update(product domain.Product) error { return nil }

func (store *sqlStore) Delete(id int) error { return nil }

func (store *sqlStore) Exists(codeValue string) bool { return false }
