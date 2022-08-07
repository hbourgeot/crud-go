package main

import (
	"fmt"
	"github.com/hbourgeot/crud-go/internal/database"
	"github.com/hbourgeot/crud-go/internal/utilities"
	"log"
	"os"
	"strings"
)

func addClient() {
	reader := utilities.NewReader()
	fmt.Print("Enter Client DNI: ")
	clientDNI, err := utilities.ReadNumber(reader)
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Print("Enter Client Name: ")
	line, err := utilities.ReadLine(reader)
	if err != nil {
		log.Fatalln(err)
		return
	}
	clientName := line

	fmt.Print("Enter Client Phone number: ")
	line, err = utilities.ReadLine(reader)
	if err != nil {
		log.Fatalln(err)
		return
	}
	clientPhone := line

	err = database.CreateClients(clientDNI, clientName, clientPhone)
	if err != nil {
		log.Fatalln(err)
		return
	}
	utilities.PrintTitle("Client registered successfully")
	utilities.PrintSubtitle("Options menu")
	utilities.PrintMenus("Generate Order", "Go to Clients menu", "Go to Main menu", "Exit")

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
		clientsMenu()
		break
	case 3:
		fmt.Println("Changing to main menu...")
		break
	case 4:
		os.Exit(0)
		break
	default:
		fmt.Println("Invalid option")
	}
}

func showClients() {
	utilities.PrintTitle("TodoTech CRUD Program")
	utilities.PrintSubtitle("Show Clients menu")
	utilities.PrintMenus("Show Client by DNI", "Show All Clients", "Go to Client menu", "Go to Main Menu", "Exit")
	reader := utilities.NewReader()
	option, err := utilities.ReadNumber(reader)
	if err != nil {
		log.Fatalln(err)
		return
	}

	switch option {
	case 1:
		fmt.Print("Enter Client DNI: ")
		clientDNI, err := utilities.ReadNumber(reader)
		if err != nil {
			log.Fatalln(err)
			return
		}

		err = database.GetClientByDNI(clientDNI)
		if err != nil {
			log.Fatalln(err)
			return
		}

		fmt.Println("\n\n\n")
		break
	case 2:
		var count int
		clients, err := database.GetAllClients()
		if err != nil {
			log.Fatalln(err)
			return
		}

		for _, c := range clients {
			fmt.Printf("Name: %s\nDNI: %d\nPhone number: %s\n", c.Name, c.DNI, c.Phone)
			if count < 15 {
				count = 15
			}
			if count < len(c.Name) {
				count = len(c.Name)
			}
			if count < len(c.Phone) {
				count = len(c.Phone)
			}
			utilities.PrintSeparator(count + 14)
		}
		fmt.Print("\n\n")
		break
	case 3:
		clientsMenu()
		break
	case 4:
		break
	default:
		fmt.Println("Invalid option")

	}
}

func updateClient() {

	reader := utilities.NewReader()

	fmt.Print("Enter DNI of the client to update: ")
	numbers, err := utilities.ReadNumber(reader)
	if err != nil {
		log.Fatalln(err)
		return
	}

	err = database.GetClientByDNI(dni)
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Print("\nEnter the name of the data to modify: ")
	line, err := utilities.ReadLine(reader)
	if err != nil {
		log.Fatalln(err)
		return
	}

	colToUpdate := strings.ToLower(line)
	var newVal string

	fmt.Print("\nEnter the value: ")
	if colToUpdate == "name" {
		line, err = utilities.ReadLine(reader)
		newVal = line
	} else if colToUpdate == "dni" || colToUpdate == "phone" {
		numbers, err = utilities.ReadNumber(reader)
		newVal = numbers
	}

	err = database.UpdateClient(colToUpdate, newVal, dni)

}

func deleteClient() {

}
func addProduct() {

}

func showProducts() {

}

func updateProduct() {

}

func deleteProduct() {

}

func createOrder() {

}

func readOrders() {

}