package main

import (
	"fmt"
	"github.com/hbourgeot/crud-go/internal/utilities"
	"log"
	"os"
)

func main() {
	var option int
	for option != 4 {

		utilities.PrintTitle("Welcome to TodoTech Crud Program")
		utilities.PrintSubtitle("Main menu")
		utilities.PrintMenus("Clients...", "Products...", "Orders...", "Exit")
		reader := utilities.NewReader()

		option, err := utilities.ReadNumber(reader)
		if err != nil {
			log.Fatalln(err)
			return
		}

		switch option {
		case 1:
			clientsMenu()
			break
		case 2:
			productsMenu()
			break
		case 3:
			ordersMenu()
			break
		case 4:
			os.Exit(0)
			break
		default:
			break
		}
	}
}

func clientsMenu() {
	utilities.PrintTitle("TodoTech Crud Program")
	utilities.PrintSubtitle("Clients menu")
	utilities.PrintMenus("Add client", "Show clients...", "Update Client", "Delete client", "Go to main menu")
	reader := utilities.NewReader()
	option, err := utilities.ReadNumber(reader)

	if err != nil {
		log.Fatalln(err)
		return
	}

	switch option {
	case 1:
		addClient()
		break
	case 2:
		showClients()
		break
	case 3:
		updateClient()
		break
	case 4:
		deleteClient()
		break
	case 5:
		fmt.Println("Changing to main menu...")
		break
	default:
		fmt.Println("Invalid option")
	}
}

func productsMenu() {
	utilities.PrintTitle("TodoTech Crud Program")
	utilities.PrintSubtitle("Products menu")
	utilities.PrintMenus("Add Product", "Show products...", "Update Product", "Delete product", "Go to main menu")
	reader := utilities.NewReader()
	option, err := utilities.ReadNumber(reader)

	if err != nil {
		log.Fatalln(err)
		return
	}

	switch option {
	case 1:
		addProduct()
		break
	case 2:
		showProducts()
		break
	case 3:
		updateProduct()
		break
	case 4:
		deleteProduct()
		break
	case 5:
		fmt.Println("Changing to main menu...")
		break
	default:
		fmt.Println("Invalid option")
	}
}

func ordersMenu() {
	utilities.PrintTitle("TodoTech Crud Program")
	utilities.PrintSubtitle("Orders menu")
	utilities.PrintMenus("Add Order", "Show orders...", "Go to Main menu")
	reader := utilities.NewReader()
	option, err := utilities.ReadNumber(reader)

	if err != nil {
		log.Fatalln(err)
		return
	}

	switch option {
	case 1:
		createOrder()
		break
	case 2:
		readOrders()
		break
	case 3:
		fmt.Println("Changing to main menu...")
		break
	default:
		fmt.Println("Invalid option")
	}
}