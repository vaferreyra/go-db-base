package store

import (
	"database/sql"
	"errors"

	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
	"github.com/go-sql-driver/mysql"
)

var (
	ErrNotFound   = errors.New("There is not product with that id")
	ErrInternal   = errors.New("Something internal has wrong")
	ErrDuplicated = errors.New("Product already exists")
)

const (
	GET_BY_ID = `SELECT id, name, quantity, code_value, is_published, expiration, price 
 			     FROM products 
			     WHERE id = ?`
	INSERT = `INSERT INTO products(name, quantity, code_value, is_published, expiration, price)
			  VALUES ?, ?, ?, ?, ?, ?`
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
		default:
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

func (store *sqlStore) Create(product domain.Product) (er error) {
	statement, err := store.Database.Prepare(INSERT)
	if err != nil {
		er = ErrInternal
		return
	}

	defer statement.Close()

	result, err := statement.Exec(
		product.Name,
		product.Quantity,
		product.CodeValue,
		product.IsPublished,
		product.Expiration,
		product.Price,
	)
	if err != nil {
		driverErr, ok := err.(*mysql.MySQLError)
		if !ok {
			er = ErrInternal
			return
		}

		switch driverErr.Number {
		case 1062:
			er = ErrDuplicated
		default:
			er = ErrInternal
		}
		return
	}

	// Check if product was inserted
	affectedRows, err := result.RowsAffected()
	if err != nil || affectedRows != 1 {
		er = ErrInternal
		return
	}

	// Get product id
	productId, err := result.LastInsertId()
	if err != nil {
		er = ErrInternal
		return
	}

	product.Id = int(productId)
	return

}

func (store *sqlStore) Update(product domain.Product) error { return nil }

func (store *sqlStore) Delete(id int) error { return nil }

func (store *sqlStore) Exists(codeValue string) bool { return false }
