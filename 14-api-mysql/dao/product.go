package dao

import (
	"database/sql"
	"errors"

	"dreadfiles/curso-golang/14-api-mysql/model"
)

// ProductDAO handles CRUD operations for Product.
type ProductDAO struct {
	DB *sql.DB
}

// NewProductDAO creates a new ProductDAO.
func NewProductDAO(db *sql.DB) *ProductDAO {
	return &ProductDAO{DB: db}
}

// Create inserts a new product into the database.
func (dao *ProductDAO) Create(product *model.Product) error {
	query := "INSERT INTO product (id, name, description, value) VALUES (?, ?, ?, ?)"
	_, err := dao.DB.Exec(query, product.ID, product.Name, product.Description, product.Value)
	return err
}

// GetByID retrieves a product by its ID.
func (dao *ProductDAO) GetByID(id int) (*model.Product, error) {
	query := "SELECT id, name, description, value FROM product WHERE id = ?"
	row := dao.DB.QueryRow(query, id)

	var product model.Product
	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Value)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// Update modifies an existing product.
func (dao *ProductDAO) Update(product *model.Product) error {
	query := "UPDATE product SET name = ?, description = ?, value = ? WHERE id = ?"
	res, err := dao.DB.Exec(query, product.Name, product.Description, product.Value, product.ID)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("no rows updated")
	}
	return nil
}

// Delete removes a product by its ID.
func (dao *ProductDAO) Delete(id int) error {
	query := "DELETE FROM product WHERE id = ?"
	res, err := dao.DB.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("no rows deleted")
	}
	return nil
}

// GetAll retrieves all products.
func (dao *ProductDAO) GetAll() ([]*model.Product, error) {
	query := "SELECT id, name, description, value FROM product"
	rows, err := dao.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*model.Product
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Value); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}
