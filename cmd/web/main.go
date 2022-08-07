package main

import (
	"github.com/hbourgeot/crud-go/internal/utilities"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var option int
	for option != 4 {

		utilities.PrintTitle("Welcome to TodoTech Crud Program")
		utilities.PrintSubtitle("Main menu")
		utilities.PrintMenus("Clients", "Products", "Orders", "Exit")
		reader := utilities.NewReader()

		num, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
			return
		}

		num = strings.TrimSuffix(num, "\r\n")

		option, err = strconv.Atoi(num)
		if err != nil {
			log.Fatalln(err)
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