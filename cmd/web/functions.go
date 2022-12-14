package main

import (
	"fmt"
	"github.com/hbourgeot/crud-go/internal/database"
	"github.com/hbourgeot/crud-go/internal/utilities"
	"log"
	"os"
	"strconv"
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

	err = database.CreateClient(clientDNI, clientName, clientPhone)
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
			log.Println(err)
		}

		fmt.Print("\n\n\n")
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
		break
	}

	utilities.PrintSubtitle("Options menu")
	utilities.PrintMenus("Generate Order", "Go to Clients menu", "Go to Main menu", "Exit")

	option, err = utilities.ReadNumber(reader)

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
	default:
		fmt.Println("Invalid option")
	}
}

func updateClient() {

	reader := utilities.NewReader()

	fmt.Print("Enter DNI of the client to update: ")
	dni, err := utilities.ReadNumber(reader)
	if err != nil {
		log.Fatalln(err)
		return
	}

	err = database.GetClientByDNI(dni)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Print("\nEnter the name of the data to modify: ")
	line, err := utilities.ReadLine(reader)
	if err != nil {
		log.Fatalln(err)
		return
	}

	colToUpdate := strings.ToLower(line)

	fmt.Print("\nEnter the value: ")
	line, err = utilities.ReadLine(reader)
	err = database.UpdateClient(colToUpdate, line, dni)
	if err != nil {
		log.Println(err)
		return
	}

	utilities.PrintTitle("Client updated successfully")
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
	default:
		fmt.Println("Invalid option")
	}
}

func deleteClient() {
	reader := utilities.NewReader()
	fmt.Print("Enter the Client DNI: ")
	dni, err := utilities.ReadNumber(reader)
	if err != nil {
		log.Fatalln(err)
		return
	}

	err = database.GetClientByDNI(dni)
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Print("Are you sure you want to delete the client? y/n: ")
	line, err := utilities.ReadLine(reader)
	if err != nil {
		log.Fatalln(err)
		return
	}
	if line == "y" {
		err = database.DeleteClient(dni)
		if err != nil {
			log.Println(err)
			return
		}
	}
	utilities.PrintTitle("Client deleted successfully")
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
	default:
		fmt.Println("Invalid option")
	}
}

func addProduct() {
	reader := utilities.NewReader()
	fmt.Print("Enter Product code: ")
	productCod, err := utilities.ReadNumber(reader)
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Print("Enter Product Name: ")
	line, err := utilities.ReadLine(reader)
	if err != nil {
		log.Fatalln(err)
		return
	}
	productName := line

	fmt.Print("Enter Product Brand: ")
	line, err = utilities.ReadLine(reader)
	if err != nil {
		log.Fatalln(err)
		return
	}
	productBrand := line

	fmt.Print("Enter Product Description: ")
	line, err = utilities.ReadLine(reader)
	if err != nil {
		log.Fatalln(err)
		return
	}
	productDesc := line

	fmt.Print("Enter Product Price: ")
	productPrice, err := utilities.ReadPrice(reader)
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Print("Enter Inventory Count: ")
	productCount, err := utilities.ReadNumber(reader)
	if err != nil {
		log.Fatalln(err)
		return
	}

	err = database.InsertProducts(productCod, productName, productBrand, productDesc, productPrice, productCount)
	if err != nil {
		log.Fatalln(err)
		return
	}
	utilities.PrintTitle("Product registered successfully")
	utilities.PrintSubtitle("Options menu")
	utilities.PrintMenus("Generate Order", "Go to Products menu", "Go to Main menu", "Exit")

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
		productsMenu()
		break
	case 3:
		fmt.Println("Changing to main menu...")
		break
	case 4:
		os.Exit(0)
	default:
		fmt.Println("Invalid option")
	}
}

