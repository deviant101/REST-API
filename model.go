package main

import (
	"database/sql"
	"errors"
)

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

func getProducts(db *sql.DB) ([]Product, error) {

	rows, err := db.Query("SELECT id, name, quantity, price FROM items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []Product{}
	for rows.Next() {
		var tempProduct Product
		if err := rows.Scan(&tempProduct.ID, &tempProduct.Name, &tempProduct.Quantity, &tempProduct.Price); err != nil {
			return nil, err
		}
		products = append(products, tempProduct)
	}
	return products, nil
}

func (p *Product) getProduct(db *sql.DB) error {

	row := db.QueryRow("SELECT name, quantity, price FROM items WHERE id=?", p.ID)
	err := row.Scan(&p.Name, &p.Quantity, &p.Price)
	if err != nil {
		return err
	}
	return nil
}

func (p *Product) createProduct(db *sql.DB) error {

	result, err := db.Exec("INSERT INTO items(name, quantity, price) VALUES(?, ?, ?)", p.Name, p.Quantity, p.Price)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	p.ID = int(id)
	return nil
}

func (p *Product) updateProduct(db *sql.DB) error {

	result, err := db.Exec("UPDATE items SET name=?, quantity=?, price=? WHERE id=?", p.Name, p.Quantity, p.Price, p.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("Product with this id does not exist")
	}
	return err
}

func (p *Product) deleteProduct(db *sql.DB) error {

	result, err := db.Exec("DELETE FROM items WHERE id=?", p.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("Product with this id does not exist")
	}
	return err
}
