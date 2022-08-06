package database

import (
	"fmt"
	"log"
)

type Products struct {
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

func GetProductsByID(id int) {
	return
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