func showProducts() {
	utilities.PrintTitle("TodoTech CRUD Program")
	utilities.PrintSubtitle("Show Products menu")
	utilities.PrintMenus("Show Product by Code", "Show All Products", "Go to Products menu", "Go to Main Menu", "Exit")
	reader := utilities.NewReader()
	option, err := utilities.ReadNumber(reader)
	if err != nil {
		log.Fatalln(err)
		return
	}

	switch option {
	case 1:
		fmt.Print("Enter Product Code: ")
		productCode, err := utilities.ReadNumber(reader)
		if err != nil {
			log.Fatalln(err)
			return
		}

		err = database.GetProductsByCode(productCode)
		if err != nil {
			log.Println(err)
		}

		fmt.Print("\n\n\n")
		break
	case 2:
		var count int
		products, err := database.GetAllProducts()
		if err != nil {
			log.Println(err)
		}

		for _, p := range products {
			fmt.Printf("Product: %s\nCod: %d\nBrand: %s\nDescription: %s\nPrice: %.2f\nInventory count: %d\n", p.Name, p.Cod, p.Brand, p.Description, p.Price, p.InventoryCount)
			if count < 15 {
				count = 15
			}
			if count < len(p.Description) {
				count = len(p.Description)
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

	utilities.PrintSubtitle("Options menu")
	utilities.PrintMenus("Generate Order", "Go to Products menu", "Go to Main menu", "Exit")

	option, err = utilities.ReadNumber(reader)

	if err != nil {
		log.Fatalln(err)
		return
	}

	switch option {
	case 1:
		createOrder()
		break
	case 2:
		productsMenu()
		break
	case 3:
		fmt.Println("Changing to main menu...")
		break
	case 4:
		os.Exit(0)
	default:
		fmt.Println("Invalid option")
	}
}

func updateProduct() {
	reader := utilities.NewReader()

	fmt.Print("Enter Code of the product to update: ")
	code, err := utilities.ReadNumber(reader)
	if err != nil {
		log.Fatalln(err)
		return
	}

	err = database.GetProductsByCode(code)
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

	fmt.Print("\nEnter the value: ")
	line, err = utilities.ReadLine(reader)

	switch colToUpdate {
	case "code", "inventory count":
		newVal, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalln(err)
			return
		}

		err = database.UpdateProducts(colToUpdate, newVal, code)
		if err != nil {
			log.Fatalln(err)
			return
		}
		break
	case "name", "brand", "description":
		err = database.UpdateProducts(colToUpdate, line, code)
		if err != nil {
			log.Fatalln(err)
			return
		}
		break
	case "price":
		newVal, err := strconv.ParseFloat(line, 64)
		if err != nil {
			log.Fatalln(err)
			return
		}

		err = database.UpdateProducts(colToUpdate, newVal, code)
		if err != nil {
			log.Println(err)
			return
		}
		break
	}

	utilities.PrintTitle("Product updated successfully")
	utilities.PrintSubtitle("Options menu")
	utilities.PrintMenus("Generate Order", "Go to Products menu", "Go to Main menu", "Exit")

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
		productsMenu()
		break
	case 3:
		fmt.Println("Changing to main menu...")
		break
	case 4:
		os.Exit(0)
	default:
		fmt.Println("Invalid option")
	}
}

func deleteProduct() {
	reader := utilities.NewReader()
	fmt.Print("Enter the Product Code: ")
	code, err := utilities.ReadNumber(reader)
	if err != nil {
		log.Fatalln(err)
		return
	}

	err = database.GetProductsByCode(code)
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Print("Are you sure you want to delete the product? y/n: ")
	line, err := utilities.ReadLine(reader)
	if err != nil {
		log.Fatalln(err)
		return
	}
	if line == "y" {
		err = database.DeleteProducts(code)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("Client deleted.")
	}

	utilities.PrintTitle("Product deleted successfully")
	utilities.PrintSubtitle("Options menu")
	utilities.PrintMenus("Generate Order", "Go to Products menu", "Go to Main menu", "Exit")

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
		productsMenu()
		break
	case 3:
		fmt.Println("Changing to main menu...")
		break
	case 4:
		os.Exit(0)
	default:
		fmt.Println("Invalid option")
	}
}

