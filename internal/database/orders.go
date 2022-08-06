package database

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

func GetOrdersByID(id int) {
	return
}

func GetAllOrders() {
	return
}

func GetOrdersByClient(dni int) {
	return
}

func GetOrdersByProduct(cod int) {
	return
}