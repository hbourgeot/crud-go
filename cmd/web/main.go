package main

import (
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

		option, err := utilities.ReadOption(reader)
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