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

func GetOrderByDNI(dni int) ([]*Orders, error) {
	db, err := makeCN()
	if err != nil {
		return nil, err
	}

	orders := []*Orders{}

	query := "SELECT * FROM orders WHERE dni = $1"
	rows, err := db.Query(query, dni)
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

func GetOrderByCode(code int) ([]*Orders, error) {
	db, err := makeCN()
	if err != nil {
		return nil, err
	}

	orders := []*Orders{}

	query := "SELECT * FROM orders WHERE cod = $1"
	rows, err := db.Query(query, code)
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

func GetAllOrders() ([]*Orders, error) {
	db, err := makeCN()
	if err != nil {
		return nil, err
	}

	orders := []*Orders{}

	query := "SELECT * FROM orders"
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