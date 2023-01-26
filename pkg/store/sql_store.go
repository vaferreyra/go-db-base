package store

import (
	"database/sql"
	"errors"

	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
)

var (
	ErrNotFound = errors.New("There is not product with that id")
	ErrInternal = errors.New("Something internal has wrong")
)

const (
	GET_BY_ID = `SELECT id, name, quantity, code_value, is_published, expiration, price 
 			     FROM products 
			     WHERE id = ?`
	INSERT = `INSERT INTO products(id, name, quantity, code_value, is_published, expiration, price)
			  VALUES ?, ?, ?, ?, ?, ?, ?`
)

type sqlStore struct {
	Database *sql.DB
}

func NewSQLStore(db *sql.DB) *sqlStore {
	return &sqlStore{db}
}

func (store *sqlStore) Read(id int) (product domain.Product, er error) {
	row := store.Database.QueryRow(GET_BY_ID, id)

	if row.Err() != nil {
		switch row.Err() {
		case sql.ErrNoRows:
			er = ErrNotFound
		case sql.ErrConnDone:
			er = ErrInternal
		}
		return
	}

	if err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Quantity,
		&product.CodeValue,
		&product.IsPublished,
		&product.Expiration,
		&product.Price,
	); err != nil {
		er = ErrInternal
		return
	}
	return
}

func (store *sqlStore) Create(product domain.Product) error {
	statement, err := store.Database.Prepare(INSERT)
	if err != nil {
		return err
	}

	defer statement.Close()

	result, err := statement.Exec()
}

func (store *sqlStore) Update(product domain.Product) error { return nil }

func (store *sqlStore) Delete(id int) error { return nil }

func (store *sqlStore) Exists(codeValue string) bool { return false }
