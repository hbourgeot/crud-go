package database

import (
	"fmt"
	"log"
)

type Products struct {
	Cod            int
	Name           string
	Brand          string
	Description    string
	Price          float32
	InventoryCount int
}

func InsertProducts(cod int, name string, brand string, desc string, price float32, count int) {
	db, err := makeCN()

	query := "INSERT INTO products (cod,name,brand,description,price,inventory_count) VALUES ($1,$2,$3,$4,$5,$6)"
	result, err := db.Exec(query, cod, name, brand, desc, price, count)
	if err != nil {
		log.Fatalln(err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Printf("Added Product, %d rows affected", rowsAffected)
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
		log.Fatalln(row.Error())
		return err
	}

	fmt.Printf("Product: %s\nCod: %d\nBrand: %s\nDescription: %s\nPrice: %.2f\nInventory count: %d\n", product.Name, product.Cod, product.Brand, product.Description, product.Price, product.InventoryCount)

	return nil
}

func GetAllProducts() {
	return
}

func UpdateProducts() {
	return
}

func DeleteProducts() {
	return
}