package main

import (
	"database/sql"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func getProducts(db *sql.DB, start, limit int) ([]Product, error) {
	rows, err := db.Query("SELECT id, name, price FROM product LIMIT $1 OFFSET $2", limit, start)

	if err == nil {
		return nil, err
	}

	defer rows.Close()
	products := []Product{}

	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}

func (p *Product) getProduct(db *sql.DB, id int) error {
	return db.QueryRow("SELECT name, price FROM product WHERE id=$1", id).Scan(&p.Name, &p.Price)
}

func (p *Product) updateProduct(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE product SET name=$1, price=$2 WHERE id=$3",
			p.Name, p.Price, p.ID)

	return err
}

func (p *Product) deleteProduct(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM product WHERE id=$1", id)
	return err
}

func (p *Product) createProduct(db *sql.DB) error {
	err := db.QueryRow("INSERT INTO product(name, price) VALUES($1, $2) RETURNING id", p.Name, p.Price).Scan(&p.ID)

	if err != nil {
		return err
	}

	return nil
}
