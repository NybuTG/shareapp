package main

import (
	"fmt"
	"net"
	"share/networking"
	"share/networking/ft"
)

func parseNetAddr(ipa net.Addr) (string) {
	ipString := ipa.String()
	return ipString
}


func server() {
	// Send ping and save users to array
	users := networking.UserCollect()

	// If users are found: Print their IP and Port
	if len(users) > 0 {
		fmt.Println("USERS:")
		for i := 0; i < len(users); i++ {
			fmt.Printf("%d. %s\n", i + 1, users[i])
		}
		fmt.Printf("\nChoose from the list: ")
		input, _ := fmt.Scanf("%d")
		ft.Send(parseNetAddr(users[input - 1]))
	}
}

func client() {
	networking.Listener()
	ft.Receive()
}


func main() {
	server()
}

