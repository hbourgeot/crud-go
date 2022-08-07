package main

import (
	"github.com/hbourgeot/crud-go/internal/utilities"
)

func clientsMenu() {
	utilities.PrintTitle("TodoTech Crud Program")
	utilities.PrintSubtitle("Clients menu")
	utilities.PrintMenus("Add client", "Show clients...", "Update Client", "Delete client")
}

func productsMenu() {
	utilities.PrintTitle("TodoTech Crud Program")
	utilities.PrintSubtitle("Products menu")
	utilities.PrintMenus("Add Product", "Show products...", "Update Product", "Delete product")
}

func ordersMenu() {
	utilities.PrintTitle("TodoTech Crud Program")
	utilities.PrintSubtitle("Orders menu")
	utilities.PrintMenus("Add Order", "Show products...", "Update Product", "Delete product")
}