package database

import (
	"fmt"
	_ "github.com/lib/pq"
)

func CreateClient(dni int, name string, phone string) error {
	db, err := makeCN()
	if err != nil {
		return err
	}

	query := "INSERT INTO clients (dni,name,phone) VALUES ($1,$2,$3)"
	result, err := db.Exec(query, dni, name, phone)
	if err != nil {

		return err
	}

	_, err = result.RowsAffected()
	if err != nil {

		return err
	}

	return nil
}

func GetClientByDNI(dni int) error {
	db, err := makeCN()
	if err != nil {
		return err
	}

	var client Clients

	query := "SELECT * FROM clients WHERE dni = $1"
	row := db.QueryRow(query, dni)
	err = row.Scan(&client.DNI, &client.Name, &client.Phone)
	if err != nil {
		return err
	}

	fmt.Printf("Name: %s\nDNI: %d\nPhone: %s\n", client.Name, client.DNI, client.Phone)

	return nil
}

func GetAllClients() ([]*Clients, error) {
	db, err := makeCN()
	if err != nil {
		return nil, err
	}

	clients := []*Clients{}

	query := "SELECT * FROM clients"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		c := &Clients{}
		err := rows.Scan(&c.DNI, &c.Name, &c.Phone)
		if err != nil {
			return nil, err
		}

		clients = append(clients, c)
	}
	defer rows.Close()

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return clients, nil
}

func UpdateClient(columnEdit string, newValue string, dni int) error {
	db, err := makeCN()
	if err != nil {
		return err
	}
	query := ""
	if columnEdit == "name" {
		query = "UPDATE clients SET name = $1 WHERE dni = $2"
	} else if columnEdit == "phone" {
		query = "UPDATE clients SET phone = $1 WHERE dni = $2"
	} else {
		query = "UPDATE clients SET dni = $1 WHERE dni = $2"
	}
	_, err = db.Exec(query, newValue, dni)
	if err != nil {
		return err
	}

	return nil
}

func DeleteClient(dni int) error {
	db, err := makeCN()
	if err != nil {
		return err
	}

	query := "DELETE FROM clients WHERE dni = $1"
	_, err = db.Exec(query, dni)
	if err != nil {
		return err
	}

	return nil
}