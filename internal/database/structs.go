package database

type Clients struct {
	Name  string
	Phone string
	DNI   int
}

type Products struct {
	Cod            int
	Name           string
	Brand          string
	Description    string
	Price          float32
	InventoryCount int
}

type Orders struct {
	ID         int
	ProductCod int
	ClientDNI  int
}