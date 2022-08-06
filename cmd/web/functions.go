package main

import "github.com/hbourgeot/crud-go/internal/utilities"

func home() {
	utilities.PrintTitle("Welcome to TodoTech Crud Program")
	utilities.PrintSubtitle("Main menu")
	utilities.PrintMenus("Clients", "Products", "Orders")
}