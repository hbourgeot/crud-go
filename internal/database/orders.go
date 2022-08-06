package database

import "fmt"

func GenerateOrder(productCod, clientDNI int) error {
	db, err := makeCN()

	query := "INSERT INTO orders (dni, cod) VALUES ($1,$2)"
	result, err := db.Exec(query, clientDNI, productCod)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func GetOrderByID(id int) error {
	db, err := makeCN()
	if err != nil {
		return err
	}

	var order Orders

	var clientName, productName string

	query := "SELECT * FROM orders WHERE dni = $1"
	row := db.QueryRow(query, id)
	err = row.Scan(&order.ClientDNI)
	if err != nil {
		return err
	}

	query = "SELECT orders.id, clients.name, products.name FROM orders INNER JOIN clients ON clients.dni = $1 INNER JOIN products ON products.cod = $2 WHERE orders.id = $3"
	row = db.QueryRow(query, order.ClientDNI, order.ProductCod, id)
	err = row.Scan(&clientName, &productName)
	if err != nil {
		return err
	}

	fmt.Printf("Order ID: %d\nClient: %s\nProduct: %s\n", id, clientName, productName)

	return nil
}

func GetAllOrders() ([]*Orders, error) {
	db, err := makeCN()
	if err != nil {
		return nil, err
	}

	orders := []*Orders{}

	query := "SELECT * FROM clients"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		o := &Orders{}
		err := rows.Scan(&o.ID, &o.ClientDNI, &o.ProductCod)
		if err != nil {
			return nil, err
		}

		orders = append(orders, o)
	}
	defer rows.Close()

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}