package database

import (
	"fmt"
)

type Products struct {
	Cod            int
	Name           string
	Brand          string
	Description    string
	Price          float32
	InventoryCount int
}

func InsertProducts(cod int, name string, brand string, desc string, price float32, count int) error {
	db, err := makeCN()

	query := "INSERT INTO products (cod,name,brand,description,price,inventory_count) VALUES ($1,$2,$3,$4,$5,$6)"
	result, err := db.Exec(query, cod, name, brand, desc, price, count)
	if err != nil {

		return err
	}

	_, err = result.RowsAffected()
	if err != nil {

		return err
	}

	return nil
}

func GetProductsByID(id int) error {
	db, err := makeCN()
	if err != nil {
		return err
	}

	var product Products

	query := "SELECT * FROM products WHERE id = $1"
	row := db.QueryRow(query, id).Scan(&product.Cod)
	if row.Error() != "" {
		return err
	}

	fmt.Printf("Product: %s\nCod: %d\nBrand: %s\nDescription: %s\nPrice: %.2f\nInventory count: %d\n", product.Name, product.Cod, product.Brand, product.Description, product.Price, product.InventoryCount)

	return nil
}

func GetAllProducts() ([]*Products, error) {
	db, err := makeCN()
	if err != nil {
		return nil, err
	}

	products := []*Products{}

	query := "SELECT * FROM products"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		p := &Products{}
		err := rows.Scan(&p.Cod, &p.Name, &p.Brand, &p.Description, &p.Price, &p.InventoryCount)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func UpdateProducts(columnEdit string, newValue any, cod int) error {
	db, err := makeCN()
	if err != nil {
		return err
	}

	query := "UPDATE products SET $1 = $2 WHERE cod = $3"
	_, err = db.Exec(query, columnEdit, newValue, cod)
	if err != nil {
		return err
	}

	return nil
}

func DeleteProducts(cod int) error {
	db, err := makeCN()
	if err != nil {
		return err
	}

	query := "DELETE FROM products WHERE cod = $1"
	_, err = db.Exec(query, cod)
	if err != nil {
		return err
	}

	return nil
}