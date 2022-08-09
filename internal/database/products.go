package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func InsertProducts(cod int, name string, brand string, desc string, price float64, count int) error {
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

func GetProductsByCode(cod int) error {
	db, err := makeCN()
	if err != nil {
		return err
	}

	var product Products

	query := "SELECT * FROM products WHERE cod = $1"
	row := db.QueryRow(query, cod)
	err = row.Scan(&product.Cod, &product.Name, &product.Brand, &product.Description, &product.Price, &product.InventoryCount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return err
		} else {
			return err
		}
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
			if errors.Is(err, sql.ErrNoRows) {
				return nil, err
			} else {
				return nil, err
			}
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

	var query string
	switch columnEdit {
	case "code":
		query = "UPDATE products SET cod = $1 WHERE cod = $2"
	case "name":
		query = "UPDATE products SET name = $1 WHERE cod = $2"
		break
	case "brand":
		query = "UPDATE products SET brand = $1 WHERE cod = $2"
		break
	case "description":
		query = "UPDATE products SET description = $1 WHERE cod = $2"
		break
	case "price":
		query = "UPDATE products SET price = $1 WHERE cod = $2"
		break
	case "inventory count":
		query = "UPDATE products SET inventory_count = $1 WHERE cod = $2"
		break
	}
	_, err = db.Exec(query, newValue, cod)
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