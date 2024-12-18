package structs

import (
	"encoding/json"
	"fmt"
	"myFirstProject/Structs/embed"
)

type Client struct {
	embed.Foo

	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address Address
	Email   string
}

type Address struct {
	Address string
	Number  int
	ZipCode string
	Country string
}

func Struct() {
	client1 := &Client{
		Name: "Franklin",
		Age:  32,
		Address: Address{ // Initialize the Address struct here
			Country: "Brazil",
			Number:  123,
		},
	}

	client1.Email = "frankmcdias@gmail.com"
	client1.Address.Number = 123
	client1.Address.Country = "Brazil"

	client1.updateClient("client update")

	client1.Bar()
	res, err := json.Marshal(client1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	// client1.greetings()
}

// func (p Client) greetings() {
// 	fmt.Printf("Hello, %s from %s\n", p.Name, p.Address.Country)
// }

func (u *Client) updateClient(newName string) {
	u.Name = newName
}
