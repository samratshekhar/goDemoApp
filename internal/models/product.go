package models

import (
	"database/sql"
	"errors"
)

type Product struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

func (p *Product) GetProduct(db *sql.DB) error {
	return db.QueryRow("SELECT name, price FROM products WHERE id=?",
		p.Id).Scan(&p.Name, &p.Price)
}

func (p *Product) CreateProduct(db *sql.DB) error {
	_, err := db.Exec(
		"INSERT INTO products(name, price) VALUES(?, ?)",
		p.Name, p.Price)

	if err != nil {
		return err
	}

	return nil
}

func (p *Product) UpdateProduct(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE products SET name=?, price=? WHERE id=?",
			p.Name, p.Price, p.Id)

	return err
}

func (p *Product) DeleteProduct(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM products WHERE id=?", p.Id)
	return err
}

func (p *Product) getProducts(db *sql.DB) ([]Product, error) {
	return nil, errors.New("Not implemented")
}
