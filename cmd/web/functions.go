package main

import (
	"fmt"
	"github.com/hbourgeot/crud-go/internal/database"
	"github.com/hbourgeot/crud-go/internal/utilities"
	"log"
	"os"
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

}

func updateClient() {

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