func createOrder() {
	reader := utilities.NewReader()

	fmt.Print("Enter Client DNI: ")
	dni, err := utilities.ReadNumber(reader)
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Print("Enter Product code: ")
	code, err := utilities.ReadNumber(reader)
	if err != nil {
		log.Fatalln(err)
		return
	}

	utilities.PrintTitle(" Order Resume ")
	err = database.GetClientByDNI(dni)
	if err != nil {
		log.Fatalln(err)
		return
	}
	err = database.GetProductsByCode(code)
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Print("\n\nGenerate order? y/n:")
	line, err := utilities.ReadLine(reader)
	if err != nil {
		log.Fatalln(err)
		return
	}

	if line == "y" {
		err = database.GenerateOrder(code, dni)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}

	utilities.PrintTitle("Order generated successfully")
	utilities.PrintSubtitle("Options menu")
	utilities.PrintMenus("Generate other Order", "Go to Orders menu", "Go to Main menu", "Exit")

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
		ordersMenu()
		break
	case 3:
		fmt.Println("Changing to main menu...")
		break
	case 4:
		os.Exit(0)
	default:
		fmt.Println("Invalid option")
	}
}

func readOrders() {
	utilities.PrintTitle("TodoTech CRUD Program")
	utilities.PrintSubtitle("Show Clients menu")
	utilities.PrintMenus("Show Orders by DNI", "Show Orders By Product Code", "Show All Orders", "Go to Orders menu", "Go to Main Menu")
	reader := utilities.NewReader()
	option, err := utilities.ReadNumber(reader)
	if err != nil {
		log.Fatalln(err)
		return
	}

	switch option {
	case 1:
		fmt.Print("Enter Client DNI: ")
		dni, err := utilities.ReadNumber(reader)
		if err != nil {
			log.Fatalln(err)
			return
		}

		var count = 20
		orders, err := database.GetOrdersByDNI(dni)
		if err != nil {
			log.Fatalln(err)
			return
		}

		for _, o := range orders {
			fmt.Printf("ID: %d\nClient DNI: %d\nProduct Code: %d\n", o.ID, o.ClientDNI, o.ProductCod)
			utilities.PrintSeparator(count)
		}
		fmt.Print("\n\n")
		break
	case 2:
		fmt.Print("Enter Client DNI: ")
		code, err := utilities.ReadNumber(reader)
		if err != nil {
			log.Fatalln(err)
			return
		}

		var count = 20
		orders, err := database.GetOrdersByCode(code)
		if err != nil {
			log.Fatalln(err)
			return
		}

		for _, o := range orders {
			fmt.Printf("ID: %d\nClient DNI: %d\nProduct Code: %d\n", o.ID, o.ClientDNI, o.ProductCod)
			utilities.PrintSeparator(count)
		}
		fmt.Print("\n\n")
		break
	case 3:
		var count = 20
		orders, err := database.GetAllOrders()
		if err != nil {
			log.Fatalln(err)
			return
		}

		for _, o := range orders {
			fmt.Printf("ID: %d\nClient DNI: %d\nProduct Code: %d\n", o.ID, o.ClientDNI, o.ProductCod)
			utilities.PrintSeparator(count)
		}
		fmt.Print("\n\n")
		break
	case 4:
		ordersMenu()
		break
	case 5:
		break
	default:
		fmt.Println("Invalid option")

	}

	utilities.PrintTitle("Order generated successfully")
	utilities.PrintSubtitle("Options menu")
	utilities.PrintMenus("Generate other Order", "Go to Orders menu", "Go to Main menu", "Exit")

	option, err = utilities.ReadNumber(reader)

	if err != nil {
		log.Fatalln(err)
		return
	}

	switch option {
	case 1:
		createOrder()
		break
	case 2:
		ordersMenu()
		break
	case 3:
		fmt.Println("Changing to main menu...")
		break
	case 4:
		os.Exit(0)
	default:
		fmt.Println("Invalid option")
	}
}