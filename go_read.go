package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Address struct {
	Type    string `json:"type"`
	Address string `json:"address"`
}

type Item struct {
	ItemID int     `json:"item_id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
}

type Order struct {
	OrderID int     `json:"order_id"`
	Items   []Item  `json:"items"` //do tipo slice
	Total   float64 `json:"total"`
}

type Settings struct {
	Theme                string `json:"theme"`
	Notifications        bool   `json:"notifications"`
	NewsletterSubscribed bool   `json:"newsletter_subscribed"`
}

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Addresses []Address `json:"addresses"` //Será um slice
	Orders    []Order   `json:"orders"`
	Settings  Settings  `json:"settings"`
}

type UserProfile struct {
	User User `json:"user"`
}

func main() {
	file, err := os.Open("cart.json") //Abre o arquivo

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close() //Fecha o arquivo "cart.json"

	byteValue, err := io.ReadAll(file)

	if err != nil {
		log.Fatal(err)
	}

	var userProfile UserProfile
	err = json.Unmarshal(byteValue, &userProfile)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("user ID: %d\n", userProfile.User.ID)
	fmt.Printf("User Name: %s\n", userProfile.User.Name)
	fmt.Printf("User Email: %s\n", userProfile.User.Email)

	//Imprime cada endereço:
	for _, address := range userProfile.User.Addresses {
		fmt.Printf("Address (%s): %s\n", address.Type, address.Address)
	}

	var accumulatedTotal float64

	// Percorre a lista de pedidos:
	for _, order := range userProfile.User.Orders {
			var currentOrderTotal float64 // Inicializa uma variável para o total do pedido atual
			
			// Percorre os itens do pedido atual e soma os preços:
			for _, item := range order.Items {
					fmt.Printf("- Item ID: %d, Name: %s, Price: %.2f\n", item.ItemID, item.Name, item.Price)
					currentOrderTotal += item.Price // Soma o preço do item ao total do pedido atual
			}
			
			// Imprime o total do pedido atual
			fmt.Printf("Total Pedido %d: %.2f\n", order.OrderID, currentOrderTotal)
			
			fmt.Println()

			// Adiciona o total do pedido atual ao total acumulado
			accumulatedTotal += currentOrderTotal
	}
	
	// Imprime o total acumulado de todos os pedidos
	fmt.Printf("Total Acumulado de Todos os Pedidos: %.2f\n", accumulatedTotal)

	fmt.Println()

	fmt.Printf("Theme: %s, Notifications: %t, Newsletter_Subscribed: %t\n",
		userProfile.User.Settings.Theme, userProfile.User.Settings.Notifications, userProfile.User.Settings.NewsletterSubscribed)

}
