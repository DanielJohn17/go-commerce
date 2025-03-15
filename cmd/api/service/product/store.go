package product

import (
	"database/sql"

	"github.com/DanielJohn17/go-commerce/cmd/api/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db}
}

func (s *Store) GetProducts() ([]types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	products := make([]types.Product, 0)
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil
}

func scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *Store) CreateProduct(p types.Product) (types.Product, error) {
	result, err := s.db.Exec("INSERT INTO products (name, description, image, price, quantity) VALUES (?, ?, ?, ?, ?)",
		p.Name, p.Description, p.Image, p.Price, p.Quantity)
	if err != nil {
		return types.Product{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return types.Product{}, err
	}

	return types.Product{
		ID:          int(id),
		Name:        p.Name,
		Description: p.Description,
		Image:       p.Image,
		Price:       p.Price,
		Quantity:    p.Quantity,
	}, nil
}
