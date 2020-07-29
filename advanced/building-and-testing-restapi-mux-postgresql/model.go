package main

import (
	"database/sql"
	"errors"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *Product) getProducts(db *sql.DB, start, limit int) error {
	return errors.New("Not implemented")
}

func (p *Product) getProduct(db *sql.DB, id int) error {
	return errors.New("Not implemented")
}

func (p *Product) updateProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *Product) deleteProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *Product) createProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}
