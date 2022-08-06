package database

import (
	"fmt"
)

func CreateClients(dni int, name string, phone string, ) error {
	db, err := makeCN()

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

func GetClientByDNI(id int) error {
	db, err := makeCN()
	if err != nil {
		return err
	}

	var client Clients

	query := "SELECT * FROM clients WHERE dni = $1"
	row := db.QueryRow(query, id).Scan(&client.DNI)
	if row.Error() != "" {
		return err
	}

	fmt.Printf("Client: %s\nDNI: %d\nPhone: %s\n", client.Name, client.DNI, client.Phone)

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

	query := "UPDATE clients SET $1 = $2 WHERE dni = $3"
	_, err = db.Exec(query, columnEdit, newValue, dni)
	if err != nil {
		return err
	}

	return nil
}

func DeleteClients(dni int) error {
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