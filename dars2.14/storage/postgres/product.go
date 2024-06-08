package postgres

import (
	"database/sql"
	"n11/Firdavs/dars2.14/model"
)

type ProductRepo struct {
	db *sql.DB
}

func Product_NewRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

func (u *ProductRepo) Product_Get(id int) (model.Product, error) {
	var product model.Product
	err := u.db.QueryRow("SELECT id, name, description, price, stock_quantity FROM products WHERE id = $1", id).Scan(
		&product.Id, &product.Name, &product.Description, &product.Price, &product.Stock_quantity)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (u *ProductRepo) Product_delete(id string) error {
	_, err := u.db.Exec(`delete from products where id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *ProductRepo) Product_Create(product model.Product) error {
	_, err := u.db.Exec(`INSERT INTO products (name, description, price, stock_quantity) 
                       VALUES ($1, $2, $3, $4)`,
		product.Name, product.Description, product.Price, product.Stock_quantity)
	if err != nil {
		return err
	}

	return nil
}
func (repo *ProductRepo) Product_Update(id int, product model.Product) error {
	query := `UPDATE products SET name=$1, description=$2, price=$3, stock_quantity=$4 WHERE id=$5`
	_, err := repo.db.Exec(query, product.Name, product.Description, product.Price, product.Stock_quantity, id)
	if err != nil {
		return err
	}
	return nil
}
